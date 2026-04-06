package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	host := getEnv("APP_HOST", "0.0.0.0")
	port := getEnv("APP_PORT", "8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"service":"nexo","status":"ok"}`))
	})

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("starting nexo on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
