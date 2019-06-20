package app

import (
	"log"
	"net/http"
)

type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Init(port *string) {
	log.Println("Initializing a server...")
	s.port = ":" + *port
}

func (s *Server) Start() {
	log.Println("Server listening on port: " + s.port)
	http.ListenAndServe(s.port, nil)
}
