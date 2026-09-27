package dpxlabd

import (
	"net/http"
)

func (o *Orchestrator) StartWWW() error {
	o.Logger.Info("Listening on", o.cfg.WWW.Address)
	o.www.mux.HandleFunc("/", o.HandleRoute)
	return http.ListenAndServe(o.cfg.WWW.Address, o.www.mux)
}

// Handle(*RouteHandlerContext) error
func (o *Orchestrator) RegisterRoute(route string, handler RouteHandler) {
	o.www.routes[route] = handler
}

// session?; log
func (o *Orchestrator) HandleRoute(w http.ResponseWriter, r *http.Request) {
	o.Logger.Info(r.RemoteAddr, trunc(r.UserAgent()), "=>", r.Method, r.URL.Path)

	var h RouteHandler

	if _h, oke := o.www.routes[r.URL.Path]; oke {
		h = _h
		return
	} else {
		h = o.www.errR
	}

	h.Handle(
		&RouteHandlerContext{
			W:    w,
			R:    r,
			Orch: o,
		},
	)

}
