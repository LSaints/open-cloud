package router

import (
	"backend/src/router/routes"

	"github.com/gorilla/mux"
)

func Initialize() *mux.Router {
	router := mux.NewRouter()
	return routes.Setup(router)
}
