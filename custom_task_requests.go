package capmonster

import "strings"

// DataDomeVersionNew is the DataDomeMetadata.DatadomeVersion of the updated solving method (i.js and c.js).
const DataDomeVersionNew = "new"

// DataDomeMetadata values are dynamic: extract them right before creating the task.
type DataDomeMetadata struct {
	// CaptchaURL is the captcha URL with the latest dynamic parameters, usually
	// https://geo.captcha-delivery.com/captcha/?initialCid=...&t=fe&... Submit it only if its t parameter is fe
	// (t=bv means the proxy is blocked).
	CaptchaURL     string `json:"captchaUrl"`
	DatadomeCookie string `json:"datadomeCookie"` // your DataDome cookies (datadome=...)
	// DatadomeVersion is DataDomeVersionNew or empty (legacy method, only i.js).
	DatadomeVersion string `json:"datadomeVersion,omitempty"`
}

// DataDomeCustomTaskRequest is DataDome (CustomTask, class DataDome). Your own Proxy is required; if it uses IP
// authorization, whitelist 65.21.190.34. The solution contains cookies that must be set in the browser.
type DataDomeCustomTaskRequest struct {
	WebsiteURL string           `json:"websiteURL"`
	Metadata   DataDomeMetadata `json:"metadata"`
	Proxy      *Proxy           `json:"-"`
	UserAgent  string           `json:"userAgent,omitempty"`
}

func (r DataDomeCustomTaskRequest) Timeouts() Timeouts { return timeouts(1, 3, 180) }

func (r DataDomeCustomTaskRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "metadata.captchaUrl", r.Metadata.CaptchaURL,
		"metadata.datadomeCookie", r.Metadata.DatadomeCookie); err != nil {
		return nil, err
	}
	if err := requireProxy(r.Proxy); err != nil {
		return nil, err
	}
	return buildTask("CustomTask", "DataDome", r, r.Proxy)
}

// ImpervaMetadata values are dynamic: extract them right before creating the task.
type ImpervaMetadata struct {
	IncapsulaScriptURL string `json:"incapsulaScriptUrl"` // e.g. _Incapsula_Resource?SWJIYLWA=719d34d3...
	IncapsulaCookies   string `json:"incapsulaCookies"`   // e.g. "incap_sess_*=...; visid_incap_*=..."
	// Reese84URLEndpoint is the endpoint the reese84 fingerprint is sent to (ends with ?d=site.com).
	Reese84URLEndpoint string `json:"reese84UrlEndpoint,omitempty"`
}

// ImpervaCustomTaskRequest is Imperva / Incapsula (CustomTask, class Imperva). Your own Proxy is required.
type ImpervaCustomTaskRequest struct {
	WebsiteURL string          `json:"websiteURL"`
	Metadata   ImpervaMetadata `json:"metadata"`
	Proxy      *Proxy          `json:"-"`
	UserAgent  string          `json:"userAgent,omitempty"`
}

func (r ImpervaCustomTaskRequest) Timeouts() Timeouts { return timeouts(1, 1, 15) }

func (r ImpervaCustomTaskRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "metadata.incapsulaScriptUrl", r.Metadata.IncapsulaScriptURL,
		"metadata.incapsulaCookies", r.Metadata.IncapsulaCookies); err != nil {
		return nil, err
	}
	if err := requireProxy(r.Proxy); err != nil {
		return nil, err
	}
	return buildTask("CustomTask", "Imperva", r, r.Proxy)
}

// TspdMetadata values are dynamic: extract them right before creating the task.
type TspdMetadata struct {
	TspdCookie     string `json:"tspdCookie"`     // cookies from the TSPD challenge page, e.g. "TS386a400d029=0826...; ...;"
	HTMLPageBase64 string `json:"htmlPageBase64"` // the entire TSPD page encoded in base64
}

// TspdCustomTaskRequest is TSPD (CustomTask, class tspd). Your own Proxy is required; its login must not contain
// the characters ';', '|', ')'. UserAgent should match the one used when the cookie and the page were obtained.
type TspdCustomTaskRequest struct {
	WebsiteURL string       `json:"websiteURL"`
	Metadata   TspdMetadata `json:"metadata"`
	Proxy      *Proxy       `json:"-"`
	UserAgent  string       `json:"userAgent,omitempty"`
}

func (r TspdCustomTaskRequest) Timeouts() Timeouts { return timeouts(1, 1, 60) }

func (r TspdCustomTaskRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "metadata.tspdCookie", r.Metadata.TspdCookie,
		"metadata.htmlPageBase64", r.Metadata.HTMLPageBase64); err != nil {
		return nil, err
	}
	if err := requireProxy(r.Proxy); err != nil {
		return nil, err
	}
	if strings.ContainsAny(r.Proxy.Login, ";|)") {
		return nil, &ValidationError{Field: "proxyLogin", Reason: "must not contain the characters ';', '|', ')' for TSPD"}
	}
	return buildTask("CustomTask", "tspd", r, r.Proxy)
}

// HuntMetadata: set APIGetLib only to generate a fingerprint, or add exactly one of Data / WidgetURL to solve
// the captcha. Values are dynamic: extract them right before creating the task.
type HuntMetadata struct {
	// APIGetLib is the full URL of the api.js file (look for hd-api or api.js in the Network tab), e.g.
	// https://example.com/hd-api/external/apps/<app-hash>/api.js.
	APIGetLib string `json:"apiGetLib"`
	Data      string `json:"data,omitempty"` // the meta.token value received from the website
	// WidgetURL is the full URL of the HUNT widget, e.g. https://captcha.example.com/widget?hash=<widget-hash>.
	WidgetURL string `json:"widgetUrl,omitempty"`
}

// HuntCustomTaskRequest is HUNT captcha (CustomTask, class HUNT). Your own Proxy is required.
type HuntCustomTaskRequest struct {
	WebsiteURL string       `json:"websiteURL"`
	Metadata   HuntMetadata `json:"metadata"`
	Proxy      *Proxy       `json:"-"`
	UserAgent  string       `json:"userAgent,omitempty"`
}

func (r HuntCustomTaskRequest) Timeouts() Timeouts { return timeouts(1, 3, 180) }

func (r HuntCustomTaskRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "metadata.apiGetLib", r.Metadata.APIGetLib); err != nil {
		return nil, err
	}
	if r.Metadata.Data != "" && r.Metadata.WidgetURL != "" {
		return nil, &ValidationError{Field: "metadata", Reason: "data and widgetUrl are mutually exclusive"}
	}
	if err := requireProxy(r.Proxy); err != nil {
		return nil, err
	}
	return buildTask("CustomTask", "HUNT", r, r.Proxy)
}

// TenDiMetadata holds the optional TenDI parameters.
type TenDiMetadata struct {
	// CaptchaURL is the captcha script, usually TCaptcha.js / TCaptcha-global.js (TJNCaptcha*.js / TJCaptcha*.js
	// for the second version).
	CaptchaURL string `json:"captchaUrl,omitempty"`
	// AidEncrypted is the Aid value from the target website; single-use, fetch a new one for every task.
	// Omit it if the site does not return an Aid.
	AidEncrypted string `json:"aidEncrypted,omitempty"`
}

// TenDiCustomTaskRequest is TenDI / Tencent captcha (CustomTask, class TenDI).
type TenDiCustomTaskRequest struct {
	WebsiteURL string         `json:"websiteURL"`
	WebsiteKey string         `json:"websiteKey"` // captchaAppId, the aid value (e.g. "189123456"), unique per website
	Metadata   *TenDiMetadata `json:"metadata,omitempty"`
	Proxy      *Proxy         `json:"-"`
	UserAgent  string         `json:"userAgent,omitempty"`
}

func (r TenDiCustomTaskRequest) Timeouts() Timeouts { return timeouts(1, 3, 180) }

func (r TenDiCustomTaskRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey); err != nil {
		return nil, err
	}
	return buildTask("CustomTask", "TenDI", r, r.Proxy)
}

// AlibabaMetadata: only SceneID and Prefix are always required; set the others only if the website uses them.
type AlibabaMetadata struct {
	SceneID string `json:"sceneId"` // captcha scenario identifier, e.g. "1ww7426c"
	// Prefix is the subdomain of the request URL that loads the task, e.g. "dlw3kug" for
	// https://dlw3kug.captcha-open.example.aliyuncs.com/.
	Prefix        string `json:"prefix"`
	UserID        string `json:"userId,omitempty"`        // unique user/session identifier on the website side
	UserUserID    string `json:"userUserId,omitempty"`    // additional (secondary) user identifier
	VerifyType    string `json:"verifyType,omitempty"`    // e.g. "1.0"
	Region        string `json:"region,omitempty"`        // e.g. "sgp" (cn gives WEB# tokens, sgp gives SG_WEB#)
	UserCertifyID string `json:"UserCertifyId,omitempty"` // verification ID of the current captcha session
	APIGetLib     string `json:"apiGetLib,omitempty"`     // captcha JS library (may change on every page render)
	// PunishURL is the full URL (with query) of the verification page, for websites where Alibaba returns a
	// separate /punish URL.
	PunishURL      string `json:"punishUrl,omitempty"`
	CookieRequired *bool  `json:"cookieRequired,omitempty"` // return cookies too (use true for Alibaba WAF sites)
}

// AlibabaCustomTaskRequest is Alibaba CAPTCHA / WAF (CustomTask, class alibaba).
type AlibabaCustomTaskRequest struct {
	WebsiteURL string          `json:"websiteURL"`
	Metadata   AlibabaMetadata `json:"metadata"`
	Proxy      *Proxy          `json:"-"`
	UserAgent  string          `json:"userAgent,omitempty"`
}

func (r AlibabaCustomTaskRequest) Timeouts() Timeouts { return timeouts(1, 3, 120) }

func (r AlibabaCustomTaskRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "metadata.sceneId", r.Metadata.SceneID,
		"metadata.prefix", r.Metadata.Prefix); err != nil {
		return nil, err
	}
	return buildTask("CustomTask", "alibaba", r, r.Proxy)
}

// FriendlyMetadata holds the Friendly Captcha parameters.
type FriendlyMetadata struct {
	// APIGetLib is the URL of the JS file, used to detect the captcha version.
	// V1: https://cdn.jsdelivr.net/npm/friendly-challenge@X.Y.Z/widget.module.min.js (X.Y.Z from the
	// x-frc-client header). V2: the site.min.js loaded on the page, e.g.
	// https://cdn.jsdelivr.net/npm/@friendlycaptcha/sdk@X.Y.Z/site.min.js.
	APIGetLib string `json:"apiGetLib"`
}

// FriendlyCustomTaskRequest is Friendly Captcha (CustomTask, class friendly).
type FriendlyCustomTaskRequest struct {
	WebsiteURL string           `json:"websiteURL"`
	WebsiteKey string           `json:"websiteKey"`
	Metadata   FriendlyMetadata `json:"metadata"`
	Proxy      *Proxy           `json:"-"`
	UserAgent  string           `json:"userAgent,omitempty"`
}

func (r FriendlyCustomTaskRequest) Timeouts() Timeouts { return timeouts(1, 3, 180) }

func (r FriendlyCustomTaskRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey,
		"metadata.apiGetLib", r.Metadata.APIGetLib); err != nil {
		return nil, err
	}
	return buildTask("CustomTask", "friendly", r, r.Proxy)
}

// AltchaMetadata holds the challenge parameters taken from the page (all four are required).
type AltchaMetadata struct {
	Challenge  string `json:"challenge"`
	Iterations string `json:"iterations"`
	Salt       string `json:"salt"`
	Signature  string `json:"signature"`
}

// AltchaCustomTaskRequest is Altcha proof-of-work captcha (CustomTask, class altcha).
type AltchaCustomTaskRequest struct {
	WebsiteURL string         `json:"websiteURL"`
	WebsiteKey string         `json:"websiteKey"` // an empty string is allowed
	Metadata   AltchaMetadata `json:"metadata"`
	Proxy      *Proxy         `json:"-"`
	UserAgent  string         `json:"userAgent,omitempty"`
}

func (r AltchaCustomTaskRequest) Timeouts() Timeouts { return timeouts(1, 1, 50) }

func (r AltchaCustomTaskRequest) Task() (map[string]any, error) {
	m := r.Metadata
	if err := required("websiteURL", r.WebsiteURL, "metadata.challenge", m.Challenge, "metadata.iterations", m.Iterations,
		"metadata.salt", m.Salt, "metadata.signature", m.Signature); err != nil {
		return nil, err
	}
	return buildTask("CustomTask", "altcha", r, r.Proxy)
}
