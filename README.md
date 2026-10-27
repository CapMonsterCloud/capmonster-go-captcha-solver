# CapMonster Cloud Go SDK: Golang Captcha Solver & Captcha Bypass for reCAPTCHA, Cloudflare Turnstile and GeeTest

<p align="center">
  <a href="https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme">
    <img src="https://img.shields.io/badge/CapMonster%20Cloud-Go%20Captcha%20Solver-00B2FF?style=for-the-badge&logo=go&logoColor=white" alt="CapMonster Cloud Go SDK - Golang captcha solver and captcha bypass API client" height="40">
  </a>
</p>

Official Go (Golang) client library for the [CapMonster Cloud](https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) AI-powered captcha solving service and captcha bypass API. Use this Go captcha solver to bypass reCAPTCHA v2, reCAPTCHA v3, reCAPTCHA v2 Enterprise, Cloudflare Turnstile and GeeTest, and to solve reCAPTCHA image-click and text (image-to-text) captchas, in Go web scraping, browser automation and backend services.

To get your API key, follow the [Getting Started Guide](https://docs.capmonster.cloud/docs/getting-start/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme): register on CapMonster Cloud, top up your balance and pass your personal API key to the client.

---

## Why Use the Go Captcha Solver SDK

- **Captcha bypass in a few lines of Go:** create a client with your API key and call the `Solve...` method for the captcha type you need.
- **reCAPTCHA bypass and Cloudflare Turnstile bypass:** request types for reCAPTCHA v2, v2 Enterprise, v3 and Cloudflare Turnstile, with and without a proxy where supported.
- **Captcha solving API for GeeTest and image captchas:** GeeTest, reCAPTCHA image-click and image-to-text recognition are covered by the same client.
- **Typed tasks:** each captcha type has its own task constructor, for example `tasks.NewRecaptchaV2TaskProxyless` or `tasks.NewImageToTextTask`.
- **Image-to-text modules:** pick a recognition module with `WithCapMonsterModule`.
- **Balance check:** `client.GetBalance()` returns your current CapMonster Cloud balance.
- **Backed by the official API:** request types follow the [CapMonster Cloud API documentation](https://docs.capmonster.cloud/docs/captchas/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

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

## Supported Captchas: Bypass reCAPTCHA, Cloudflare Turnstile, GeeTest and More

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

**Is there an official Go (Golang) captcha solver from CapMonster Cloud?**
Yes. This repository is the official Go client library for the CapMonster Cloud captcha solving service.

**How do I bypass reCAPTCHA v2 in Go?**
Create a client with your API key, build the task with `tasks.NewRecaptchaV2TaskProxyless`, call `client.SolveRecaptchaV2Proxyless` and use the returned token on the page. See the [reCAPTCHA v2 documentation](https://docs.capmonster.cloud/docs/captchas/no-captcha-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

**How do I bypass Cloudflare Turnstile in Go?**
Use the Turnstile request types from the table above. According to the [Turnstile documentation](https://docs.capmonster.cloud/docs/captchas/turnstile-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme), all Turnstile subtypes are supported automatically: manual, non-interactive and invisible.

**Which captcha types does the Go SDK solve?**
See the table above. You can also browse the [Supported Captchas Overview](https://docs.capmonster.cloud/docs/captchas/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) for all task types of the service.

**How do I get an API key?**
Register on CapMonster Cloud, top up your balance and copy your personal API key. The [Getting Started Guide](https://docs.capmonster.cloud/docs/getting-start/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) describes each step.

**Which Go version does the module target?**
Go 1.18 (see `go.mod`).

**How fast are captchas solved?**
On average between 1 and 30 seconds, depending on the captcha type. More details are in the [CapMonster Cloud FAQ](https://docs.capmonster.cloud/docs/faq/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

**How long is a reCAPTCHA v2 solution valid?**
According to the [reCAPTCHA v2 documentation](https://docs.capmonster.cloud/docs/captchas/no-captcha-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme), the returned token is valid for 60 seconds after the captcha is solved, so use it right away.

**Where is the full API reference?**
See the [API Methods Reference](https://docs.capmonster.cloud/docs/api/methods/create-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

---

## Responsible Use

CapMonster Cloud is intended for legitimate business purposes only, such as testing your own website or a website you are authorized to test, and recognizing images you provide. Read more about [ethical and transparent captcha automation](https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

---

## Documentation & Links

- [Getting Started Guide](https://docs.capmonster.cloud/docs/getting-start/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme)
- [Supported Captchas Overview](https://docs.capmonster.cloud/docs/captchas/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme)
- [API Methods Reference](https://docs.capmonster.cloud/docs/api/methods/create-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme)
- [CapMonster Cloud FAQ](https://docs.capmonster.cloud/docs/faq/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme)
- [CapMonster Cloud website](https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme)

---

## License

Released under the [MIT License](./LICENSE).
