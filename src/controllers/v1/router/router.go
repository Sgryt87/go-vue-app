package router

import (
	"CODE/goVue/pkg/types/routes"
	StatusHandler "CODE/goVue/src/controllers/v1/status"
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// token verification
		token := r.Header.Get("X-App-Token")
		if len(token) < 1 {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}
		log.Println("V1 middleware reached!")
		log.Println("TOKEN->", token)
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
