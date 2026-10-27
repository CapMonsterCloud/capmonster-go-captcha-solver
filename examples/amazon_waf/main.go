// Usage: CAPMONSTER_API_KEY=... go run ./examples/amazon_waf
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

	solution, err := client.Solve(context.Background(), capmonster.AmazonWafVisibleRequest{
		WebsiteURL:    "https://example.com/",
		WebsiteKey:    "AQIDAHjcYu/GjX+QlghicBgQ/7bFaQZ+m5FKCMDnO+vTp9wjBQ2MNq0uv7gT8+IFxrGrLgsA",
		CaptchaScript: "https://234324vgvc23.yejk.captcha-sdk.awswaf.com/234324vgvc23/jsapi.js",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
