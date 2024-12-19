package main

import (
	"errors"
	"fmt"
	"go-http/tools"
	"log"
	"net/http"
	"time"
)

func main() {
	zerolo
	fmt.Println("Hello go-http")
	fmt.Println(time.Now())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello go-http")
		if err != nil {
			log.Printf("Error writing response to /: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "This is the About page of the Simple Go App.")
		if err != nil {
			log.Printf("Error writing response to /about: %v\n", err)
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
		log.Println("Starting server on port 8080...")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	tools.WaitForShutdown(server)
}
