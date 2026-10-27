// Usage: CAPMONSTER_API_KEY=... go run ./examples/balance
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

	balance, err := client.GetBalance(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Balance: %.4f USD\n", balance)
}
