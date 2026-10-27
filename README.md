# CapMonster Cloud Go SDK: Golang Captcha Solver & Captcha Bypass for reCAPTCHA, Cloudflare Turnstile and GeeTest

> **🔄 Migrating from 2Captcha or Capsolver?** CapMonster Cloud is a fast, cost-effective **2captcha alternative** and **capsolver alternative** for Go: reCAPTCHA v3 from **$0.90 per 1,000 solves**, which is **up to 38% cheaper than 2Captcha** and **10% cheaper than Capsolver**, with an API response time of **≤ 2 s**. Already have software built for 2Captcha? CapMonster Cloud [supports the 2Captcha API format](https://docs.capmonster.cloud/docs/external-services/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme). Moving your Go code takes just a few lines: see the **Migration examples** below. Full comparison and disclaimer are in the comparison section below.

<p align="center">
  <a href="https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme">
    <img src="https://img.shields.io/badge/CapMonster%20Cloud-Go%20Captcha%20Solver-00B2FF?style=for-the-badge&logo=go&logoColor=white" alt="CapMonster Cloud Go SDK - captcha solver and captcha bypass API client" height="40">
  </a>
</p>

<p align="center">
  <strong>Fast, AI-driven CAPTCHA solving client for Go. Seamlessly bypass Cloudflare Turnstile, reCAPTCHA v2/v3/Enterprise and GeeTest in web scraping and browser automation.</strong>
</p>

<p align="center">
  <a href="https://github.com/CapMonsterCloud/capmonster-go-captcha-solver/stargazers"><img src="https://img.shields.io/github/stars/CapMonsterCloud/capmonster-go-captcha-solver?style=flat-square&color=yellow" alt="GitHub Stars"></a>
  <a href="https://github.com/CapMonsterCloud/capmonster-go-captcha-solver/network/members"><img src="https://img.shields.io/github/forks/CapMonsterCloud/capmonster-go-captcha-solver?style=flat-square" alt="GitHub Forks"></a>
  <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square" alt="License: MIT"></a>
</p>

---

Official Go (Golang) client library for the [CapMonster Cloud](https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) AI-powered captcha solving service and captcha bypass API. Use this Go captcha solver to bypass reCAPTCHA v2, reCAPTCHA v3, reCAPTCHA v2 Enterprise, Cloudflare Turnstile and GeeTest, and to solve reCAPTCHA image-click and text (image-to-text) captchas, in Go web scraping, browser automation and backend services. As the official CapMonster SDK for Go, it is a cost-effective 2captcha alternative and capsolver alternative.

**[👉 Get your Free API Key & Free Trial Balance on CapMonster Cloud](https://dash.capmonster.cloud/Account/SignUp?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme)**

---

## ⚡ Highlights

- ⚡ **Fast solving:** AI-based recognition, on average between 1 and 30 seconds depending on the captcha type.
- 🤖 **Zero human workers:** 100% automated machine learning pipeline with no manual latencies.
- 🛡️ **Modern anti-bot bypass:** Support for Cloudflare Turnstile, reCAPTCHA Enterprise, GeeTest and FunCaptcha.
- 🧩 **Typed requests:** each captcha type is a plain struct, for example `capmonster.RecaptchaV2Request` or `capmonster.ImageToTextRequest`, passed to one `client.Solve` method.
- 🖼️ **Image-to-text modules:** pick a recognition module with the `CapMonsterModule` field, for example `capmonster.ModuleGoogle`.
- 💳 **Balance check:** `client.GetBalance(ctx)` returns your current CapMonster Cloud balance.
- 💰 **Cost-effective:** Low-cost pricing structure based on pay-as-you-go balance.
- 📉 **Cheaper than alternatives:** reCAPTCHA v3 from $0.90 per 1,000 solves, versus $1.45 at 2Captcha and $1.00 at Capsolver.
- 🔄 **Easy migration:** replace your 2Captcha or Capsolver solver call and keep the rest of your scraper or bot unchanged.
- 📖 **Official docs:** request types follow the [CapMonster Cloud API documentation](https://docs.capmonster.cloud/docs/captchas/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

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

## 📦 Installation

Requires Go 1.21 or newer.

```bash
go get github.com/CapMonsterCloud/capmonster-go-captcha-solver/v2
```

Import the module as `capmonster "github.com/CapMonsterCloud/capmonster-go-captcha-solver/v2"`, as shown in the Quick Start below.

---

## 🚀 Quick Start: Solve and Bypass Captchas in Go

The example below checks your balance, bypasses a reCAPTCHA v2 (without proxy) and recognizes an image-to-text captcha:

```go
package main

import (
	"context"
	"fmt"
	"log"

	capmonster "github.com/CapMonsterCloud/capmonster-go-captcha-solver/v2"
)

func main() {
	ctx := context.Background()
	client := capmonster.New("YourClientKey")

	// get balance
	if balance, err := client.GetBalance(ctx); err != nil {
		log.Println(err)
	} else {
		fmt.Println(balance)
	}

	// solve reCAPTCHA v2 (without proxy)
	solution, err := client.Solve(ctx, capmonster.RecaptchaV2Request{
		WebsiteURL: "https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high",
		WebsiteKey: "6Lcg7CMUAAAAANphynKgn9YAgA4tQ2KI_iqRyTwd",
	})
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(solution.String("gRecaptchaResponse"))
	}

	// solve ImageToText with a recognition module
	imgReq, err := capmonster.ImageToTextFromFile("captcha.png")
	if err != nil {
		log.Fatal(err)
	}
	imgReq.CapMonsterModule = capmonster.ModuleGoogle
	if solution, err := client.Solve(ctx, imgReq); err != nil {
		log.Println(err)
	} else {
		fmt.Println(solution.String("text"))
	}
}
```

---

## 🔄 Migrating from 2Captcha or Capsolver

Already use 2Captcha or Capsolver in your Go scraper or bot? Replace only the solver call and keep the rest of your code unchanged. CapMonster Cloud follows the same flow: send the website URL and sitekey, receive the token.

### Before: 2Captcha (2captcha-go)

```go
// Before: 2Captcha
client := api2captcha.NewClient("YOUR_2CAPTCHA_API_KEY")
captcha := api2captcha.ReCaptcha{
	SiteKey: "YOUR_SITE_KEY",
	Url:     "https://target-website.com/page-with-recaptcha",
}
code, err := client.Solve(captcha.ToRequest())
```

### Before: Capsolver (capsolver-go)

```go
// Before: Capsolver (API key is read from the CAPSOLVER_API_KEY environment variable)
capSolver := capsolver.CapSolver{}
s, err := capSolver.Solve(map[string]any{
	"type":       "ReCaptchaV2taskProxyLess",
	"websiteURL": "https://target-website.com/page-with-recaptcha",
	"websiteKey": "YOUR_SITE_KEY",
})
code := s.Solution.GRecaptchaResponse
```

### After: CapMonster Cloud

```go
package main

import (
	"context"
	"fmt"

	capmonster "github.com/CapMonsterCloud/capmonster-go-captcha-solver/v2"
)

func main() {
	cm := capmonster.New("YOUR_CAPMONSTER_API_KEY")

	solution, err := cm.Solve(context.Background(), capmonster.RecaptchaV2Request{
		WebsiteURL: "https://target-website.com/page-with-recaptcha",
		WebsiteKey: "YOUR_SITE_KEY",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(solution.String("gRecaptchaResponse"))
}
```

Only the client, the API key and the request struct change: the token (`gRecaptchaResponse`) is used exactly as before, so the rest of your scraper stays the same.

### Keep existing 2Captcha-based software

Software that already speaks the 2Captcha API can work with CapMonster Cloud without code changes. See the [guide for connecting other services' APIs](https://docs.capmonster.cloud/docs/external-services/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

---

## 🛡️ Supported CAPTCHA Types: Bypass reCAPTCHA, Cloudflare Turnstile, GeeTest and More

Use this Go captcha solver for captcha bypass in web scraping and automation: reCAPTCHA bypass (v2, v2 Enterprise, v3), Cloudflare Turnstile solver, GeeTest solver, FunCaptcha solving and image-to-text recognition.

All task types follow the official [CapMonster Cloud API documentation](https://docs.capmonster.cloud/docs/captchas/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme):

| Captcha bypass / solving | Request types | Documentation |
| :--- | :--- | :--- |
| reCAPTCHA v2 bypass | `RecaptchaV2Request` (optional `Proxy`) | [RecaptchaV2Task](https://docs.capmonster.cloud/docs/captchas/no-captcha-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| reCAPTCHA v2 Enterprise bypass | `RecaptchaV2EnterpriseRequest` (optional `Proxy`) | [RecaptchaV2EnterpriseTask](https://docs.capmonster.cloud/docs/captchas/recaptcha-v2-enterprise-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| reCAPTCHA v3 bypass | `RecaptchaV3ProxylessRequest` | [RecaptchaV3TaskProxyless](https://docs.capmonster.cloud/docs/captchas/recaptcha-v3-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| reCAPTCHA image-click solver (ComplexImageTask) | `RecaptchaComplexImageTaskRequest` | [ComplexImageTask reCAPTCHA](https://docs.capmonster.cloud/docs/captchas/recaptcha-click/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| Cloudflare Turnstile bypass | `TurnstileRequest` (optional `Proxy`), `TurnstileChallengeTokenRequest`, `TurnstileChallengeCookieRequest`, `TurnstileWaitRoomRequest` | [TurnstileTask](https://docs.capmonster.cloud/docs/captchas/turnstile-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| GeeTest solver | `GeetestV3Request`, `GeetestV4Request` | [GeeTestTask](https://docs.capmonster.cloud/docs/captchas/geetest-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| Text captcha solver (image to text) | `ImageToTextRequest` | [ImageToTextTask](https://docs.capmonster.cloud/docs/captchas/image-to-text/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |

---

## 🛠️ How It Works

```text
[ Your Go service ]
         │
         ▼
[ Extract sitekey & website URL ]
         │
         ▼
[ CapMonster Cloud Go SDK ] ──► createTask API request
         │
         ▼
[ Neural Network Processing ] ──► Poll getTaskResult until the task is ready
         │
         ▼
[ Receive Token & Autosubmit ] ──► Complete bypass
```

---

## ⚙️ Best Practices

- **Use tokens promptly:** CAPTCHA tokens expire quickly (a reCAPTCHA v2 token is valid for 60 seconds after solving), so submit or inject the solution immediately after receiving the response.
- **Keep session context consistent:** When a task uses a proxy, use the same proxy in your scraping or browser session and in the CapMonster task.
- **Use the right task type:** Choose the request type that matches the target protection and follow the required parameters in the [documentation](https://docs.capmonster.cloud/docs/captchas/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).
- **Monitor your balance:** Keep your account balance positive in the [Dashboard](https://dash.capmonster.cloud/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) and handle API errors in your code.

---

## 💡 Use Cases

- **Web scraping and data collection** from websites you are authorized to access, without manual captcha solving.
- **Browser automation:** get the solution token in Go and inject it into the page with your automation tool.
- **Automated testing (QA)** of your own website and its captcha-protected forms.
- **Backend services** that need a captcha solving API inside a Go application.

---

## ❓ FAQ

**Is CapMonster Cloud a 2Captcha alternative for Go?**
Yes. It is a captcha solving service with an official Go SDK. For reCAPTCHA v3 it starts at $0.90 per 1,000 solves versus $1.45 at 2Captcha, which is up to 38% cheaper. For Cloudflare Turnstile it starts at $1.30 versus $1.45, which is 10% cheaper.

**Is CapMonster Cloud a Capsolver alternative for Go?**
Yes. For reCAPTCHA v3 it is 10% cheaper than Capsolver ($0.90 versus $1.00 per 1,000 solves) and shows an API response time of ≤ 2 s versus < 3 s. See the comparison table above and its disclaimer.

**How do I migrate from 2Captcha or Capsolver to CapMonster Cloud in Go?**
Create a client with `capmonster.New("YOUR_CAPMONSTER_API_KEY")`, fill the matching request struct and call `client.Solve(ctx, req)`. See the migration examples above. Software that already supports the 2Captcha API can also be connected through the [CapMonster Cloud guide](https://docs.capmonster.cloud/docs/external-services/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

**Is there an official Go captcha solver from CapMonster Cloud?**
Yes. This repository is the official Go client library for the CapMonster Cloud captcha solving service.

**How do I bypass reCAPTCHA v2 in Go?**
Create a client with your API key, pass a `capmonster.RecaptchaV2Request` to `client.Solve` and use `solution.String("gRecaptchaResponse")` on the page. See the [reCAPTCHA v2 documentation](https://docs.capmonster.cloud/docs/captchas/no-captcha-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

**How do I bypass Cloudflare Turnstile in Go?**
Use the Turnstile request types from the supported types table above. According to the [Turnstile documentation](https://docs.capmonster.cloud/docs/captchas/turnstile-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme), all Turnstile subtypes are supported automatically: manual, non-interactive and invisible.

**Which captcha types does the Go SDK solve?**
reCAPTCHA v2, v2 Enterprise and v3, Cloudflare Turnstile, GeeTest, FunCaptcha, image-to-text and reCAPTCHA image-click tasks. See the supported types table above and the [Supported Captchas Overview](https://docs.capmonster.cloud/docs/captchas/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

**How do I get an API key?**
Register on CapMonster Cloud, top up your balance and copy your API key. The [Getting Started Guide](https://docs.capmonster.cloud/docs/getting-start/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) describes each step.

**Which Go version does the module target?**
Go 1.21 or newer (see `go.mod`).

**How fast are captchas solved?**
On average between 1 and 30 seconds, depending on the captcha type. More details are in the [CapMonster Cloud FAQ](https://docs.capmonster.cloud/docs/faq/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

**How long is a reCAPTCHA v2 solution valid?**
According to the [reCAPTCHA v2 documentation](https://docs.capmonster.cloud/docs/captchas/no-captcha-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme), the returned token is valid for 60 seconds after the captcha is solved, so use it right away.

**Where is the full API reference?**
See the [API Methods Reference](https://docs.capmonster.cloud/docs/methods/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

---

## ⚖️ Responsible Use

CapMonster Cloud is intended for legitimate business purposes only, such as testing your own website or a website you are authorized to test, and recognizing images you provide. Read more about [ethical and transparent captcha automation](https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

---

## 📚 Documentation & Support

- 📖 [Getting Started Guide](https://docs.capmonster.cloud/docs/getting-start/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme)
- 🎯 [Supported Captchas Overview](https://docs.capmonster.cloud/docs/captchas/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme)
- ⚙️ [API Methods Reference](https://docs.capmonster.cloud/docs/methods/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme)
- ❓ [CapMonster Cloud FAQ](https://docs.capmonster.cloud/docs/faq/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme)
- 🌐 [Browser Extension Guides](https://docs.capmonster.cloud/docs/extension/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme)
- 💬 [Support & Community](https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme#support)

## Other official SDKs & integrations

- [Python](https://github.com/CapMonsterCloud/capmonster-python-captcha-solver)
- [Node.js](https://github.com/CapMonsterCloud/capmonster-nodejs-captcha-solver)
- [.NET](https://github.com/CapMonsterCloud/capmonster-dotnet-captcha-solver)
- [PHP](https://github.com/CapMonsterCloud/capmonster-php-captcha-solver)
- [Rust](https://github.com/CapMonsterCloud/capmonster-rust-captcha-solver)
- [n8n](https://github.com/CapMonsterCloud/capmonster-n8n-captcha-solver)
- [MCP server](https://github.com/CapMonsterCloud/capmonster-mcp-captcha-solver)
- [MCP server for Patchright](https://github.com/CapMonsterCloud/capmonster-mcp-patchright-captcha-solver)
- [API Docs](https://github.com/CapMonsterCloud/capmonster-captcha-solver-docs)

---

## 📄 License

[MIT](./LICENSE) © [CapMonster Cloud](https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme)
