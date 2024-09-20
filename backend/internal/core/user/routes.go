package user

import (
	"backend/pkg/router/types"
	"net/http"
)

var UserRoutes = []types.Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Action:                Create,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Action:                GetAll,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodGet,
		Action:                Get,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodPut,
		Action:                Update,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodDelete,
		Action:                Delete,
		RequireAuthentication: false,
	},
}
