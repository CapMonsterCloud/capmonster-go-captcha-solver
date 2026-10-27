// Usage: CAPMONSTER_API_KEY=... go run ./examples/tendi
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

	solution, err := client.Solve(context.Background(), capmonster.TenDiCustomTaskRequest{
		WebsiteURL: "https://example.com/",
		WebsiteKey: "189123456",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
