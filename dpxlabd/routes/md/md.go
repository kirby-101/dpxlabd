package md

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

type MarkdownRoute struct {
	AllowedMethods []string
	AllowedRoutes  map[string][]byte
}

// ! convert md to html
func NewMarkdownRoute(routes map[string]dpxlabd.RouteHandler) *MarkdownRoute {
	var r = &MarkdownRoute{
		AllowedMethods: []string{http.MethodGet},
		AllowedRoutes:  map[string][]byte{},
	}

	err := fs.WalkDir(
		markdownFS,
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
			r.AllowedRoutes[routerPath] = []byte{} //! content

			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	return r
}

func (r *MarkdownRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	if !routes.AssertHttpMethod(ctx, r.AllowedMethods) {
		return
	}

	var route = ctx.R.URL.Path

	// is valid route?
	if _, oke := r.AllowedRoutes[route]; !oke {
		routes.Error(ctx, http.StatusNotFound)

		return
	}

	// serve markdown

	//catServeMarkdownFS(ctx, ctx.R.URL.Path)
}
