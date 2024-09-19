package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	Action                func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func Setup(router *mux.Router) *mux.Router {
	routes := userRoutes
	for _, route := range routes {
		router.HandleFunc(route.URI, route.Action).Methods(route.Method)
	}
	return router
}
