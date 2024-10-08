package routes

import (
	"backend/internal/core/instance"
	"backend/internal/core/login"
	"backend/internal/core/user"
	"backend/pkg/middlewares"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) *mux.Router {
	routes := user.UserRoutes
	routes = append(routes, login.LoginRoutes)
	routes = append(routes, instance.InstanceRoutes...)

	for _, route := range routes {
		if route.RequireAuthentication {
			router.HandleFunc(route.URI, middlewares.Authenticate(route.Action)).Methods(route.Method)
		}
		router.HandleFunc(route.URI, middlewares.Logger(route.Action)).Methods(route.Method)
	}

	return router
}
