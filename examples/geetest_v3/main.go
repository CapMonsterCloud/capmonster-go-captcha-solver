// Usage: CAPMONSTER_API_KEY=... go run ./examples/geetest_v3
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

	solution, err := client.Solve(context.Background(), capmonster.GeetestV3Request{
		WebsiteURL: "https://example.com/",
		GT:         "81388ea1fc187e0c335c0a8907ff2625",
		Challenge:  "2e2f0f65240058b683cb6ea21c303eea6n", // fresh challenge, valid for one solve only
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
