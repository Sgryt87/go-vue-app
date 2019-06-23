package router

import (
	"CODE/goVue/pkg/types/routes"
	HomeHandler "CODE/goVue/src/controllers/home"
	"github.com/go-xorm/xorm"
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Global middleware reached!")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) routes.Routes {
	HomeHandler.Init(db)
	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
	}
}
