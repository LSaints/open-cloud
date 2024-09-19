package routes

import (
	"backend/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Action:                controllers.Create,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Action:                controllers.GetAll,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodGet,
		Action:                controllers.Get,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodPut,
		Action:                controllers.Update,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodDelete,
		Action:                controllers.Delete,
		RequireAuthentication: false,
	},
}
