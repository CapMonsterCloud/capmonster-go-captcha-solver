// Usage: CAPMONSTER_API_KEY=... go run ./examples/concurrent [N]
//
// Solves N reCAPTCHA v2 tasks one by one and then in parallel goroutines, and compares the throughput.
// A Client is safe for concurrent use.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	capmonster "github.com/CapMonsterCloud/capmonster-go-captcha-solver/v2"
)

func main() {
	n := 3
	if len(os.Args) > 1 {
		v, err := strconv.Atoi(os.Args[1])
		if err != nil || v < 1 {
			log.Fatal("N must be a positive integer")
		}
		n = v
	}

	client := capmonster.New(os.Getenv("CAPMONSTER_API_KEY"))
	ctx := context.Background()
	req := capmonster.RecaptchaV2Request{
		WebsiteURL: "https://lessons.zennolab.com/captchas/recaptcha/v2_simple.php?level=high",
		WebsiteKey: "6Lcg7CMUAAAAANphynKgn9YAgA4tQ2KI_iqRyTwd",
	}

	start := time.Now()
	for i := 0; i < n; i++ {
		if _, err := client.Solve(ctx, req); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("sequential: %.2f solves/sec\n", float64(n)/time.Since(start).Seconds())

	start = time.Now()
	errs := make([]error, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, errs[i] = client.Solve(ctx, req)
		}(i)
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("concurrent: %.2f solves/sec\n", float64(n)/time.Since(start).Seconds())
}
