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

func NewAssetsRoute(routes map[string]dpxlabd.RouteHandler) *AssetsRoute {
	var r = &AssetsRoute{
		AllowedMethods: []string{http.MethodGet},
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

			return nil
		},
	)

	return r
}

func (r *AssetsRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	if !routes.AssertHttpMethod(ctx, r.AllowedMethods) {
		return
	}

	e, _ := assetsFS.ReadDir("content")

	for _, dir := range e {
		ctx.Orch.Logger.Debug(dir.Name())
	}

	ctx.Orch.Logger.Debug(ctx.R.URL.Path, fmt.Sprintf("%#v", e))

	// is valid route?

	catServeAssetsFS(ctx, ctx.R.URL.Path)
}
