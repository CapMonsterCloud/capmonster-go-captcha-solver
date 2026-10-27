// Package capmonster is the Go client for the CapMonster Cloud captcha-solving API.
//
//	client := capmonster.New("YOUR_API_KEY")
//	solution, err := client.Solve(ctx, &capmonster.RecaptchaV2Request{
//		WebsiteURL: "https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high",
//		WebsiteKey: "6Lcg7CMUAAAAANphynKgn9YAgA4tQ2KI_iqRyTwd",
//	})
//	if err != nil { ... }
//	fmt.Println(solution.String("gRecaptchaResponse"))
//
// Every request type (files *_request.go / *_requests.go) mirrors the API task: set the exported fields, empty
// optional fields are not sent. Required fields and value ranges are checked locally before anything is sent
// (*ValidationError). Errors returned by the API are *APIError with the API errorCode.
package capmonster
