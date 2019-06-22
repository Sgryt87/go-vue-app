package app

import (
	"CODE/goVue/src/system/router"
	"github.com/go-xorm/xorm"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"time"
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
	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATH", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-type", "Origin", "Cache-control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	newServer := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(newServer.ListenAndServe())
	//http.ListenAndServe(s.port, handler)
}
