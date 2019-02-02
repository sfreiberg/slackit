package slackit

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Payload holds the information that will be sent to slack
type Payload struct {
	Text      string `json:"text"`
	IconURL   string `json:"icon_url,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
	Username  string `json:"username,omitempty"`
}

// PostMessage posts the payload to slack
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
	return resp.Body.Close()
}
