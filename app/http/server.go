package http

import (
	"fmt"
	"log"
	"net/http"
)

// Server represents an HTTP server.type
type Server struct {
	Handler *Handler
	Addr    string
}

// NewServer returns a new instance of Server.
func NewServer(port string) *Server {
	s := &Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: NewHandler(),
	}

	return s
}

// ListenAndServe start the server
func (s *Server) ListenAndServe() {
	log.Printf("[HTTP server] listening at %s", s.Addr)
	log.Fatal(http.ListenAndServe(s.Addr, s.Handler.Router))
}
