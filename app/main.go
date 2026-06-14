package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Status  string      `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{
		Status: "ok",
	})
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var body map[string]interface{}
	json.NewDecoder(r.Body).Decode(&body)

	json.NewEncoder(w).Encode(Response{
		Message: "received",
		Data:    body,
	})
}

func main() {
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/data", dataHandler)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
