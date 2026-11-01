package static

import (
	"io/fs"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kirby-101/dpxlabd"
	"github.com/kirby-101/dpxlabd/routes"
)

type StaticRoute struct {
	subroutes map[string][]byte
	srv       *dpxlabd.Server
}

func New(srv *dpxlabd.Server) *StaticRoute {
	var self = &StaticRoute{
		subroutes: map[string][]byte{},
		srv:       srv,
	}

	err := fs.WalkDir(
		routes.AssetsFS,
		routes.StaticFilesPath,
		self.enumerateSubroutes,
	)
	if err != nil {
		panic(err)
	}

	return self
}

func (route *StaticRoute) enumerateSubroutes(path string, dir fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if dir.IsDir() {
		return nil
	}

	// Read Static File
	contents, err := routes.AssetsFS.ReadFile(path)
	if err != nil {
		return err
	}

	// assets/static/static.file => /static/static.file
	var subroute = "/" + strings.TrimPrefix(path, "assets/")

	// Store Static File
	route.subroutes[subroute] = contents

	return nil
}

func (route *StaticRoute) Register(router *mux.Router) *mux.Route {
	return router.PathPrefix("/static").
		Name("static").
		Methods(http.MethodGet).
		HandlerFunc(route.Handle)
}

func (route *StaticRoute) Handle(w http.ResponseWriter, r *http.Request) {
	var subroute = r.URL.Path

	// valid subroute?
	contents, oke := route.subroutes[subroute]
	if !oke {
		routes.Error(http.StatusNotFound).ServeHTTP(w, r)

		return
	}

	// Set Content-Type
	if extension := mime.TypeByExtension(filepath.Ext(subroute)); extension != "" {
		w.Header().Set("Content-Type", extension)
	}

	// Serve File
	w.Write(contents)
}
