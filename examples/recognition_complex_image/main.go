// Usage: CAPMONSTER_API_KEY=... go run ./examples/recognition_complex_image <task> <image>... (e.g. oocl_rotate_double_new bg.png ring.png circle.png)
package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	capmonster "github.com/CapMonsterCloud/capmonster-go-captcha-solver/v2"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: recognition_complex_image <task> <image>...")
	}
	var images []string
	for _, path := range os.Args[2:] {
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		images = append(images, base64.StdEncoding.EncodeToString(data))
	}

	client := capmonster.New(os.Getenv("CAPMONSTER_API_KEY"))

	solution, err := client.Solve(context.Background(), capmonster.RecognitionComplexImageTaskRequest{
		ImagesBase64: images,
		Metadata:     capmonster.RecognitionMetadata{Task: os.Args[1]},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(solution))
}
