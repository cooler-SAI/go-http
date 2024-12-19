package main

import (
	"errors"
	"fmt"
	"go-http/tools"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: zerolog.ConsoleWriter{Out: os.Stdout}})

	log.Info().Msg("Hello go-http")
	log.Info().Time("timestamp", time.Now()).Msg("Server is starting")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello go-http")
		if err != nil {
			log.Error().Err(err).Msg("Error writing response to /")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "This is the About page of the Simple Go App.")
		if err != nil {
			log.Error().Err(err).Msg("Error writing response to /about")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		log.Info().Msg("Starting server on port 8080...")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Error starting server")
		}
	}()

	tools.WaitForShutdown(server)
}
