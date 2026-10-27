// Usage: CAPMONSTER_API_KEY=... go run ./examples/alibaba
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

	solution, err := client.Solve(context.Background(), capmonster.AlibabaCustomTaskRequest{
		WebsiteURL: "https://www.example.com",
		Metadata: capmonster.AlibabaMetadata{
			SceneID:   "1ww7426c",
			Prefix:    "dlw3kug",
			APIGetLib: "https://o.example.com/captcha-frontend/aliyunCaptcha/AliyunCaptcha.js?t=2041",
		},
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
