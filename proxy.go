package capmonster

import "fmt"

// ProxyType is the protocol of a proxy server.
type ProxyType string

const (
	ProxyHTTP   ProxyType = "http"
	ProxyHTTPS  ProxyType = "https"
	ProxySOCKS4 ProxyType = "socks4"
	ProxySOCKS5 ProxyType = "socks5"
)

// Proxy is used by CapMonster workers to open the captcha page (not a local or transparent proxy).
// It is flattened into the task as proxyType, proxyAddress, proxyPort, proxyLogin, proxyPassword.
type Proxy struct {
	Type     ProxyType
	Address  string // IPv4/IPv6 address or hostname
	Port     int
	Login    string
	Password string
}

// String keeps the password out of logs (also used by %v and %+v).
func (p Proxy) String() string {
	pass := ""
	if p.Password != "" {
		pass = ":***"
	}
	login := ""
	if p.Login != "" {
		login = p.Login + pass + "@"
	}
	return fmt.Sprintf("%s://%s%s:%d", p.Type, login, p.Address, p.Port)
}

// GoString keeps the password out of %#v output.
func (p Proxy) GoString() string { return "capmonster.Proxy(" + p.String() + ")" }

func (p *Proxy) validate() error {
	switch p.Type {
	case ProxyHTTP, ProxyHTTPS, ProxySOCKS4, ProxySOCKS5:
	default:
		return &ValidationError{Field: "proxyType", Reason: fmt.Sprintf("must be http, https, socks4 or socks5, got %q", p.Type)}
	}
	if p.Address == "" {
		return &ValidationError{Field: "proxyAddress", Reason: "must not be empty"}
	}
	if p.Port < 1 || p.Port > 65535 {
		return &ValidationError{Field: "proxyPort", Reason: fmt.Sprintf("must be between 1 and 65535, got %d", p.Port)}
	}
	return nil
}

func (p *Proxy) fields() map[string]any {
	m := map[string]any{"proxyType": string(p.Type), "proxyAddress": p.Address, "proxyPort": p.Port}
	if p.Login != "" {
		m["proxyLogin"] = p.Login
	}
	if p.Password != "" {
		m["proxyPassword"] = p.Password
	}
	return m
}
