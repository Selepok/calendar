package response

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Token struct {
	Token string `json:"token"`
}

type Error struct {
	Error string `json:"error"`
}

func Respond(w http.ResponseWriter, status int, data interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("respond:", err)
	}
}
