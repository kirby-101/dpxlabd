package adm

import (
	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/adm/auth"
)

/*
   nur logged in
   /adm/proxy?s=10.0.0.1:8080
*/

type ProxyRoute struct{}

/*
type RouteHandler interface {
	Route() string
	AccessLevel() auth.Role
	Handle(*RouteHandlerContext)
}
*/

func NewProxyRoute(_routes map[string]dpxlabd.RouteHandler) *ProxyRoute {
	return &ProxyRoute{}
}

func (r *ProxyRoute) Route() string {
	return "/adm/proxy"
}

// needs admin privs
func (r *ProxyRoute) AccessLevel() auth.Role {
	return auth.RoleAdmin
}

func (r *ProxyRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	// valid args?

	// target?

	// serve
}
