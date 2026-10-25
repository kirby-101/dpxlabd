package dpxlabd

import (
	"net/http"

	"github.com/kirby-101/dpxlab/dpxlabd/routes/adm/auth"
)

type RouteHandlerContext struct {
	W    http.ResponseWriter
	R    *http.Request
	Orch *Orchestrator
}

type RouteHandler interface {
	Route() string // prefix route
	AccessLevel() auth.Role
	Handle(*RouteHandlerContext)
}
