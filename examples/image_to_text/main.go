// Usage: CAPMONSTER_API_KEY=... go run ./examples/image_to_text [image file, default examples/images/sample.jpg]
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	capmonster "github.com/CapMonsterCloud/capmonster-go-captcha-solver/v2"
)

func main() {
	path := "examples/images/sample.jpg"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	req, err := capmonster.ImageToTextFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	req.CapMonsterModule = capmonster.ModuleUniversal
	req.RecognizingThreshold = capmonster.Int(50)
	req.Case = capmonster.Bool(true)

	client := capmonster.New(os.Getenv("CAPMONSTER_API_KEY"))

	solution, err := client.Solve(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
