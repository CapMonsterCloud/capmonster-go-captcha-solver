# CapMonster Cloud Go SDK: Golang Captcha Solver & Captcha Bypass for reCAPTCHA, Cloudflare Turnstile and GeeTest

> **🔄 Migrating from 2Captcha or Capsolver?** CapMonster Cloud is a fast, cost-effective **2captcha alternative** and **capsolver alternative** for Go: reCAPTCHA v3 from **$0.90 per 1,000 solves**, which is **up to 38% cheaper than 2Captcha** and **10% cheaper than Capsolver**, with an API response time of **≤ 2 s**. Already have software built for 2Captcha? CapMonster Cloud [supports the 2Captcha API format](https://docs.capmonster.cloud/docs/external-services/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme). Moving your Go code takes just a few lines: see the **Migration examples** below. Full comparison and disclaimer are in the comparison section below.

<p align="center">
  <a href="https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme">
    <img src="https://img.shields.io/badge/CapMonster%20Cloud-Go%20Captcha%20Solver-00B2FF?style=for-the-badge&logo=go&logoColor=white" alt="CapMonster Cloud Go SDK - Golang captcha solver and captcha bypass API client" height="40">
  </a>
</p>

Official Go (Golang) client library for the [CapMonster Cloud](https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) AI-powered captcha solving service and captcha bypass API. Use this Go captcha solver to bypass reCAPTCHA v2, reCAPTCHA v3, reCAPTCHA v2 Enterprise, Cloudflare Turnstile and GeeTest, and to solve reCAPTCHA image-click and text (image-to-text) captchas, in Go web scraping, browser automation and backend services. As the official CapMonster SDK for Go, it is a cost-effective 2captcha alternative and capsolver alternative.

To get your API key, follow the [Getting Started Guide](https://docs.capmonster.cloud/docs/getting-start/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme): register on CapMonster Cloud, top up your balance and pass your personal API key to the client.

---

## Why Use the Go Captcha Solver SDK

- **Captcha bypass in a few lines of Go:** create a client with your API key and call the `Solve...` method for the captcha type you need.
- **reCAPTCHA bypass and Cloudflare Turnstile bypass:** request types for reCAPTCHA v2, v2 Enterprise, v3 and Cloudflare Turnstile, with and without a proxy where supported.
- **Captcha solving API for GeeTest and image captchas:** GeeTest, reCAPTCHA image-click and image-to-text recognition are covered by the same client.
- **Typed tasks:** each captcha type has its own task constructor, for example `tasks.NewRecaptchaV2TaskProxyless` or `tasks.NewImageToTextTask`.
- **Image-to-text modules:** pick a recognition module with `WithCapMonsterModule`.
- **Balance check:** `client.GetBalance()` returns your current CapMonster Cloud balance.
- **Cheaper than alternatives:** reCAPTCHA v3 from $0.90 per 1,000 solves, versus $1.45 at 2Captcha and $1.00 at Capsolver.
- **Easy migration:** replace your 2Captcha or Capsolver solver call and keep the rest of your scraper or bot unchanged.
- **Backed by the official API:** request types follow the [CapMonster Cloud API documentation](https://docs.capmonster.cloud/docs/captchas/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

---

## 📊 CapMonster Cloud vs 2Captcha vs Capsolver

Looking for a 2captcha alternative or a capsolver alternative for Go web scraping and automation? Compare price, speed and success rate.

### reCAPTCHA v3

| Parameter | CapMonster Cloud | 2Captcha | CapSolver | DeathByCaptcha |
| :--- | :---: | :---: | :---: | :---: |
| Price per 1,000 solves (from) | **$0.90** | $1.45 | $1.00 | $2.89 |
| API response time | **≤ 2 s** | ≤ 2 s | < 3 s | ≤ 3 s |
| Solving success rate | **up to 99%** | ~77–92% | ~95–98% | ~64–90% |

*Disclaimer: 2Captcha, CapSolver, and DeathByCaptcha are trademarks of their respective owners. CapMonster Cloud is not affiliated with these entities. Performance metrics and pricing comparisons are based on publicly available data and internal benchmarking conducted on 27.09.2026.*

### Cloudflare Turnstile

| Parameter | CapMonster Cloud | 2Captcha | Anti-Captcha | DeathByCaptcha |
| :--- | :---: | :---: | :---: | :---: |
| Price per 1,000 solves (from) | **$1.30** | $1.45 | $2.00 | $2.89 |
| Time to solve | **6.24 s\*** | 16.9 s | 20.33 s | ~13 s |
| Solving success rate | **99%\*** | 96% | 96% | 96% |

*Disclaimer: 2Captcha, Anti-Captcha, and DeathByCaptcha are trademarks of their respective owners. CapMonster Cloud is not affiliated with these entities. Performance metrics and pricing comparisons are based on publicly available data and internal benchmarking as of 12.10.2026.*

**What this means in practice:**

- **reCAPTCHA v3:** $0.90 versus $1.45 at 2Captcha is up to 38% cheaper, and $0.90 versus $1.00 at Capsolver is 10% cheaper.
- **Cloudflare Turnstile:** $1.30 versus $1.45 at 2Captcha is 10% cheaper, and a solving time of 6.24 s versus 16.9 s is more than 60% faster.

---

## Quick Start: Solve and Bypass Captchas in Go

The example below checks your balance, bypasses a reCAPTCHA v2 (without proxy) and recognizes an image-to-text captcha:

```go
package main

import (
"fmt"

"github.com/ZennoLab/capmonstercloud-client-go/pkg/client"
"github.com/ZennoLab/capmonstercloud-client-go/pkg/tasks"
)

func main() {
client := client.New("YourClientKey")

//get balance
if balance, err := client.GetBalance(); err != nil {
fmt.Println(err)
} else {
fmt.Println(balance)
}

// solve RecaptchaV2 (without proxy)
recaptchaV2Task := tasks.NewRecaptchaV2TaskProxyless(
"[https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high](https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high)",
"6Lcg7CMUAAAAANphynKgn9YAgA4tQ2KI_iqRyTwd",
)
noCache := false
if result, err := client.SolveRecaptchaV2Proxyless(recaptchaV2Task, noCache, nil); err != nil {
fmt.Println(err)
} else {
fmt.Println(result)
}

// solve ImageToTextTask with module
imgToTextTask := tasks.NewImageToTextTask("BODY").WithCapMonsterModule(tasks.CapMonsterModuleGoogle)
if result, err := client.SolveImageToText(imgToTextTask, nil); err != nil {
fmt.Println(err)
} else {
fmt.Println(result)
}
}
```

---

## 🔄 Migrating from 2Captcha or Capsolver

Already use 2Captcha or Capsolver in your Go scraper or bot? Replace only the solver call and keep the rest of your code unchanged. CapMonster Cloud follows the same flow: send the website URL and sitekey, receive the token.

### Migration from 2Captcha (2captcha-go)

```go
// Before: 2Captcha
client := api2captcha.NewClient("YOUR_2CAPTCHA_API_KEY")
captcha := api2captcha.ReCaptcha{
	SiteKey: "YOUR_SITE_KEY",
	Url:     "[https://target-website.com/page-with-recaptcha](https://target-website.com/page-with-recaptcha)",
}
code, err := client.Solve(captcha.ToRequest())
```

### Migration from Capsolver (capsolver-go)

```go
// Before: Capsolver (API key is read from the CAPSOLVER_API_KEY environment variable)
capSolver := capsolver.CapSolver{}
s, err := capSolver.Solve(map[string]any{
	"type":       "ReCaptchaV2taskProxyLess",
	"websiteURL": "[https://target-website.com/page-with-recaptcha](https://target-website.com/page-with-recaptcha)",
	"websiteKey": "YOUR_SITE_KEY",
})
code := s.Solution.GRecaptchaResponse
```

### After: CapMonster Cloud

```go
package main

import (
	"fmt"

	"github.com/ZennoLab/capmonstercloud-client-go/pkg/client"
	"github.com/ZennoLab/capmonstercloud-client-go/pkg/tasks"
)

func main() {
	cm := client.New("YOUR_CAPMONSTER_API_KEY")

	task := tasks.NewRecaptchaV2TaskProxyless(
		"[https://target-website.com/page-with-recaptcha](https://target-website.com/page-with-recaptcha)",
		"YOUR_SITE_KEY",
	)

	solution, err := cm.SolveRecaptchaV2Proxyless(task, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(solution.GRecaptchaResponse)
}
```

Only the client, the API key and the task constructor change: the token (`GRecaptchaResponse`) is used exactly as before, so the rest of your scraper stays the same.

### Keep existing 2Captcha-based software

Software that already speaks the 2Captcha API can work with CapMonster Cloud without code changes. See the [guide for connecting other services' APIs](https://docs.capmonster.cloud/docs/external-services/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

---

## Supported Captchas: Bypass reCAPTCHA, Cloudflare Turnstile, GeeTest and More

Use this Go captcha solver for captcha bypass in web scraping and automation: reCAPTCHA bypass (v2, v2 Enterprise, v3), Cloudflare Turnstile solver, GeeTest solver, FunCaptcha solving and image-to-text recognition.

Every request type below follows the official [CapMonster Cloud API documentation](https://docs.capmonster.cloud/docs/captchas/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme):

| Captcha bypass / solving | Request types | Documentation |
| :--- | :--- | :--- |
| reCAPTCHA v2 bypass | `RecaptchaV2ProxylessRequest`, `RecaptchaV2Request` | [RecaptchaV2Task](https://docs.capmonster.cloud/docs/captchas/no-captcha-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| reCAPTCHA v2 Enterprise bypass | `RecaptchaV2EnterpriseProxylessRequest`, `RecaptchaV2EnterpriseRequest` | [RecaptchaV2EnterpriseTask](https://docs.capmonster.cloud/docs/captchas/recaptcha-v2-enterprise-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| reCAPTCHA v3 bypass | `RecaptchaV3ProxylessRequest` | [RecaptchaV3TaskProxyless](https://docs.capmonster.cloud/docs/captchas/recaptcha-v3-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| reCAPTCHA image-click solver (ComplexImageTask) | `RecaptchaComplexImageTaskRequest` | [ComplexImageTask reCAPTCHA](https://docs.capmonster.cloud/docs/captchas/recaptcha-click/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| Cloudflare Turnstile bypass | `TurnstileProxylessRequest`, `TurnstileRequest` | [TurnstileTask](https://docs.capmonster.cloud/docs/captchas/turnstile-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| GeeTest solver | `GeeTestProxylessRequest`, `GeeTestRequest` | [GeeTestTask](https://docs.capmonster.cloud/docs/captchas/geetest-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| Text captcha solver (image to text) | `ImageToTextRequest` | [ImageToTextTask](https://docs.capmonster.cloud/docs/captchas/image-to-text/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |

---

## Use Cases

- **Web scraping and data collection** from websites you are authorized to access, without manual captcha solving.
- **Browser automation:** get the solution token in Go and inject it into the page with your automation tool.
- **Automated testing (QA)** of your own website and its captcha-protected forms.
- **Backend services** that need a captcha solving API inside a Go application.

---

## FAQ

**Is CapMonster Cloud a 2Captcha alternative for Go?**
Yes. It is a captcha solving service with an official Go SDK. For reCAPTCHA v3 it starts at $0.90 per 1,000 solves versus $1.45 at 2Captcha, which is up to 38% cheaper. For Cloudflare Turnstile it starts at $1.30 versus $1.45,
