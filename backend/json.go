package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error parsing to json format")
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "applications/json")
	w.WriteHeader(code)
	w.Write(data)
}
