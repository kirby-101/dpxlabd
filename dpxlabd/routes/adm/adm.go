package adm

import (
	"net/http"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

type AdmRoute struct {
	AllowedRoutes map[string]func(*dpxlabd.RouteHandlerContext)
}

func NewAdmRoute(_routes map[string]dpxlabd.RouteHandler) *AdmRoute {
	var r = &AdmRoute{}

	// Obfuscate & Register AllowedRoutes
	r.AllowedRoutes = routes.ObfuscateRoutes(
		map[string]func(*dpxlabd.RouteHandlerContext){
			"/adm":        r.HandleDashboard,
			"/adm/login":  r.HandleLogin,
			"/adm/logout": r.HandleLogout,
		},
	)

	return r
}

func (r *AdmRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	var route = ctx.R.URL.Path

	// is valid route?
	h, oke := r.AllowedRoutes[route]
	if !oke {
		routes.Error(ctx, http.StatusNotFound)

		return
	}

	// serve
	h(ctx)
}
