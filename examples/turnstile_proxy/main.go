// Usage: CAPMONSTER_API_KEY=... go run ./examples/turnstile_proxy
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	capmonster "github.com/CapMonsterCloud/capmonster-go-captcha-solver/v2"
)

func main() {
	client := capmonster.New(os.Getenv("CAPMONSTER_API_KEY"))

	// Your own deadline on top of the SDK's per-type timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Use the same proxy in your own scraping session and in the task.
	proxy := &capmonster.Proxy{Type: capmonster.ProxyHTTP, Address: "203.0.113.10", Port: 8080, Login: "login", Password: "password"}

	solution, err := client.Solve(ctx, capmonster.TurnstileRequest{
		WebsiteURL: "https://target-website.com/login",
		WebsiteKey: "0x4AAAAAAABnPIDnK2k_e-2",
		Proxy:      proxy,
	})
	var apiErr *capmonster.APIError
	switch {
	case errors.As(err, &apiErr):
		log.Fatalf("API error %s: %s", apiErr.Code, apiErr.Description)
	case errors.Is(err, capmonster.ErrTimeout):
		log.Fatal("not solved in time")
	case err != nil:
		log.Fatal(err)
	}
	// put the token into the cf-turnstile-response field and submit the form
	fmt.Println("Turnstile token:", solution.String("token"))
}
