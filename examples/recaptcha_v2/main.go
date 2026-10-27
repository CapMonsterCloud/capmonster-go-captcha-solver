// Usage: CAPMONSTER_API_KEY=... go run ./examples/recaptcha_v2
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
	ctx := context.Background()

	balance, err := client.GetBalance(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance: %.4f USD\n", balance)

	solution, err := client.Solve(ctx, capmonster.RecaptchaV2Request{
		WebsiteURL: "https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high",
		WebsiteKey: "6Lcg7CMUAAAAANphynKgn9YAgA4tQ2KI_iqRyTwd",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gRecaptchaResponse:", solution.String("gRecaptchaResponse"))
}
