# Changelog

All notable changes to this module are documented here. The format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/), versions follow [SemVer](https://semver.org/)
and Go module rules (a new major version gets a new `/vN` import path).

## [2.0.0]

Full rewrite under a new import path, `github.com/CapMonsterCloud/capmonster-go-captcha-solver/v2`
(1.x was `github.com/ZennoLab/capmonstercloud-client-go`, which did not match the repository and could not be
fetched by the repository's own address). 1.x users are not affected until they switch imports; see
"Migrating from 1.x" in README.md.

### Added
- One package, `capmonster`: `Client.Solve` (createTask + polling with a per-type schedule), low-level
  `CreateTask` / `GetTaskResult`, `GetBalance`. Every call takes a `context.Context`.
- Typed errors for `errors.As`: `*ValidationError`, `*APIError` (with the API `Code`), `*TimeoutError`
  (`errors.Is(err, ErrTimeout)`), `*StatusError`, `*MalformedResponseError`.
- Options: `WithServiceURL`, `WithSoftID`, `WithHTTPClient`, `WithClientTimeout`, `WithClientProxy`
  (http/https/socks5 proxy for the SDK's own requests), `WithTimeoutsOverride`.
- Request types matching the current API documentation: reCAPTCHA v2 / v2 Enterprise / v3 / v3 Enterprise,
  reCAPTCHA click and Recognition ComplexImageTask, ImageToText, FunCaptcha, GeeTest v3/v4,
  Turnstile (standard, Cloudflare Challenge token / cf_clearance, Waiting Room), Amazon WAF (3 variants),
  Binance, MTCaptcha, Yidun, Prosopo, and CustomTask: DataDome, Imperva, TSPD, HUNT, Basilisk, TenDI, Alibaba,
  Friendly Captcha, Altcha.
- Local validation before anything is sent; proxy passwords are masked in `fmt` output.

### Removed
- One `Solve<Type>` method per captcha type and per-type solution structs: `Solve` returns a `Solution`
  (`String(key)`, `Decode(&v)`).
- hCaptcha and hCaptcha / FunCaptcha ComplexImageTask: no longer in the API documentation.
- `*Proxyless` task types: the API now uses one task type per captcha with an optional proxy.
- `noCache` and `callbackUrl` arguments.

### Fixed
- The `User-Agent` header was sent as `UserAgent`.
- The `getTaskResult` URL had a trailing slash.
- API error codes are no longer replaced with fixed English strings.
