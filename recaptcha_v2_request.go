package capmonster

// RecaptchaV2Request is reCAPTCHA v2 (RecaptchaV2Task). Without Proxy the task is solved from CapMonster's own
// IPs. Solution: gRecaptchaResponse.
type RecaptchaV2Request struct {
	WebsiteURL string `json:"websiteURL"` // URL of the page with the captcha
	WebsiteKey string `json:"websiteKey"` // reCAPTCHA site key
	Proxy      *Proxy `json:"-"`
	// DataSValue is the one-time data-s token of some custom reCAPTCHA v2 implementations.
	DataSValue  string `json:"recaptchaDataSValue,omitempty"`
	UserAgent   string `json:"userAgent,omitempty"` // current browser User-Agent (Windows)
	Cookies     string `json:"cookies,omitempty"`   // "name1=value1; name2=value2"
	IsInvisible *bool  `json:"isInvisible,omitempty"`
}

func (r RecaptchaV2Request) Timeouts() Timeouts { return timeouts(1, 3, 180) }

func (r RecaptchaV2Request) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey); err != nil {
		return nil, err
	}
	return buildTask("RecaptchaV2Task", "", r, r.Proxy)
}
