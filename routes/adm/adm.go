package adm

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kirby-101/dpxlabd"
)

type AdmRoute struct {
	srv *dpxlabd.Server
}

func New(srv *dpxlabd.Server) *AdmRoute {
	return &AdmRoute{srv: srv}
}

func (route *AdmRoute) Register(router *mux.Router) *mux.Route {
	return router.PathPrefix("/adm").
		Name("adm").
		Methods(http.MethodGet, http.MethodPost).
		HandlerFunc(route.Handle)
}

/*
 * /adm
 * /adm/auth
 * /adm/login
 * /adm/logout
 * /adm/dashboard
 * /adm/service
 * /adm/safe
 */

/*
- /adm/auth
  - GET /adm/login : login.html
  - GET /adm/auth : serve pow challenge
  - POST /adm/auth : submit pow, login creds => token
  - GET /adm/logout : unset token cookie & redirect zu /adm/login
*/
func (route *AdmRoute) Handle(w http.ResponseWriter, r *http.Request) {
	var subroute = r.URL.Path
	_ = subroute

	// subroute?
	// content type?
	// serve?
}
