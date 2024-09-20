package router

import (
	"backend/pkg/router/routes"

	"github.com/gorilla/mux"
)

func Initialize() *mux.Router {
	router := mux.NewRouter()
	return routes.Setup(router)
}
