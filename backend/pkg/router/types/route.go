package types

import "net/http"

type Route struct {
	URI                   string
	Method                string
	Action                func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}
