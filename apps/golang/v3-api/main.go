package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Item struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {
	http.HandleFunc("/", homeRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Item{
		ID:          1,
		Description: "Item 4854857 ai papai",
	})
}
