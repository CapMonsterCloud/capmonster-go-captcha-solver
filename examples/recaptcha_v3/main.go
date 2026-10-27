// Usage: CAPMONSTER_API_KEY=... go run ./examples/recaptcha_v3
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

	solution, err := client.Solve(context.Background(), capmonster.RecaptchaV3ProxylessRequest{
		WebsiteURL: "https://lessons.zennolab.com/captchas/recaptcha/v3.php?level=beta",
		WebsiteKey: "6Le0xVgUAAAAAIt20XEB4rVhYOODgTl00d8juDob",
		MinScore:   0.9,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
