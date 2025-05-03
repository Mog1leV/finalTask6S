package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	server := server.NewServer(logger)

	if err := server.RunServer(); err != nil {
		server.Logger.Fatalf("error starting server: %s", err)
	}
}
