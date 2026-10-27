package capmonster

import "fmt"

// RecaptchaV3ProxylessRequest is reCAPTCHA v3 (RecaptchaV3TaskProxyless). Proxies are not supported.
type RecaptchaV3ProxylessRequest struct {
	WebsiteURL string `json:"websiteURL"`
	WebsiteKey string `json:"websiteKey"` // ...api.js?render=THIS_ONE
	// MinScore is the minimum score of the token, 0.1..0.9 (0 = not sent).
	MinScore     float64 `json:"minScore,omitempty"`
	PageAction   string  `json:"pageAction,omitempty"`   // widget action, default "verify"
	IsEnterprise *bool   `json:"isEnterprise,omitempty"` // solve as reCAPTCHA v3 Enterprise
}

func (r RecaptchaV3ProxylessRequest) Timeouts() Timeouts { return timeouts(1, 3, 180) }

func (r RecaptchaV3ProxylessRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey); err != nil {
		return nil, err
	}
	if err := checkMinScore(r.MinScore); err != nil {
		return nil, err
	}
	return buildTask("RecaptchaV3TaskProxyless", "", r, nil)
}

// RecaptchaV3EnterpriseRequest is reCAPTCHA v3 Enterprise (RecaptchaV3EnterpriseTask). Proxies are not supported.
type RecaptchaV3EnterpriseRequest struct {
	WebsiteURL string  `json:"websiteURL"`
	WebsiteKey string  `json:"websiteKey"`           // ...enterprise.js?render=THIS_ONE
	MinScore   float64 `json:"minScore,omitempty"`   // 0.1..0.9 (0 = not sent)
	PageAction string  `json:"pageAction,omitempty"` // if it differs from the default "verify"
}

func (r RecaptchaV3EnterpriseRequest) Timeouts() Timeouts { return timeouts(1, 3, 180) }

func (r RecaptchaV3EnterpriseRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey); err != nil {
		return nil, err
	}
	if err := checkMinScore(r.MinScore); err != nil {
		return nil, err
	}
	return buildTask("RecaptchaV3EnterpriseTask", "", r, nil)
}

func checkMinScore(v float64) error {
	if v != 0 && (v < 0.1 || v > 0.9) {
		return &ValidationError{Field: "minScore", Reason: fmt.Sprintf("must be between 0.1 and 0.9, got %v", v)}
	}
	return nil
}
