package slackit_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/sfreiberg/slackit"
)

func TestPostMessage(t *testing.T) {
	url := os.Getenv("SLACK_URL")
	if url == "" {
		t.Skip("SLACK_URL not set; skipping integration test.")
	}
	payload := &slackit.Payload{Text: "Hello from slackit!"}

	if err := slackit.PostMessage(url, payload); err != nil {
		t.Fatalf("Error posting to slack: %s", err)
	}
}

func TestPostMessageSendsPayload(t *testing.T) {
	var got slackit.Payload
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", ct)
		}
		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Errorf("decoding request body: %v", err)
		}
		w.Write([]byte("ok"))
	}))
	defer srv.Close()

	want := slackit.Payload{Text: "hi", Username: "bot"}
	if err := slackit.PostMessage(srv.URL, &want); err != nil {
		t.Fatalf("PostMessage() error: %v", err)
	}
	if got != want {
		t.Errorf("server received %+v, want %+v", got, want)
	}
}

func TestPostMessageErrorStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "invalid_payload", http.StatusBadRequest)
	}))
	defer srv.Close()

	err := slackit.PostMessage(srv.URL, &slackit.Payload{Text: "hi"})
	if err == nil {
		t.Fatal("PostMessage() with a 400 response: expected an error")
	}
	if !strings.Contains(err.Error(), "400") || !strings.Contains(err.Error(), "invalid_payload") {
		t.Errorf("error %q should include the status and slack's reason", err)
	}
}
