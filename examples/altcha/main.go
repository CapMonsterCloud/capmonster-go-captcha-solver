// Usage: CAPMONSTER_API_KEY=... go run ./examples/altcha
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

	solution, err := client.Solve(context.Background(), capmonster.AltchaCustomTaskRequest{
		WebsiteURL: "https://example.com/login",
		WebsiteKey: "altcha-public-key-123",
		Metadata: capmonster.AltchaMetadata{
			Challenge:  "challenge_string_here",
			Iterations: "1000",
			Salt:       "salt_string_here",
			Signature:  "signature_string_here",
		},
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
