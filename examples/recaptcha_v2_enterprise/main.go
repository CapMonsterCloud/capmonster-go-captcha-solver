// Usage: CAPMONSTER_API_KEY=... go run ./examples/recaptcha_v2_enterprise
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	capmonster "github.com/CapMonsterCloud/capmonster-go-captcha-solver/v2"
)

func main() {
	client := capmonster.New(os.Getenv("CAPMONSTER_API_KEY"))

	solution, err := client.Solve(context.Background(), capmonster.RecaptchaV2EnterpriseRequest{
		WebsiteURL:        "https://example.com/",
		WebsiteKey:        "6Lf26sUnAAAAAIKLuWNYgRsFUfmI-3Lex3xT5N-s",
		EnterprisePayload: map[string]string{"s": "SOME_ADDITIONAL_TOKEN"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
