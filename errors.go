package capmonster

import (
	"errors"
	"fmt"
	"time"
)

// ErrTimeout matches a *TimeoutError with errors.Is.
var ErrTimeout = errors.New("capmonster: solution timeout")

// ValidationError means the request is inconsistent (empty required field, value out of range...).
// It is returned locally, before anything is sent.
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("capmonster: invalid request: %s %s", e.Field, e.Reason)
}

// APIError means the API returned errorId != 0. Code is the stable identifier to branch on, e.g.
// ERROR_KEY_DOES_NOT_EXIST, ERROR_ZERO_BALANCE, ERROR_CAPTCHA_UNSOLVABLE.
type APIError struct {
	Method      string // createTask, getTaskResult or getBalance
	Code        string
	Description string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("capmonster: %s failed: [%s] %s", e.Method, e.Code, e.Description)
}

// StatusError means the API answered with an unexpected HTTP status and no error JSON.
type StatusError struct {
	Method string
	Status int
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("capmonster: unexpected HTTP status %d from %s", e.Status, e.Method)
}

// MalformedResponseError means the API answered with something that is not the expected JSON.
type MalformedResponseError struct {
	Method string
	Reason string
}

func (e *MalformedResponseError) Error() string {
	return fmt.Sprintf("capmonster: malformed %s response: %s", e.Method, e.Reason)
}

// TimeoutError means the solution did not arrive within the maximum waiting time of the task type.
type TimeoutError struct {
	TaskID  int64
	Timeout time.Duration
}

func (e *TimeoutError) Error() string {
	return fmt.Sprintf("capmonster: failed to get a solution of task %d within %s", e.TaskID, e.Timeout)
}

// Is makes errors.Is(err, ErrTimeout) work.
func (e *TimeoutError) Is(target error) bool { return target == ErrTimeout }
