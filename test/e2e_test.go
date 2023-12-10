package test

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	r, err := http.Get("http://localhost:3001/ping")
	if err != nil {
		t.Error(err)
	}
	defer r.Body.Close()

	var exp map[string]string
	err = json.NewDecoder(r.Body).Decode(&exp)

	if v, ok := exp["message"]; v != "pong" || !ok {
		t.Error("The response did not contain the expected message")
	}

	if r.StatusCode != 200 {
		t.Error("Status not OK")
	}
}

func TestIndex(t *testing.T) {
	r, err := http.Get("http://localhost:3001/")
	if err != nil {
		t.Error(err)
	}
	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "text/html; charset=utf-8" {
		t.Error("Invalid content type")
	}

	if r.StatusCode != 200 {
		t.Error("Status not OK")
	}
}
