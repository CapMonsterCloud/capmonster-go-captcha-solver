package capmonster

// Amazon AWS WAF (AmazonTask): three mutually exclusive parameter sets, one request type each.
// CookieSolution: false (default) returns captcha_voucher and existing_token, true returns the aws-waf-token cookie.

// AmazonWafVisibleRequest is option 1: the visible AWS WAF captcha (jsapi.js / captcha.js only).
type AmazonWafVisibleRequest struct {
	WebsiteURL     string `json:"websiteURL"`
	WebsiteKey     string `json:"websiteKey"`    // value of the apiKey field used when rendering the captcha
	CaptchaScript  string `json:"captchaScript"` // link to jsapi.js (<Integration URL>jsapi.js)
	Proxy          *Proxy `json:"-"`
	UserAgent      string `json:"userAgent,omitempty"` // only this option accepts a User-Agent
	CookieSolution *bool  `json:"cookieSolution,omitempty"`
}

func (r AmazonWafVisibleRequest) Timeouts() Timeouts { return amazonTimeouts() }

func (r AmazonWafVisibleRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey, "captchaScript", r.CaptchaScript); err != nil {
		return nil, err
	}
	return buildTask("AmazonTask", "", r, r.Proxy)
}

// AmazonWafChallengeRequest is option 2: challenge plus captcha. The values come from window.gokuProps on the
// 405 page of the protected site.
type AmazonWafChallengeRequest struct {
	WebsiteURL      string `json:"websiteURL"`
	ChallengeScript string `json:"challengeScript"` // link to challenge.js
	WebsiteKey      string `json:"websiteKey"`      // window.gokuProps.key
	Context         string `json:"context"`         // window.gokuProps.context
	IV              string `json:"iv"`              // window.gokuProps.iv
	// CaptchaScript is the link to captcha.js; may be absent if only a challenge is present.
	CaptchaScript  string `json:"captchaScript,omitempty"`
	Proxy          *Proxy `json:"-"`
	CookieSolution *bool  `json:"cookieSolution,omitempty"`
}

func (r AmazonWafChallengeRequest) Timeouts() Timeouts { return amazonTimeouts() }

func (r AmazonWafChallengeRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "challengeScript", r.ChallengeScript, "websiteKey", r.WebsiteKey,
		"context", r.Context, "iv", r.IV); err != nil {
		return nil, err
	}
	return buildTask("AmazonTask", "", r, r.Proxy)
}

// AmazonWafChallengeOnlyRequest is option 3: invisible captcha, only challenge.js is loaded. context and iv
// are sent as empty strings, as the API requires.
type AmazonWafChallengeOnlyRequest struct {
	WebsiteURL      string `json:"websiteURL"`
	ChallengeScript string `json:"challengeScript"`
	Proxy           *Proxy `json:"-"`
	CookieSolution  *bool  `json:"cookieSolution,omitempty"`
}

func (r AmazonWafChallengeOnlyRequest) Timeouts() Timeouts { return amazonTimeouts() }

func (r AmazonWafChallengeOnlyRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "challengeScript", r.ChallengeScript); err != nil {
		return nil, err
	}
	task, err := buildTask("AmazonTask", "", r, r.Proxy)
	if err == nil {
		task["context"], task["iv"] = "", ""
	}
	return task, err
}

func amazonTimeouts() Timeouts { return timeouts(1, 3, 180) }
