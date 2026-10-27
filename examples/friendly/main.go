// Usage: CAPMONSTER_API_KEY=... go run ./examples/friendly
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

	solution, err := client.Solve(context.Background(), capmonster.FriendlyCustomTaskRequest{
		WebsiteURL: "https://www.example.com",
		WebsiteKey: "FCMSITEKEY000000000000000000000000000000",
		Metadata:   capmonster.FriendlyMetadata{APIGetLib: "https://cdn.jsdelivr.net/npm/friendly-challenge@0.9.18/widget.module.min.js"},
		UserAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
