package routes

import (
	"backend/internal/core/login"
	"backend/internal/core/user"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) *mux.Router {
	routes := user.UserRoutes
	routes = append(routes, login.LoginRoutes)
	for _, route := range routes {
		router.HandleFunc(route.URI, route.Action).Methods(route.Method)
	}
	return router
}
