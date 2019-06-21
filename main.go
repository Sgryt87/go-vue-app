package main

import (
	"CODE/goVue/src/system/app"
	DB "CODE/goVue/src/system/db"
	"flag"
	"github.com/joho/godotenv"
	"os"
)

var port string
var dbhost string
var dbport string
var dbuser string
var dbpass string
var dboptions string
var dbdatabase string

// will run when the package is called
func init() {
	flag.StringVar(&port, "port", "8000", "Set the server connection")
	flag.StringVar(&dbhost, "dbhost", "127.0.0.1", "Set the port of the application")
	flag.StringVar(&dbport, "dbport", "3306", "Set the port of the database")
	flag.StringVar(&dbuser, "dbuser", "root", "DB User")
	flag.StringVar(&dbpass, "dbpass", "root", "DB Pass")
	flag.StringVar(&dboptions, "dboptions", "parseTime=true", "DB options")
	flag.StringVar(&dbdatabase, "dbdatabase", "go_vue", "DB name")
	flag.Parse()

	confFileName := "config.ini"
	if err := godotenv.Load(confFileName); err != nil {
		println("Cannot load " + confFileName)
		panic(err)
	}

	if host := os.Getenv("DB_HOST"); len(host) > 0 {
		dbhost = host
	}

	if database := os.Getenv("DB_DATABASE"); len(database) > 0 {
		dbdatabase = database
	}

	if user := os.Getenv("DB_DATABASE"); len(user) > 0 {
		dbuser = user
	}
	if password := os.Getenv("DB_DATABASE"); len(password) > 0 {
		dbpass = password
	}
	if port := os.Getenv("DB_DATABASE"); len(port) > 0 {
		dbport = port
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {
	db, err := DB.Connect(&dbhost, &dbport, &dbuser, &dbpass, &dbdatabase, &dboptions)
	if err != nil {
		panic(err)
	}

	s := app.NewServer()
	s.Init(&port, db)
	s.Start()
}
