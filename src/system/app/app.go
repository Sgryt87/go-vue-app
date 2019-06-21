package app

import (
	"CODE/goVue/src/system/router"
	"github.com/go-xorm/xorm"
	"log"
	"net/http"
)

type Server struct {
	port string
	Db   *xorm.Engine
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Init(port *string, db *xorm.Engine) {
	log.Println("Initializing a server...")
	s.port = ":" + *port
	s.Db = db
}

func (s *Server) Start() {
	log.Println("Server started on port: " + s.port)
	r := router.NewRouter()
	r.Init()
	http.ListenAndServe(s.port, r.Router)
}
