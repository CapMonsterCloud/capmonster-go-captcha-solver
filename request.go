package capmonster

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Request is a task that can be sent to CapMonster Cloud. All *Request types of this package implement it;
// implement it yourself only for a task type the SDK does not know yet.
type Request interface {
	// Timeouts is the polling schedule recommended for this captcha type.
	Timeouts() Timeouts
	// Task validates the request and returns the `task` object for createTask.
	Task() (map[string]any, error)
}

// Timeouts is the polling schedule of getTaskResult.
type Timeouts struct {
	FirstRequestDelay time.Duration // pause between createTask and the first getTaskResult
	RequestsInterval  time.Duration // pause between two getTaskResult calls
	Timeout           time.Duration // total waiting time before a *TimeoutError
}

func timeouts(first, interval, total float64) Timeouts {
	sec := func(v float64) time.Duration { return time.Duration(v * float64(time.Second)) }
	return Timeouts{FirstRequestDelay: sec(first), RequestsInterval: sec(interval), Timeout: sec(total)}
}

// Bool returns a pointer to v, for optional *bool fields.
func Bool(v bool) *bool { return &v }

// Int returns a pointer to v, for optional *int fields.
func Int(v int) *int { return &v }

// buildTask turns a request struct (exported fields with json tags, empty optional fields `omitempty`)
// into the task object: adds `type`, `class` (CustomTask / ComplexImageTask only) and the proxy fields.
func buildTask(taskType, class string, fields any, proxy *Proxy) (map[string]any, error) {
	raw, err := json.Marshal(fields)
	if err != nil {
		return nil, fmt.Errorf("capmonster: encode task: %w", err)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber() // keep integers such as hct exact
	task := map[string]any{}
	if err := dec.Decode(&task); err != nil {
		return nil, fmt.Errorf("capmonster: encode task: %w", err)
	}
	task["type"] = taskType
	if class != "" {
		task["class"] = class
	}
	if proxy != nil {
		if err := proxy.validate(); err != nil {
			return nil, err
		}
		for k, v := range proxy.fields() {
			task[k] = v
		}
	}
	return task, nil
}

// required checks pairs of (API field name, value) and fails on the first empty value.
func required(nameValues ...string) error {
	for i := 0; i+1 < len(nameValues); i += 2 {
		if nameValues[i+1] == "" {
			return &ValidationError{Field: nameValues[i], Reason: "must not be empty"}
		}
	}
	return nil
}

func requireProxy(p *Proxy) error {
	if p == nil {
		return &ValidationError{Field: "proxy", Reason: "is required for this task type"}
	}
	return nil
}
