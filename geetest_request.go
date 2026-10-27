package capmonster

// GeetestV3Request is GeeTest v3 (GeeTestTask, version 3).
type GeetestV3Request struct {
	// WebsiteURL is the page where the captcha is initialized (usually the Referer of the gettype.php request,
	// which may differ from the URL in the address bar).
	WebsiteURL string `json:"websiteURL"`
	GT         string `json:"gt"` // the gt key of the domain
	// Challenge is a dynamic key: obtain a fresh one before creating each task and before the captcha is
	// initialized on the page (afterwards it becomes invalid).
	Challenge string `json:"challenge"`
	Proxy     *Proxy `json:"-"`
	// GeetestAPIServerSubdomain must differ from api.geetest.com; may be required for some sites.
	GeetestAPIServerSubdomain string `json:"geetestApiServerSubdomain,omitempty"`
	// GeetestGetLib is the path to the captcha script, as a JSON string; may be required for some sites.
	GeetestGetLib string `json:"geetestGetLib,omitempty"`
	UserAgent     string `json:"userAgent,omitempty"`
}

func (r GeetestV3Request) Timeouts() Timeouts { return timeouts(1, 1, 80) }

func (r GeetestV3Request) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "gt", r.GT, "challenge", r.Challenge); err != nil {
		return nil, err
	}
	task, err := buildTask("GeeTestTask", "", r, r.Proxy)
	if err == nil {
		task["version"] = 3
	}
	return task, err
}

// GeetestV4Request is GeeTest v4 (GeeTestTask, version 4).
type GeetestV4Request struct {
	WebsiteURL string `json:"websiteURL"`
	GT         string `json:"gt"` // the captcha_id
	// InitParameters are extra v4 parameters, e.g. {"riskType": "slide"}.
	InitParameters            map[string]any `json:"initParameters,omitempty"`
	Proxy                     *Proxy         `json:"-"`
	GeetestAPIServerSubdomain string         `json:"geetestApiServerSubdomain,omitempty"`
	GeetestGetLib             string         `json:"geetestGetLib,omitempty"` // JSON string
	UserAgent                 string         `json:"userAgent,omitempty"`
}

func (r GeetestV4Request) Timeouts() Timeouts { return timeouts(1, 1, 80) }

func (r GeetestV4Request) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "gt", r.GT); err != nil {
		return nil, err
	}
	task, err := buildTask("GeeTestTask", "", r, r.Proxy)
	if err == nil {
		task["version"] = 4
	}
	return task, err
}
