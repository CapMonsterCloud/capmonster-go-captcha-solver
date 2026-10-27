// Usage: CAPMONSTER_API_KEY=... go run ./examples/binance
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

	solution, err := client.Solve(context.Background(), capmonster.BinanceTaskRequest{
		WebsiteURL: "https://accounts.binance.com/ru/login?loginChannel=&return_to=",
		WebsiteKey: "login",
		ValidateID: "2b8137c0b9b44189800368819354e114",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
