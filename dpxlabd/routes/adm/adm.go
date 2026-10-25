package adm

import (
	"net/http"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/adm/auth"
)

type AdmRoute struct {
	r

	AllowedRoutes map[string]func(*dpxlabd.RouteHandlerContext)
}

func NewAdmRoute(_routes map[string]dpxlabd.RouteHandler) *AdmRoute {
	var r = &AdmRoute{
		// Obfuscate & Register AllowedRoutes
		AllowedRoutes: routes.ObfuscateRoutes(
			map[string]func(*dpxlabd.RouteHandlerContext){
				"/adm":        r.HandleDashboard,
				"/adm/auth":   auth.NewAdmAuthRoute().Handle,
				"/adm/login":  r.HandleLogin,
				"/adm/logout": r.HandleLogout,
			},
		),
	}

	return r
}

func (r *AdmRoute) Route() string {
	return "/adm"
}

func (r *AdmRoute) AccessLevel() auth.Role {
	return auth.RoleInvalid
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
