package assets

import (
	"io/fs"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

type AssetsRoute struct {
	AllowedMethods []string
	AllowedRoutes  map[string]bool
}

func NewAssetsRoute(routes map[string]dpxlabd.RouteHandler) *AssetsRoute {
	var r = &AssetsRoute{
		AllowedMethods: []string{http.MethodGet},
		AllowedRoutes:  map[string]bool{},
	}

	err := fs.WalkDir(
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
	if err != nil {
		panic(err)
	}

	return r
}

func (r *AssetsRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	if !routes.AssertHttpMethod(ctx, r.AllowedMethods) {
		return
	}

	var route = ctx.R.URL.Path

	// is valid route?
	if _, oke := r.AllowedRoutes[route]; !oke {
		routes.Error(ctx, http.StatusNotFound)

		return
	}

	if ext := mime.TypeByExtension(filepath.Ext(route)); ext != "" {
		ctx.W.Header().Set("Content-Type", ext)
	}

	catServeAssetsFS(ctx, ctx.R.URL.Path)
}
