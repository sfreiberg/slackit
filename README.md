# slackit

[![GoDoc](https://godoc.org/github.com/sfreiberg/slackit?status.png)](https://godoc.org/github.com/sfreiberg/slackit)

## About

Slackit is a very simple library for posting messages to slack. I use this code all over my projects so even though it's simple I thought it would be useful to put it into a library.

## Example

```go
url := os.Getenv("SLACK_URL")
payload := &slackit.Payload{Text: "Hello from slackit!"}

err := slackit.PostMessage(url, payload)
if err != nil {
    log.Printf("Error posting to slack: %s\n", err)
}
```