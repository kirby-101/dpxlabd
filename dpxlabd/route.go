package dpxlabd

import (
	"net/http"
)

type RouteHandlerContext struct {
	W    http.ResponseWriter
	R    *http.Request
	Orch *Orchestrator
}

type RouteHandler interface {
	Handle(*RouteHandlerContext)
}
