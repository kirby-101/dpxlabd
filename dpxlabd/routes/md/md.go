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

func NewMarkdownRoute(_routes map[string]dpxlabd.RouteHandler) *MarkdownRoute {
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

			// Read Markdown File
			md_data, err := markdownFS.ReadFile(path)
			if err != nil {
				return err
			}

			// Convert Markdown to HTML & Execute BasePage Template
			html_data, err := markdown2HTML(md_data)
			if err != nil {
				return err
			}

			// content/asset.file => /assets/asset.file
			var routerPath = "/assets/" + strings.TrimPrefix(path, "content/")

			_routes[routerPath] = r
			r.AllowedRoutes[routerPath] = html_data

			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	//! generate index page

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
	site, oke := r.AllowedRoutes[route]
	if !oke {
		routes.Error(ctx, http.StatusInternalServerError)
	}

	ctx.W.Header().Set("Content-Type", routes.HTMLContentType)
	ctx.W.Write(site)
}
