package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"student-information-system/internal/api"
	"syscall"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: api.Router(),
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Server running on port: 8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf(
				"Something is wrong... shutting down server.\n"+
					"ERROR: [command/app/main.go main()] -> %v", err,
			)
		}
	}()

	<-interrupt
	log.Println("Server getting down...")

	_ = server.Shutdown(nil)
	log.Println("Server is down.")
}
