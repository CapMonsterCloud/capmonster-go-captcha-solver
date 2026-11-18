// Usage: CAPMONSTER_API_KEY=... go run ./examples/proton
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

	solution, err := client.Solve(context.Background(), capmonster.ProtonRequest{
		WebsiteURL: "https://example.com/",
		Metadata: capmonster.ProtonMetadata{
			HumanVerificationToken: "YOUR_HUMAN_VERIFICATION_TOKEN",
			PageAction:             capmonster.ProtonPageActionLogin,
		},
		Proxy: proxy,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
