// Usage: CAPMONSTER_API_KEY=... go run ./examples/geetest_v4
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

	solution, err := client.Solve(context.Background(), capmonster.GeetestV4Request{
		WebsiteURL:     "https://faucetpay.io/account/login",
		GT:             "4eb8b0c2b27f3365b9244d9da81638c6",
		InitParameters: map[string]any{"riskType": "slide"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
