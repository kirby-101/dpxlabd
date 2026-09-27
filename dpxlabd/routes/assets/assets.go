package assets

import (
	"fmt"
	"net/http"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

type AssetsRoute struct {
	AllowedMethods []string
}

func NewAssetsRoute() *AssetsRoute {
	return &AssetsRoute{
		AllowedMethods: []string{http.MethodGet},
	}
}

func (r *AssetsRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	if !routes.AssertHttpMethod(ctx, r.AllowedMethods) {
		return
	}

	e, _ := assetsFS.ReadDir("/")
	ctx.Orch.Logger.Debug(fmt.Sprintf("%#v", e))

	catServeAssetsFS(ctx, ctx.R.URL.Path)
}
