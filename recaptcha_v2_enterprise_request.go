package capmonster

// RecaptchaV2EnterpriseRequest is reCAPTCHA v2 Enterprise (RecaptchaV2EnterpriseTask). Without Proxy the task is
// solved from CapMonster's own IPs; a proxy is only needed when the website does not accept the token.
type RecaptchaV2EnterpriseRequest struct {
	WebsiteURL string `json:"websiteURL"`
	WebsiteKey string `json:"websiteKey"`
	Proxy      *Proxy `json:"-"`
	// EnterprisePayload holds additional Enterprise parameters, e.g. {"s": "..."} (the s value passed to
	// grecaptcha.enterprise.render).
	EnterprisePayload map[string]string `json:"enterprisePayload,omitempty"`
	APIDomain         string            `json:"apiDomain,omitempty"`  // www.google.com or www.recaptcha.net
	PageAction        string            `json:"pageAction,omitempty"` // if it differs from the default "verify"
	UserAgent         string            `json:"userAgent,omitempty"`
	Cookies           string            `json:"cookies,omitempty"`
}

func (r RecaptchaV2EnterpriseRequest) Timeouts() Timeouts { return timeouts(1, 3, 180) }

func (r RecaptchaV2EnterpriseRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey); err != nil {
		return nil, err
	}
	return buildTask("RecaptchaV2EnterpriseTask", "", r, r.Proxy)
}
