package assets

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

type AssetsRoute struct {
	AllowedRoutes  map[string]bool
	AllowedMethods []string
}

func NewAssetsRoute(routes map[string]dpxlabd.RouteHandler) *AssetsRoute {
	var r = &AssetsRoute{
		AllowedMethods: []string{http.MethodGet},
		AllowedRoutes:  map[string]bool{},
	}

	fs.WalkDir(
		assetsFS,
		"content",
		func(path string, dir fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if dir.IsDir() {
				return nil
			}

			// content/asset.file => /assets/asset.file
			var routerPath = "/assets/" + strings.TrimPrefix(path, "content/")
			routes[routerPath] = r
			r.AllowedRoutes[routerPath] = true

			return nil
		},
	)

	return r
}

func (r *AssetsRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	if !routes.AssertHttpMethod(ctx, r.AllowedMethods) {
		return
	}

	// is valid route?
	if _, oke := r.AllowedRoutes[ctx.R.URL.Path]; !oke {
		routes.Error(ctx, http.StatusNotFound)
		return
	}

	catServeAssetsFS(ctx, ctx.R.URL.Path)
}
