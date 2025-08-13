package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Health-check placeholder
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

}
