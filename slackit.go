package slackit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Payload holds the information that will be sent to slack
type Payload struct {
	Text      string `json:"text"`
	IconURL   string `json:"icon_url,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
	Username  string `json:"username,omitempty"`
}

// PostMessage posts the payload to slack. It returns an error if the request
// fails or slack responds with a non-2xx status code.
func PostMessage(url string, p *Payload) error {
	jsonPayload, err := json.Marshal(p)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	return post(req)
}

func post(req *http.Request) error {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Slack explains the failure in the body, e.g. "invalid_payload".
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return fmt.Errorf("slackit: %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}
	return nil
}
