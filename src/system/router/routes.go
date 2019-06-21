package router

import (
	"CODE/goVue/pkg/types/routes"
	HomeHandler "CODE/goVue/src/controllers/home"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() routes.Routes {

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
	}
}
