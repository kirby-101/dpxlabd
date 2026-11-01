package markdown

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kirby-101/dpxlabd"
	"github.com/kirby-101/dpxlabd/routes"
)

type MarkdownRoute struct {
	subroutes map[string][]byte
	srv       *dpxlabd.Server
}

func New(srv *dpxlabd.Server) *MarkdownRoute {
	var route = &MarkdownRoute{
		subroutes: map[string][]byte{},
		srv:       srv,
	}

	err := fs.WalkDir(
		routes.AssetsFS,
		routes.MarkdownPagesPath,
		route.enumerateSubroutes,
	)
	if err != nil {
		panic(err)
	}

	return route
}

func (route *MarkdownRoute) enumerateSubroutes(path string, dir fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if dir.IsDir() {
		return nil
	}

	// Read Markdown File
	markdownContents, err := routes.AssetsFS.ReadFile(path)
	if err != nil {
		return err
	}

	// Markdown => HTML
	htmlContents, err := markdown2html(markdownContents)
	if err != nil {
		return err
	}

	// assets/markdown/file.md => /markdown/file.md

	var subroute = "/" + strings.TrimPrefix(
		strings.TrimSuffix(path, ".md"),
		"assets/",
	)

	// Store Markdown Page
	route.subroutes[subroute] = htmlContents

	return nil
}

func (route *MarkdownRoute) Register(router *mux.Router) *mux.Route {
	return router.PathPrefix("/markdown").
		Name("markdown").
		Methods(http.MethodGet).
		HandlerFunc(route.Handle)
}

func (route *MarkdownRoute) Handle(w http.ResponseWriter, r *http.Request) {
	var subroute = r.URL.Path

	// valid subroute?
	contents, oke := route.subroutes[subroute]
	if !oke {
		routes.Error(http.StatusNotFound)

		return
	}

	// Set Content-Type: HTML
	routes.SetContentTypeHTML(w)

	// Serve Markdown=>HTML Page
	w.Write(contents)
}
