package tools

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func WaitForShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("\nShutting down server...")

	if err := server.Close(); err != nil {
		fmt.Printf("Error during shutdown: %v\n", err)
	} else {
		fmt.Println("Server stopped gracefully.")
	}
}
