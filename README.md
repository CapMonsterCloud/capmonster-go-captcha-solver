# CapMonster Cloud Go SDK: Go Captcha Solver for reCAPTCHA, Cloudflare Turnstile & GeeTest

<p align="center">
  <a href="https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme">
    <img src="https://img.shields.io/badge/CapMonster%20Cloud-Go%20Captcha%20Solver-00B2FF?style=for-the-badge&logo=go&logoColor=white" alt="CapMonster Cloud Go SDK - captcha solver for Go (Golang)" height="40">
  </a>
</p>

Official Go (Golang) client library for the [CapMonster Cloud](https://capmonster.cloud/en/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) captcha recognition service. Use this Go captcha solver to solve (bypass) reCAPTCHA v2, reCAPTCHA v3, reCAPTCHA v2 Enterprise, Cloudflare Turnstile, GeeTest, reCAPTCHA image-click and text (image-to-text) captchas in web scraping, browser automation and backend services written in Go.

To get your API key, follow the [Getting Started Guide](https://docs.capmonster.cloud/docs/getting-start/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme): register on CapMonster Cloud, top up your balance and pass your personal API key to the client.

---

## Key Features

- **Simple client:** create a client with your API key and call a `Solve...` method for the captcha type you need.
- **Balance check:** `client.GetBalance()` returns your current CapMonster Cloud balance.
- **Typed tasks:** each captcha type has its own task constructor, for example `tasks.NewRecaptchaV2TaskProxyless` or `tasks.NewImageToTextTask`.
- **Image-to-text modules:** pick a recognition module with `WithCapMonsterModule`.
- **Backed by the official API:** request types follow the [CapMonster Cloud API documentation](https://docs.capmonster.cloud/docs/captchas/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme).

---

## Usage

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

## Supported captchas

Every request type below follows the official [CapMonster Cloud API documentation](https://docs.capmonster.cloud/docs/captchas/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme):

| Captcha type | Request types | Documentation |
| :--- | :--- | :--- |
| reCAPTCHA v2 | `RecaptchaV2ProxylessRequest`, `RecaptchaV2Request` | [RecaptchaV2Task](https://docs.capmonster.cloud/docs/captchas/no-captcha-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| reCAPTCHA v2 Enterprise | `RecaptchaV2EnterpriseProxylessRequest`, `RecaptchaV2EnterpriseRequest` | [RecaptchaV2EnterpriseTask](https://docs.capmonster.cloud/docs/captchas/recaptcha-v2-enterprise-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| reCAPTCHA v3 | `RecaptchaV3ProxylessRequest` | [RecaptchaV3TaskProxyless](https://docs.capmonster.cloud/docs/captchas/recaptcha-v3-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| reCAPTCHA image-click (ComplexImageTask) | `RecaptchaComplexImageTaskRequest` | [ComplexImageTask reCAPTCHA](https://docs.capmonster.cloud/docs/captchas/recaptcha-click/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| Cloudflare Turnstile | `TurnstileProxylessRequest`, `TurnstileRequest` | [TurnstileTask](https://docs.capmonster.cloud/docs/captchas/turnstile-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| GeeTest | `GeeTestProxylessRequest`, `GeeTestRequest` | [GeeTestTask](https://docs.capmonster.cloud/docs/captchas/geetest-task/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |
| Text captcha (image to text) | `ImageToTextRequest` | [ImageToTextTask](https://docs.capmonster.cloud/docs/captchas/image-to-text/?utm_source=github&utm_medium=referral&utm_campaign=go_repo_readme) |

---

## FAQ

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
