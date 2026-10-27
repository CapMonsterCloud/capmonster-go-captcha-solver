// Usage: CAPMONSTER_API_KEY=... go run ./examples/client_proxy
//
// WithClientProxy routes the SDK's own HTTP requests to api.capmonster.cloud through a proxy.
// It is not sent to the solver: to solve through a proxy, see examples/proxy_usage.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	capmonster "github.com/CapMonsterCloud/capmonster-go-captcha-solver/v2"
)

func main() {
	client, err := capmonster.NewWithOptions(os.Getenv("CAPMONSTER_API_KEY"),
		capmonster.WithClientProxy("http://login:password@203.0.113.10:8080"), // replace with your proxy
	)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := client.GetBalance(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance: %.4f USD\n", balance)
}
