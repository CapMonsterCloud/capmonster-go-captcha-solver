package capmonster

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// Version of the SDK, also sent in the User-Agent. Bumped on every release (see CHANGELOG.md).
	Version = "2.0.0"
	// DefaultServiceURL is the CapMonster Cloud API endpoint.
	DefaultServiceURL = "https://api.capmonster.cloud"
	// DefaultSoftID is sent as softId with every task.
	DefaultSoftID = 58
	// DefaultClientTimeout is the timeout of ONE HTTP request (not of the whole solving).
	DefaultClientTimeout = 20 * time.Second

	userAgent = "Zennolab.CapMonsterCloud.Client.Go/" + Version
)

// Client talks to CapMonster Cloud. It is safe for concurrent use; create one and reuse it.
type Client struct {
	apiKey           string
	serviceURL       string
	softID           int
	httpClient       *http.Client
	timeoutsOverride *Timeouts
}

// Option configures a Client.
type Option func(*Client) error

// WithServiceURL replaces the API endpoint.
func WithServiceURL(serviceURL string) Option {
	return func(c *Client) error {
		c.serviceURL = strings.TrimRight(serviceURL, "/")
		return nil
	}
}

// WithSoftID replaces the softId sent with every task.
func WithSoftID(softID int) Option {
	return func(c *Client) error {
		c.softID = softID
		return nil
	}
}

// WithHTTPClient makes the client send its requests through httpClient (custom transport, timeout, proxy...).
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		if httpClient == nil {
			return errors.New("capmonster: nil http client")
		}
		c.httpClient = httpClient
		return nil
	}
}

// WithClientTimeout sets the timeout of one HTTP request to the API.
func WithClientTimeout(timeout time.Duration) Option {
	return func(c *Client) error {
		c.httpClient.Timeout = timeout
		return nil
	}
}

// WithClientProxy sends the SDK's own requests to the API through a proxy, e.g.
// "http://user:pass@10.0.0.1:3128" or "socks5://10.0.0.1:1080".
// This is not the proxy used to solve a captcha: that one is set on the request.
func WithClientProxy(proxyURL string) Option {
	return func(c *Client) error {
		u, err := url.Parse(proxyURL)
		if err != nil || u.Host == "" {
			return fmt.Errorf("capmonster: invalid client proxy URL %q", proxyURL)
		}
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.Proxy = http.ProxyURL(u)
		c.httpClient.Transport = transport
		return nil
	}
}

// WithTimeoutsOverride replaces the per-captcha-type polling schedule (mostly useful for tests).
func WithTimeoutsOverride(t Timeouts) Option {
	return func(c *Client) error {
		c.timeoutsOverride = &t
		return nil
	}
}

// New creates a client. It panics only if an Option returns an error; use NewWithOptions to handle it.
func New(apiKey string, opts ...Option) *Client {
	c, err := NewWithOptions(apiKey, opts...)
	if err != nil {
		panic(err)
	}
	return c
}

// NewWithOptions creates a client and reports invalid options as an error.
func NewWithOptions(apiKey string, opts ...Option) (*Client, error) {
	c := &Client{
		apiKey:     apiKey,
		serviceURL: DefaultServiceURL,
		softID:     DefaultSoftID,
		httpClient: &http.Client{Timeout: DefaultClientTimeout},
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// GetBalance returns the account balance in USD.
func (c *Client) GetBalance(ctx context.Context) (float64, error) {
	var resp struct {
		apiResponse
		Balance *float64 `json:"balance"`
	}
	if err := c.call(ctx, "getBalance", map[string]any{"clientKey": c.apiKey}, &resp); err != nil {
		return 0, err
	}
	if err := resp.err("getBalance"); err != nil {
		return 0, err
	}
	if resp.Balance == nil {
		return 0, &MalformedResponseError{Method: "getBalance", Reason: "no balance field"}
	}
	return *resp.Balance, nil
}

// CreateTask creates a task and returns its id (low-level; most users want Solve).
func (c *Client) CreateTask(ctx context.Context, req Request) (int64, error) {
	task, err := req.Task()
	if err != nil {
		return 0, err
	}
	var resp struct {
		apiResponse
		TaskID *int64 `json:"taskId"`
	}
	body := map[string]any{"clientKey": c.apiKey, "task": task, "softId": c.softID}
	if err := c.call(ctx, "createTask", body, &resp); err != nil {
		return 0, err
	}
	if err := resp.err("createTask"); err != nil {
		return 0, err
	}
	if resp.TaskID == nil {
		return 0, &MalformedResponseError{Method: "createTask", Reason: "no taskId field"}
	}
	return *resp.TaskID, nil
}

// TaskResult is the state of a task returned by GetTaskResult.
type TaskResult struct {
	Ready    bool
	Solution Solution // nil while processing
}

// GetTaskResult makes one getTaskResult call (low-level).
func (c *Client) GetTaskResult(ctx context.Context, taskID int64) (TaskResult, error) {
	var resp struct {
		apiResponse
		Status   string          `json:"status"`
		Solution json.RawMessage `json:"solution"`
	}
	err := c.call(ctx, "getTaskResult", map[string]any{"clientKey": c.apiKey, "taskId": taskID}, &resp)
	var statusErr *StatusError
	// The service sometimes answers 500/503 while the task is still being processed.
	if errors.As(err, &statusErr) && (statusErr.Status == 500 || statusErr.Status == 503) {
		return TaskResult{}, nil
	}
	if err != nil {
		return TaskResult{}, err
	}
	if err := resp.err("getTaskResult"); err != nil {
		return TaskResult{}, err
	}
	if resp.Status != "ready" {
		return TaskResult{}, nil
	}
	return TaskResult{Ready: true, Solution: Solution(resp.Solution)}, nil
}

// Solve creates a task, polls until it is solved and returns the solution object.
//
// It returns a *TimeoutError (errors.Is(err, ErrTimeout)) when the captcha type's maximum waiting time is
// exceeded, and ctx.Err() when ctx is cancelled first.
func (c *Client) Solve(ctx context.Context, req Request) (Solution, error) {
	timeouts := req.Timeouts()
	if c.timeoutsOverride != nil {
		timeouts = *c.timeoutsOverride
	}
	taskID, err := c.CreateTask(ctx, req)
	if err != nil {
		return nil, err
	}

	deadline := time.Now().Add(timeouts.Timeout)
	if err := sleep(ctx, timeouts.FirstRequestDelay); err != nil {
		return nil, err
	}
	for time.Now().Before(deadline) {
		result, err := c.GetTaskResult(ctx, taskID)
		if err != nil {
			return nil, err
		}
		if result.Ready {
			return result.Solution, nil
		}
		if err := sleep(ctx, min(timeouts.RequestsInterval, time.Until(deadline))); err != nil {
			return nil, err
		}
	}
	return nil, &TimeoutError{TaskID: taskID, Timeout: timeouts.Timeout}
}

// apiResponse is the envelope shared by all API methods.
type apiResponse struct {
	ErrorID          *int   `json:"errorId"`
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}

func (r apiResponse) err(method string) error {
	if *r.ErrorID != 0 {
		return &APIError{Method: method, Code: r.ErrorCode, Description: r.ErrorDescription}
	}
	return nil
}

func (c *Client) call(ctx context.Context, method string, body any, out any) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("capmonster: %s: encode request: %w", method, err)
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.serviceURL+"/"+method, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("capmonster: %s: %w", method, err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("User-Agent", userAgent)

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("capmonster: %s: %w", method, err)
	}
	defer httpResp.Body.Close()
	raw, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return fmt.Errorf("capmonster: %s: read response: %w", method, err)
	}

	var envelope apiResponse
	isEnvelope := json.Unmarshal(raw, &envelope) == nil && envelope.ErrorID != nil
	isSuccess := httpResp.StatusCode >= 200 && httpResp.StatusCode < 300
	// An error JSON (errorId != 0) is reported as an API error whatever the HTTP status is.
	if isEnvelope && (isSuccess || *envelope.ErrorID != 0) {
		if err := json.Unmarshal(raw, out); err != nil {
			return &MalformedResponseError{Method: method, Reason: err.Error()}
		}
		return nil
	}
	if !isSuccess {
		return &StatusError{Method: method, Status: httpResp.StatusCode}
	}
	return &MalformedResponseError{Method: method, Reason: fmt.Sprintf("unexpected body: %.200s", raw)}
}

func sleep(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return ctx.Err()
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
