# Zennolab.CapMonsterCloud.Client

Official Go client library for [capmonster.cloud](https://capmonster.cloud/) captcha recognition service

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
		"https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high",
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
Supported captcha recognition requests:

- [GeeTestProxylessRequest](https://docs.capmonster.cloud/docs/captchas/geetest-task)
- [GeeTestRequest](https://docs.capmonster.cloud/docs/captchas/geetest-task)
- [ImageToTextRequest](https://docs.capmonster.cloud/docs/captchas/image-to-text)
- [RecaptchaV2ProxylessRequest](https://docs.capmonster.cloud/docs/captchas/no-captcha-task)
- [RecaptchaV2Request](https://docs.capmonster.cloud/docs/captchas/no-captcha-task)
- [RecaptchaV3ProxylessRequest](https://docs.capmonster.cloud/docs/captchas/recaptcha-v3-task)
- [RecaptchaV2EnterpriseProxylessRequest](https://docs.capmonster.cloud/docs/captchas/recaptcha-v2-enterprise-task)
- [RecaptchaV2EnterpriseRequest](https://docs.capmonster.cloud/docs/captchas/recaptcha-v2-enterprise-task)
- [TurnstileProxylessRequest](https://docs.capmonster.cloud/docs/captchas/tunrstile-task)
- [TurnstileRequest](https://docs.capmonster.cloud/docs/captchas/tunrstile-task)
- [RecaptchaComplexImageTaskRequest](https://docs.capmonster.cloud/docs/captchas/recaptcha-click)
