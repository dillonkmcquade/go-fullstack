package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// GET /ping should return {"message": "pong"}
func Healthcheck(w http.ResponseWriter, r *http.Request) {
	pong, err := json.Marshal(map[string]string{"message": "pong"})
	if err != nil {
		log.Println(err)
		http.Error(w, "json marshal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(pong)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
