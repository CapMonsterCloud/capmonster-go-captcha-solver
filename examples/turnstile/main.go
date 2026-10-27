// Usage: CAPMONSTER_API_KEY=... go run ./examples/turnstile
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

	solution, err := client.Solve(context.Background(), capmonster.TurnstileRequest{
		WebsiteURL: "http://tsmanaged.zlsupport.com/",
		WebsiteKey: "0x4AAAAAAABUYP0XeMJF0xoy",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
