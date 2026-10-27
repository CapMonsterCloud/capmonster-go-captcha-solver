// Usage: CAPMONSTER_API_KEY=... go run ./examples/basilisk
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

	solution, err := client.Solve(context.Background(), capmonster.BasiliskCustomTaskRequest{
		WebsiteURL: "https://example.com/account/register",
		WebsiteKey: "b7890hre5cf2544b2759c19fb2600897",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
