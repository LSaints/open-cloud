package routes

import (
	"backend/src/features/user"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) *mux.Router {
	routes := user.UserRoutes
	for _, route := range routes {
		router.HandleFunc(route.URI, route.Action).Methods(route.Method)
	}
	return router
}
