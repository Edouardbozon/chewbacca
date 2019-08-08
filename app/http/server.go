package http

import (
	"log"
	"net/http"
)

// DefaultAddr is the default bind address.
const DefaultAddr = ":8080"

// Server represents an HTTP server.type
type Server struct {
	Handler *Handler
	Addr    string
}

// NewServer returns a new instance of Server.
func NewServer() *Server {
	s := &Server{
		Addr:    DefaultAddr,
		Handler: NewHandler(),
	}

	return s
}

// ListenAndServe start the server
func (s *Server) ListenAndServe() {
	router := s.Handler.Router
	srv := &http.Server{
		Handler: router,
	}

	http.Handle("/", router)
	log.Fatal(srv.ListenAndServe())
}
