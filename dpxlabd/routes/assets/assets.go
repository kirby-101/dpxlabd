package assets

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"

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

func (r *AssetsRoute) RegisterAssetsRoutes(routes map[string]dpxlabd.RouteHandler) {
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

			return nil
		},
	)
}

func (r *AssetsRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	if !routes.AssertHttpMethod(ctx, r.AllowedMethods) {
		return
	}

	e, _ := assetsFS.ReadDir("/")
	ctx.Orch.Logger.Debug(fmt.Sprintf("%#v", e))

	catServeAssetsFS(ctx, ctx.R.URL.Path)
}
