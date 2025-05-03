package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger     *log.Logger
	HTTPserver *http.Server
}

func (s *Server) RunServer() error {
	s.Logger.Println("Starting the Server!")
	return s.HTTPserver.ListenAndServe()
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	serv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HTTPserver: serv,
	}
}
