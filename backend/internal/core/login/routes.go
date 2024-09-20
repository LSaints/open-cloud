package login

import (
	"backend/pkg/router/types"
	"net/http"
)

var LoginRoutes = types.Route{
	URI:                   "/login",
	Method:                http.MethodPost,
	Action:                Login,
	RequireAuthentication: false,
}
