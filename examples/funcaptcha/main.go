// Usage: CAPMONSTER_API_KEY=... go run ./examples/funcaptcha
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

	solution, err := client.Solve(context.Background(), capmonster.FuncaptchaRequest{
		WebsiteURL:               "https://example.com/",
		WebsitePublicKey:         "69A21A01-CC7B-B9C6-0F9A-E7FA06677FFC",
		FuncaptchaAPIJSSubdomain: "example-api.arkoselabs.com",
		Proxy:                    proxy,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
