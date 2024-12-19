package main

import (
	"errors"
	"fmt"
	"go-http/tools"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP Requests",
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	log.Info().Msg("Hello go-http")
	log.Info().Time("timestamp", time.Now()).Msg("Server is starting")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpRequestsTotal.WithLabelValues("/").Inc()
		_, err := fmt.Fprintf(w, "Hello go-http")
		if err != nil {
			log.Error().Err(err).Msg("Error writing response to /")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		httpRequestsTotal.WithLabelValues("/about").Inc()
		_, err := fmt.Fprintln(w, "This is the About page of the Simple Go App.")
		if err != nil {
			log.Error().Err(err).Msg("Error writing response to /about")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	http.Handle("/metrics", promhttp.Handler())

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

	go func() {
		tools.WaitForShutdown(server)
		log.Info().Msg("Server is shutting down")
		os.Exit(0)
	}()

	select {}
}
