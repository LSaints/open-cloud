package instance

import (
	"backend/pkg/router/types"
	"net/http"
)

var InstanceRoutes = []types.Route{
	{
		URI:                   "/instances",
		Method:                http.MethodPost,
		Action:                Create,
		RequireAuthentication: true,
	},
	{
		URI:                   "/instances/provision/{id}",
		Method:                http.MethodPost,
		Action:                ExecuteProvision,
		RequireAuthentication: true,
	},
	{
		URI:                   "/instances",
		Method:                http.MethodGet,
		Action:                GetAll,
		RequireAuthentication: true,
	},
	{
		URI:                   "/instances/{id}",
		Method:                http.MethodGet,
		Action:                Get,
		RequireAuthentication: true,
	},
	{
		URI:                   "/instances/{id}",
		Method:                http.MethodPut,
		Action:                Update,
		RequireAuthentication: true,
	},
	{
		URI:                   "/instances/{id}",
		Method:                http.MethodDelete,
		Action:                Delete,
		RequireAuthentication: true,
	},
	{
		URI:                   "/instances/halt/{id}",
		Method:                http.MethodPatch,
		Action:                HaltInstance,
		RequireAuthentication: true,
	},
	{
		URI:                   "/instances/status/{id}",
		Method:                http.MethodGet,
		Action:                GetStatusInstance,
		RequireAuthentication: true,
	},
	{
		URI:                   "/instances/start/{id}",
		Method:                http.MethodPatch,
		Action:                StartInstance,
		RequireAuthentication: true,
	},
}
