package routes

import (
	"net/http"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/adm/auth"
)

type IndexRoute struct {
	AllowedMethods []string
}

func NewIndexRoute() *IndexRoute {
	return &IndexRoute{
		AllowedMethods: []string{http.MethodGet},
	}
}

func (r IndexRoute) Route() string {
	return "/"
}

// public
func (r IndexRoute) AccessLevel() auth.Role {
	return auth.RoleInvalid
}

func (r *IndexRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	if !AssertHttpMethod(ctx, r.AllowedMethods) {
		return
	}

	CatServeHTMLFS(ctx, IndexPagePath)
}
