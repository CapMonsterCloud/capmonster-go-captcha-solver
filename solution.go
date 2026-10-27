package capmonster

import "encoding/json"

// Solution is the raw `solution` object of a solved task. Its shape depends on the captcha type
// (gRecaptchaResponse, token, text, cookies, ...): read one string field with String or decode it into
// your own struct with Decode.
type Solution json.RawMessage

// String returns a top-level string field of the solution, or "" if it is missing or not a string.
func (s Solution) String(key string) string {
	var m map[string]json.RawMessage
	if json.Unmarshal(s, &m) != nil {
		return ""
	}
	var v string
	if json.Unmarshal(m[key], &v) != nil {
		return ""
	}
	return v
}

// Decode unmarshals the solution into v (a pointer to a struct or a map).
func (s Solution) Decode(v any) error { return json.Unmarshal(s, v) }

// MarshalJSON keeps the solution as-is when it is re-encoded.
func (s Solution) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte("null"), nil
	}
	return s, nil
}
