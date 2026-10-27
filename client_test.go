package capmonster

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

// fakeAPI replays canned responses per API method (the last one of a method repeats) and records request bodies.
type fakeAPI struct {
	mu        sync.Mutex
	responses map[string][]fakeResponse
	bodies    map[string][]map[string]any
	userAgent string
}

type fakeResponse struct {
	status int
	body   string
}

func newFakeAPI(t *testing.T) (*fakeAPI, *Client) {
	f := &fakeAPI{responses: map[string][]fakeResponse{}, bodies: map[string][]map[string]any{}}
	srv := httptest.NewServer(f)
	t.Cleanup(srv.Close)
	client := New("KEY", WithServiceURL(srv.URL+"/"), WithTimeoutsOverride(Timeouts{
		FirstRequestDelay: 10 * time.Millisecond, RequestsInterval: 10 * time.Millisecond, Timeout: 300 * time.Millisecond,
	}))
	return f, client
}

func (f *fakeAPI) on(method string, status int, body string) *fakeAPI {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.responses[method] = append(f.responses[method], fakeResponse{status, body})
	return f
}

func (f *fakeAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	method := strings.TrimPrefix(r.URL.Path, "/")
	raw, _ := io.ReadAll(r.Body)
	var body map[string]any
	_ = json.Unmarshal(raw, &body)
	f.bodies[method] = append(f.bodies[method], body)
	f.userAgent = r.Header.Get("User-Agent")
	queue := f.responses[method]
	if len(queue) == 0 {
		http.Error(w, "no fake response for "+method, http.StatusTeapot)
		return
	}
	resp := queue[0]
	if len(queue) > 1 {
		f.responses[method] = queue[1:]
	}
	w.WriteHeader(resp.status)
	_, _ = w.Write([]byte(resp.body))
}

func request() RecaptchaV2Request {
	return RecaptchaV2Request{WebsiteURL: "https://example.com", WebsiteKey: "SITEKEY"}
}

func TestSolveCreatesTaskAndPollsUntilReady(t *testing.T) {
	api, client := newFakeAPI(t)
	api.on("createTask", 200, `{"errorId":0,"taskId":42}`).
		on("getTaskResult", 200, `{"errorId":0,"status":"processing"}`).
		on("getTaskResult", 200, `{"errorId":0,"status":"ready","solution":{"gRecaptchaResponse":"TOKEN"}}`)

	solution, err := client.Solve(context.Background(), request())
	if err != nil {
		t.Fatal(err)
	}
	if got := solution.String("gRecaptchaResponse"); got != "TOKEN" {
		t.Errorf("token = %q", got)
	}
	var decoded struct{ GRecaptchaResponse string }
	if err := solution.Decode(&decoded); err != nil || decoded.GRecaptchaResponse != "TOKEN" {
		t.Errorf("Decode: %v %+v", err, decoded)
	}
	assertJSONEqual(t, api.bodies["createTask"][0], `{"clientKey":"KEY","softId":58,
		"task":{"type":"RecaptchaV2Task","websiteURL":"https://example.com","websiteKey":"SITEKEY"}}`)
	assertJSONEqual(t, api.bodies["getTaskResult"], `[{"clientKey":"KEY","taskId":42},{"clientKey":"KEY","taskId":42}]`)
	if api.userAgent != "Zennolab.CapMonsterCloud.Client.Go/"+Version {
		t.Errorf("User-Agent = %q", api.userAgent)
	}
}

func TestCreateTaskAPIErrorIsReported(t *testing.T) {
	api, client := newFakeAPI(t)
	api.on("createTask", 200, `{"errorId":1,"errorCode":"ERROR_KEY_DOES_NOT_EXIST","errorDescription":"bad key"}`)

	_, err := client.Solve(context.Background(), request())
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Code != "ERROR_KEY_DOES_NOT_EXIST" || apiErr.Method != "createTask" {
		t.Fatalf("want APIError ERROR_KEY_DOES_NOT_EXIST, got %v", err)
	}
}

func TestErrorJSONWithHTTPErrorStatusIsAnAPIError(t *testing.T) {
	api, client := newFakeAPI(t)
	api.on("createTask", 400, `{"errorId":1,"errorCode":"ERROR_ZERO_BALANCE","errorDescription":"no money"}`)

	_, err := client.CreateTask(context.Background(), request())
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Code != "ERROR_ZERO_BALANCE" {
		t.Fatalf("want APIError ERROR_ZERO_BALANCE, got %v", err)
	}
}

func TestGetResultAPIErrorIsReported(t *testing.T) {
	api, client := newFakeAPI(t)
	api.on("createTask", 200, `{"errorId":0,"taskId":1}`).
		on("getTaskResult", 200, `{"errorId":12,"errorCode":"ERROR_CAPTCHA_UNSOLVABLE","errorDescription":"nope"}`)

	_, err := client.Solve(context.Background(), request())
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Code != "ERROR_CAPTCHA_UNSOLVABLE" || apiErr.Method != "getTaskResult" {
		t.Fatalf("want APIError ERROR_CAPTCHA_UNSOLVABLE, got %v", err)
	}
}

func TestHTTP500And503WhilePollingMeanStillProcessing(t *testing.T) {
	api, client := newFakeAPI(t)
	api.on("getTaskResult", 500, "Internal Server Error").on("getTaskResult", 503, "")

	for i := 0; i < 2; i++ {
		res, err := client.GetTaskResult(context.Background(), 7)
		if err != nil || res.Ready {
			t.Fatalf("call %d: want processing, got %+v %v", i, res, err)
		}
	}
}

func TestOtherHTTPStatusIsAnError(t *testing.T) {
	api, client := newFakeAPI(t)
	api.on("createTask", 502, "<html>bad gateway</html>")

	_, err := client.CreateTask(context.Background(), request())
	var statusErr *StatusError
	if !errors.As(err, &statusErr) || statusErr.Status != 502 || statusErr.Method != "createTask" {
		t.Fatalf("want StatusError 502, got %v", err)
	}
}

func TestPollingGivesUpWithTimeout(t *testing.T) {
	api, client := newFakeAPI(t)
	api.on("createTask", 200, `{"errorId":0,"taskId":5}`).on("getTaskResult", 200, `{"errorId":0,"status":"processing"}`)

	started := time.Now()
	_, err := client.Solve(context.Background(), request())
	var timeoutErr *TimeoutError
	if !errors.Is(err, ErrTimeout) || !errors.As(err, &timeoutErr) || timeoutErr.TaskID != 5 {
		t.Fatalf("want TimeoutError for task 5, got %v", err)
	}
	if time.Since(started) > time.Second {
		t.Errorf("took %s", time.Since(started))
	}
}

func TestContextCancellationStopsPolling(t *testing.T) {
	api, client := newFakeAPI(t)
	api.on("createTask", 200, `{"errorId":0,"taskId":5}`).on("getTaskResult", 200, `{"errorId":0,"status":"processing"}`)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	if _, err := client.Solve(ctx, request()); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("want context.DeadlineExceeded, got %v", err)
	}
}

func TestInvalidRequestIsRejectedBeforeAnyHTTPCall(t *testing.T) {
	api, client := newFakeAPI(t)

	_, err := client.Solve(context.Background(), RecaptchaV2Request{WebsiteKey: "K"})
	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("want ValidationError, got %v", err)
	}
	if len(api.bodies) != 0 {
		t.Errorf("unexpected HTTP calls: %v", api.bodies)
	}
}

func TestGetBalance(t *testing.T) {
	api, client := newFakeAPI(t)
	api.on("getBalance", 200, `{"errorId":0,"balance":12.5}`)

	balance, err := client.GetBalance(context.Background())
	if err != nil || balance != 12.5 {
		t.Fatalf("got %v %v", balance, err)
	}
	assertJSONEqual(t, api.bodies["getBalance"][0], `{"clientKey":"KEY"}`)
}

func TestGetBalanceError(t *testing.T) {
	api, client := newFakeAPI(t)
	api.on("getBalance", 200, `{"errorId":1,"errorCode":"ERROR_KEY_DOES_NOT_EXIST","errorDescription":"bad key"}`)

	_, err := client.GetBalance(context.Background())
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Method != "getBalance" {
		t.Fatalf("want APIError from getBalance, got %v", err)
	}
}

func TestMalformedResponse(t *testing.T) {
	api, client := newFakeAPI(t)
	api.on("getBalance", 200, "not json")

	_, err := client.GetBalance(context.Background())
	var malformed *MalformedResponseError
	if !errors.As(err, &malformed) {
		t.Fatalf("want MalformedResponseError, got %v", err)
	}
}

func TestDefaultScheduleComesFromTheRequest(t *testing.T) {
	api, _ := newFakeAPI(t)
	srv := httptest.NewServer(api)
	defer srv.Close()
	client := New("KEY", WithServiceURL(srv.URL))
	api.on("createTask", 200, `{"errorId":0,"taskId":1}`).on("getTaskResult", 200, `{"errorId":0,"status":"ready","solution":{"text":"abc"}}`)

	// ImageToText: first poll after 0.35 s
	started := time.Now()
	solution, err := client.Solve(context.Background(), ImageToTextRequest{Body: []byte("ABC")})
	if err != nil || solution.String("text") != "abc" {
		t.Fatalf("got %s %v", solution, err)
	}
	if elapsed := time.Since(started); elapsed < 300*time.Millisecond {
		t.Errorf("first poll after %s, want >= 350ms", elapsed)
	}
}

func TestOptions(t *testing.T) {
	if _, err := NewWithOptions("KEY", WithClientProxy("::not a url")); err == nil {
		t.Error("invalid client proxy URL accepted")
	}
	if _, err := NewWithOptions("KEY", WithClientProxy("socks5://user:pass@127.0.0.1:1080"), WithClientTimeout(time.Second), WithSoftID(1)); err != nil {
		t.Error(err)
	}
	if _, err := NewWithOptions("KEY", WithHTTPClient(nil)); err == nil {
		t.Error("nil http client accepted")
	}
}
