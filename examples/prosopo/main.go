// Usage: CAPMONSTER_API_KEY=... go run ./examples/prosopo
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

	solution, err := client.Solve(context.Background(), capmonster.ProsopoTaskRequest{
		WebsiteURL: "https://example.com/login",
		WebsiteKey: "5EZU3LG31uzq1Mwi8inwqxmfvFDpj7VzwDnZwj4Q3syyxBwV",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
