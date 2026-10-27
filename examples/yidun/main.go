// Usage: CAPMONSTER_API_KEY=... go run ./examples/yidun
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

	solution, err := client.Solve(context.Background(), capmonster.YidunRequest{
		WebsiteURL: "https://example.com/",
		WebsiteKey: "6cw0f0485d5d46auacf9b735d20218a5",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
