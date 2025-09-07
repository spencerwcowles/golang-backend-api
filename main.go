package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/spencerwcowles/golang-backend-api/middleware"
)

func handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router := http.NewServeMux()

	router.HandleFunc("/", handleTest)

	srv := http.Server{
		Addr:    ":8000",
		Handler: middleware.Logging(logger, router),
	}

	srv.ListenAndServe()
}
