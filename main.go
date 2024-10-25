package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/yKanazawa/sendgrid-dev/route"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		arg := args[1]
		if arg == "health" {
			resp, err := http.Get("http://localhost:3030/health")
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				os.Exit(1)
			}
			defer func() {
				_ = resp.Body.Close()
			}()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				os.Exit(1)
			}
			if resp.StatusCode != http.StatusOK {
				fmt.Printf("status=%s, body=%s\n", resp.Status, string(body))
				os.Exit(1)
			}
		}
		return
	}
	if os.Getenv("SENDGRID_DEV_API_SERVER") == "" {
		os.Setenv("SENDGRID_DEV_API_SERVER", ":3030")
	}
	fmt.Println("SENDGRID_DEV_API_SERVER", os.Getenv("SENDGRID_DEV_API_SERVER"))

	if os.Getenv("SENDGRID_DEV_API_KEY") == "" {
		os.Setenv("SENDGRID_DEV_API_KEY", "SG.xxxxx")
	}
	fmt.Println("SENDGRID_DEV_API_KEY", os.Getenv("SENDGRID_DEV_API_KEY"))

	if os.Getenv("SENDGRID_DEV_SMTP_SERVER") == "" {
		os.Setenv("SENDGRID_DEV_SMTP_SERVER", "127.0.0.1:1025")
	}
	fmt.Println("SENDGRID_DEV_SMTP_SERVER", os.Getenv("SENDGRID_DEV_SMTP_SERVER"))
	fmt.Println("SENDGRID_DEV_SMTP_USERNAME", os.Getenv("SENDGRID_DEV_SMTP_USERNAME"))
	fmt.Println("SENDGRID_DEV_SMTP_PASSWORD", os.Getenv("SENDGRID_DEV_SMTP_PASSWORD"))

	router := route.Init()
	router.Logger.Fatal(router.Start(os.Getenv("SENDGRID_DEV_API_SERVER")))
}
