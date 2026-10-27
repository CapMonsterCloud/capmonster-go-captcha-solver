package capmonster

// BasiliskCustomTaskRequest is Basilisk captcha (CustomTask, class Basilisk).
type BasiliskCustomTaskRequest struct {
	WebsiteURL string `json:"websiteURL"`
	// WebsiteKey is the data-sitekey attribute of the captcha container, or the site_key field of the POST
	// payload to https://basiliskcaptcha.com/challenge/check-site.
	WebsiteKey string `json:"websiteKey"`
	Proxy      *Proxy `json:"-"`
	UserAgent  string `json:"userAgent,omitempty"`
}

func (r BasiliskCustomTaskRequest) Timeouts() Timeouts { return timeouts(1, 3, 100) }

func (r BasiliskCustomTaskRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey); err != nil {
		return nil, err
	}
	return buildTask("CustomTask", "Basilisk", r, r.Proxy)
}

// ProsopoTaskRequest is Prosopo Procaptcha (ProsopoTask), token based.
type ProsopoTaskRequest struct {
	WebsiteURL string `json:"websiteURL"`
	WebsiteKey string `json:"websiteKey"` // the siteKey found on the page
	Proxy      *Proxy `json:"-"`
	UserAgent  string `json:"userAgent,omitempty"`
}

func (r ProsopoTaskRequest) Timeouts() Timeouts { return timeouts(1, 3, 180) }

func (r ProsopoTaskRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey); err != nil {
		return nil, err
	}
	return buildTask("ProsopoTask", "", r, r.Proxy)
}

// MTCaptchaRequest is MTCaptcha (MTCaptchaTask).
type MTCaptchaRequest struct {
	WebsiteURL  string `json:"websiteURL"`
	WebsiteKey  string `json:"websiteKey"` // passed on the page as sk; typically MTPublic-abCDEFJAB
	Proxy       *Proxy `json:"-"`
	PageAction  string `json:"pageAction,omitempty"`  // the act parameter, only if it differs from the default %24
	IsInvisible *bool  `json:"isInvisible,omitempty"` // the captcha has a hidden confirmation field
	UserAgent   string `json:"userAgent,omitempty"`
}

func (r MTCaptchaRequest) Timeouts() Timeouts { return timeouts(1, 3, 180) }

func (r MTCaptchaRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey); err != nil {
		return nil, err
	}
	return buildTask("MTCaptchaTask", "", r, r.Proxy)
}

// YidunRequest is NetEase Yidun captcha (YidunTask). Setting Challenge selects the Enterprise variant.
type YidunRequest struct {
	WebsiteURL string `json:"websiteURL"`
	WebsiteKey string `json:"websiteKey"` // the siteKey found on the page
	Proxy      *Proxy `json:"-"`
	UserAgent  string `json:"userAgent,omitempty"`
	// YidunGetLib is the full URL of the JS file that loads the captcha; recommended when Challenge, Hcg or Hct
	// are present on the site.
	YidunGetLib string `json:"yidunGetLib,omitempty"`
	// YidunAPIServerSubdomain is the API server subdomain without https://; must differ from c.dun.163.com /
	// c-v6.dun.163.com. Needed for custom (often Enterprise) servers.
	YidunAPIServerSubdomain string `json:"yidunApiServerSubdomain,omitempty"`
	Challenge               string `json:"challenge,omitempty"` // unique id of the current captcha (Enterprise)
	Hcg                     string `json:"hcg,omitempty"`       // captcha hash, obtained together with Challenge
	Hct                     int64  `json:"hct,omitempty"`       // numeric timestamp used in Enterprise validation
}

func (r YidunRequest) Timeouts() Timeouts { return timeouts(1, 3, 180) }

func (r YidunRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey); err != nil {
		return nil, err
	}
	return buildTask("YidunTask", "", r, r.Proxy)
}

// BinanceTaskRequest is Binance captcha (BinanceTask). Use only to log in with your own account.
// ValidateID is dynamic: extract it right before creating the task.
type BinanceTaskRequest struct {
	WebsiteURL string `json:"websiteURL"`
	// WebsiteKey is the section of the site: bizId, bizType or bizCode from the traffic, e.g. "login".
	WebsiteKey string `json:"websiteKey"`
	// ValidateID is validateId, securityId or securityCheckResponseValidateId from the traffic.
	ValidateID string `json:"validateId"`
	Proxy      *Proxy `json:"-"`
	UserAgent  string `json:"userAgent,omitempty"`
}

func (r BinanceTaskRequest) Timeouts() Timeouts { return timeouts(1, 1, 20) }

func (r BinanceTaskRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websiteKey", r.WebsiteKey, "validateId", r.ValidateID); err != nil {
		return nil, err
	}
	return buildTask("BinanceTask", "", r, r.Proxy)
}

// FuncaptchaRequest is FunCaptcha / Arkose Labs (FunCaptchaTask). Your own Proxy is required: use the same
// proxy to obtain the captcha parameters and to create the task.
type FuncaptchaRequest struct {
	WebsiteURL       string `json:"websiteURL"`
	WebsitePublicKey string `json:"websitePublicKey"` // pk, e.g. EX72CCFB-26EX-40E5-91E6-85EX70BE98ED
	Proxy            *Proxy `json:"-"`
	// FuncaptchaAPIJSSubdomain is the Arkose Labs subdomain (surl) without https://, only if it differs from
	// client-api.arkoselabs.com.
	FuncaptchaAPIJSSubdomain string `json:"funcaptchaApiJSSubdomain,omitempty"`
	Data                     string `json:"data,omitempty"` // required if the site uses data[blob], e.g. {"blob":"..."}
	Cookies                  string `json:"cookies,omitempty"`
	UserAgent                string `json:"userAgent,omitempty"`
}

func (r FuncaptchaRequest) Timeouts() Timeouts { return timeouts(1, 1, 80) }

func (r FuncaptchaRequest) Task() (map[string]any, error) {
	if err := required("websiteURL", r.WebsiteURL, "websitePublicKey", r.WebsitePublicKey); err != nil {
		return nil, err
	}
	if err := requireProxy(r.Proxy); err != nil {
		return nil, err
	}
	return buildTask("FunCaptchaTask", "", r, r.Proxy)
}
