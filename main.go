package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", Index)

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is not set")
	}

	addr := fmt.Sprintf(":%s", port)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	server.ListenAndServe()
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!"))
}
