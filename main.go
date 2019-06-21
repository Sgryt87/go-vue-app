package main

import (
	"CODE/goVue/src/system/app"
	DB "CODE/goVue/src/system/db"
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
	if db, err := DB.Connect(); err != nil {
		panic(err)
	}

	s := app.NewServer()
	s.Init(&port, db)
	s.Start()
}
