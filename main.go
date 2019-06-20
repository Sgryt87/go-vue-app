package main

import (
	"CODE/goVue/src/system/app"
	"flag"
	"github.com/joho/godotenv"
	"os"
)

var port string

// will run when the package is called
func init() {
	flag.StringVar(&port, "port", "8000", "Server listening on this assigned port")
	flag.Parse()
	confFileName := "config.ini"
	if err := godotenv.Load(confFileName); err != nil {
		panic("Cannot load " + confFileName)
	}
	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {
	s := app.NewServer()
	s.Init(&port)
	s.Start()
}
