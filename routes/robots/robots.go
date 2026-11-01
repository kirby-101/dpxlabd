package robots

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kirby-101/dpxlabd"
	"github.com/kirby-101/dpxlabd/routes"
)

const robotsTxtFilePath = "assets/robots.txt"

// * RFC 9309: /robots.txt
type RobotsTxtRoute struct {
	contents []byte
	srv      *dpxlabd.Server
}

func New(srv *dpxlabd.Server) *RobotsTxtRoute {
	// Read & Store /assets/robots.txt
	contents, err := routes.AssetsFS.ReadFile(robotsTxtFilePath)
	if err != nil {
		panic(err)
	}

	return &RobotsTxtRoute{
		contents: contents,
		srv:      srv,
	}
}

func (route *RobotsTxtRoute) Register(router *mux.Router) *mux.Route {
	return router.HandleFunc("/robots.txt", route.Handle).
		Name("robots.txt").
		Methods(http.MethodHead, http.MethodGet)
}

func (route *RobotsTxtRoute) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	if r.Method == http.MethodHead {
		return
	}

	routes.SetContentTypeText(w)
	w.Write(route.contents)
}
