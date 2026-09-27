package routes

import (
	"net/http"

	"github.com/kirby-101/dpxlab/dpxlabd"
)

/*
 * /login => /adm
 * pkg.dpxlab.de
 * pkgsched status
 * rad mirror?
 * github link
 * /md => (Impressum;etc)
 */

type IndexRoute struct {
	AllowedMethods []string
}

func NewIndexRoute() *IndexRoute {
	return &IndexRoute{
		AllowedMethods: []string{http.MethodGet},
	}
}

func (r *IndexRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	if AssertHttpMethod(ctx, r.AllowedMethods) {
		return
	}

	CatServeFS(ctx, htmlFS, indexPage)
}
