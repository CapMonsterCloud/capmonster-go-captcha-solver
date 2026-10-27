// Usage: CAPMONSTER_API_KEY=... go run ./examples/image_to_text path/to/captcha.png
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	capmonster "github.com/CapMonsterCloud/capmonster-go-captcha-solver/v2"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: image_to_text <image file>")
	}
	req, err := capmonster.ImageToTextFromFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	req.CapMonsterModule = capmonster.ModuleUniversal

	client := capmonster.New(os.Getenv("CAPMONSTER_API_KEY"))
	solution, err := client.Solve(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Text:", solution.String("text"))
}
