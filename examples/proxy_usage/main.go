// Usage: CAPMONSTER_API_KEY=... go run ./examples/proxy_usage
//
// The proxy is passed to CapMonster Cloud and used by the solver; use the same proxy in your own session.
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

	proxy := &capmonster.Proxy{Type: capmonster.ProxyHTTP, Address: "203.0.113.10", Port: 8080, Login: "login", Password: "password"} // replace with your proxy

	solution, err := client.Solve(context.Background(), capmonster.RecaptchaV2Request{
		WebsiteURL: "https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high",
		WebsiteKey: "6Lcg7CMUAAAAANphynKgn9YAgA4tQ2KI_iqRyTwd",
		Proxy:      proxy,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
