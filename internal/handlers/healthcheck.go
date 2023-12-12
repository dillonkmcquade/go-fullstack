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
	}
	w.Write(pong)
}
