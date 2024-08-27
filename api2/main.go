package main

import (
	"api2/api"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all systems offline")
}

func run() error {
	apiKey := os.Getenv("OMDB_KEY")
	handler := api.NewHandler(apiKey)

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Minute,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
