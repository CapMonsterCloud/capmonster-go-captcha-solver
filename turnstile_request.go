package capmonster

// Cloudflare protections (TurnstileTask), one request type per mode. WebsiteKey is the Turnstile site key,
// e.g. 0x4AAAAAAABUYP0XeMJF0xoy; for the cf_clearance and wait_room modes any string may be passed (e.g. "xxxxxx").

// TurnstileRequest is the regular Turnstile widget (manual, non-interactive and invisible subtypes are detected
// automatically). Solution: token.
type TurnstileRequest struct {
	WebsiteURL string `json:"websiteURL"`
	WebsiteKey string `json:"websiteKey"`
	Proxy      *Proxy `json:"-"`
	UserAgent  string `json:"userAgent,omitempty"`
	PageAction string `json:"pageAction,omitempty"` // the action found in the callback used when the captcha is loaded
	Data       string `json:"data,omitempty"`       // the cData value of the widget
}

func (r TurnstileRequest) Timeouts() Timeouts { return turnstileTimeouts() }

func (r TurnstileRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey); err != nil {
		return nil, err
	}
	return buildTask("TurnstileTask", "", r, r.Proxy)
}

// TurnstileChallengeTokenRequest is a Cloudflare Challenge solved to a token (cloudflareTaskType "token").
// The values are the parameters passed to window.turnstile.render(el, paramsObj).
type TurnstileChallengeTokenRequest struct {
	WebsiteURL string `json:"websiteURL"`
	WebsiteKey string `json:"websiteKey"`
	PageAction string `json:"pageAction"` // the action of the callback; usually "managed" or "non-interactive"
	UserAgent  string `json:"userAgent"`
	Data       string `json:"data"`     // the cData parameter
	PageData   string `json:"pageData"` // the chlPageData parameter
	Proxy      *Proxy `json:"-"`
	APIJSURL   string `json:"apiJsUrl,omitempty"` // link to the captcha script
}

func (r TurnstileChallengeTokenRequest) Timeouts() Timeouts { return turnstileTimeouts() }

func (r TurnstileChallengeTokenRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey, "pageAction", r.PageAction,
		"userAgent", r.UserAgent, "data", r.Data, "pageData", r.PageData); err != nil {
		return nil, err
	}
	return cloudflareTask("token", r, r.Proxy)
}

// TurnstileChallengeCookieRequest is a Cloudflare Challenge solved to the cf_clearance cookie
// (cloudflareTaskType "cf_clearance"). Your own Proxy is required.
type TurnstileChallengeCookieRequest struct {
	WebsiteURL     string `json:"websiteURL"`
	WebsiteKey     string `json:"websiteKey"`
	HTMLPageBase64 string `json:"htmlPageBase64"` // base64 of the "Just a moment" page returned with status 403
	UserAgent      string `json:"userAgent"`
	Proxy          *Proxy `json:"-"`
}

func (r TurnstileChallengeCookieRequest) Timeouts() Timeouts { return turnstileTimeouts() }

func (r TurnstileChallengeCookieRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey, "htmlPageBase64", r.HTMLPageBase64,
		"userAgent", r.UserAgent); err != nil {
		return nil, err
	}
	if err := requireProxy(r.Proxy); err != nil {
		return nil, err
	}
	return cloudflareTask("cf_clearance", r, r.Proxy)
}

// TurnstileWaitRoomRequest is the Cloudflare Waiting Room (cloudflareTaskType "wait_room"); the solution is
// cookies. Your own Proxy is required.
type TurnstileWaitRoomRequest struct {
	WebsiteURL string `json:"websiteURL"`
	WebsiteKey string `json:"websiteKey"`
	// HTMLPageBase64 is base64 of the page containing <title>Waiting Room powered by Cloudflare</title>.
	HTMLPageBase64 string `json:"htmlPageBase64"`
	UserAgent      string `json:"userAgent"`
	Proxy          *Proxy `json:"-"`
}

func (r TurnstileWaitRoomRequest) Timeouts() Timeouts { return turnstileTimeouts() }

func (r TurnstileWaitRoomRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey, "htmlPageBase64", r.HTMLPageBase64,
		"userAgent", r.UserAgent); err != nil {
		return nil, err
	}
	if err := requireProxy(r.Proxy); err != nil {
		return nil, err
	}
	return cloudflareTask("wait_room", r, r.Proxy)
}

func turnstileTimeouts() Timeouts { return timeouts(1, 1, 80) }

func cloudflareTask(cloudflareTaskType string, fields any, proxy *Proxy) (map[string]any, error) {
	task, err := buildTask("TurnstileTask", "", fields, proxy)
	if err != nil {
		return nil, err
	}
	task["cloudflareTaskType"] = cloudflareTaskType
	return task, nil
}
