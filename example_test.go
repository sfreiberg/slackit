package slackit_test

import (
	"log"
	"os"

	"github.com/sfreiberg/slackit"
)

func ExamplePostMessage() {
	url := os.Getenv("SLACK_URL")
	payload := &slackit.Payload{Text: "Hello from slackit!"}

	err := slackit.PostMessage(url, payload)
	if err != nil {
		log.Printf("Error posting to slack: %s\n", err)
	}
}
