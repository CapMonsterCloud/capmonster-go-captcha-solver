// Usage: CAPMONSTER_API_KEY=... go run ./examples/recaptcha_image
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

	solution, err := client.Solve(context.Background(), capmonster.RecaptchaComplexImageTaskRequest{
		ImageURLs: []string{"https://i.postimg.cc/yYjg75Kv/payloadtraffic.jpg"},
		Metadata: capmonster.RecaptchaMetadata{
			Grid:           capmonster.Grid3x3,
			Task:           "Click on traffic lights",
			TaskDefinition: "/m/015qff",
		},
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
