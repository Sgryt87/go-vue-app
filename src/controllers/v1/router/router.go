package router

import (
	"CODE/goVue/pkg/types/routes"
	StatusHandler "CODE/goVue/src/controllers/v1/status"
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("V1 middleware reached!")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() (SubRoute map[string]routes.SubRoutePackage) {
	// ROUTES
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": routes.SubRoutePackage{
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
			},
			Middleware: Middleware,
		},
	}
	return
}
