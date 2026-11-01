package index

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kirby-101/dpxlabd"
	"github.com/kirby-101/dpxlabd/routes"
)

type IndexRoute struct {
	contents []byte
	srv      *dpxlabd.Server
}

func New(srv *dpxlabd.Server) *IndexRoute {
	// Read & Store index.html
	contents, err := routes.AssetsFS.ReadFile(routes.HTMLPagePathIndex)
	if err != nil {
		panic(err)
	}

	return &IndexRoute{
		contents: contents,
		srv:      srv,
	}
}

func (route *IndexRoute) Register(router *mux.Router) *mux.Route {
	return router.HandleFunc("/", route.Handle).
		Name("index").
		Methods(http.MethodGet)
}

func (route *IndexRoute) Handle(w http.ResponseWriter, r *http.Request) {
	routes.SetContentTypeHTML(w)
	w.Write(route.contents)
}
