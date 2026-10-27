// Usage: CAPMONSTER_API_KEY=... go run ./examples/tspd
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

	solution, err := client.Solve(context.Background(), capmonster.TspdCustomTaskRequest{
		WebsiteURL: "https://example.com/",
		Metadata: capmonster.TspdMetadata{
			TspdCookie:     "TS386a400d029=08267010245; TS386a400d078=08267dbb3b0c",
			HTMLPageBase64: "PCFET0NUWVBFIGh0bWw+PC9odG1sPg==", // the whole TSPD challenge page in base64
		},
		Proxy:     proxy,
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
