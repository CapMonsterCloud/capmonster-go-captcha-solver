// Usage: CAPMONSTER_API_KEY=... go run ./examples/heytap
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

	solution, err := client.Solve(context.Background(), capmonster.HeyTapRequest{
		WebsiteURL: "https://example.com/",
		WebsiteKey: "YOUR_WEBSITE_KEY",
		Metadata:   &capmonster.HeyTapMetadata{APIServer: "api.example.com"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
