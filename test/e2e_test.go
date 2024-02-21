package test

import (
	"encoding/json"
	"net/http"
	"testing"
)

// *** End to End tests ***
// *** Server must be running in order to use these tests ***

func TestHealthCheck(t *testing.T) {
	r, err := http.Get("http://localhost:3001/ping")
	if err != nil {
		t.Error(err)
	}
	defer r.Body.Close()

	var exp map[string]string
	err = json.NewDecoder(r.Body).Decode(&exp)
	if err != nil {
		t.Error("JSON decoding error")
	}

	if v, ok := exp["message"]; v != "pong" || !ok {
		t.Error("response should be {'message': 'pong'}")
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
		t.Error("Content type should be text/html")
	}

	if r.StatusCode != 200 {
		t.Error("Status not OK")
	}
}
