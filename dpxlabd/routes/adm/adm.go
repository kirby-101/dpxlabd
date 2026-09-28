package adm

import "github.com/kirby-101/dpxlab/dpxlabd"

/*
 * /adm/login
 */

type AdmRoute struct {
	AllowedRoutes map[string]bool
}

func NewAdmRoute(routes map[string]dpxlabd.RouteHandler) *AdmRoute {
	var r = &AdmRoute{}

	return r
}

func (r *AdmRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {}
