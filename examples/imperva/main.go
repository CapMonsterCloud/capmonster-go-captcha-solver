// Usage: CAPMONSTER_API_KEY=... go run ./examples/imperva
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

	solution, err := client.Solve(context.Background(), capmonster.ImpervaCustomTaskRequest{
		WebsiteURL: "https://example.com/login",
		Metadata: capmonster.ImpervaMetadata{
			IncapsulaScriptURL: "_Incapsula_Resource?SWJIYLWA=719d34d31c8e3a6e6fffd425f7e032f3",
			IncapsulaCookies:   "incap_ses_1166_2930313=br7iX33ZNCtf3HlpEXcuEDzz72cAAAAA0suDnBGrq/iA0J4oERYzjQ==; visid_incap_2930313=P3hgPVm9S8Oond1L0sXhZqfK72cAAAAAQUIPAAAAAABoMSY9xZ34RvRseJRiY6s+;",
			Reese84URLEndpoint: "Built-with-the-For-hopence-Hurleysurfecting-the-",
		},
		Proxy:     proxy,
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
