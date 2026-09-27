package dpxlabd

import (
	"fmt"
	"net/http"
)

func (o *Orchestrator) StartWWW() error {
	o.Logger.Info("Listening on", o.cfg.WWW.Address)
	o.www.mux.HandleFunc("/", o.HandleRoute)

	o.Logger.Debug("routes", fmt.Sprintf("%#v", o.www.routes))

	return http.ListenAndServe(o.cfg.WWW.Address, o.www.mux)
}

// Handle(*RouteHandlerContext) error
func (o *Orchestrator) RegisterRoute(route string, handler RouteHandler) {
	o.www.routes[route] = handler
}

func (o *Orchestrator) HandleRoute(w http.ResponseWriter, r *http.Request) {
	// Default RouteHandler
	var h = o.www.errR

	// exact match?
	if _h, oke := o.www.routes[r.URL.Path]; oke {
		// 200
		h = _h
		o.Logger.Info(r.RemoteAddr, trunc(r.UserAgent()), "=>", r.Method, "using handler", r.URL.Path)
	} else {
		// 404
		o.Logger.Debug(r.RemoteAddr, trunc(r.UserAgent()), "=>", r.Method, "unknown handler", r.URL.Path)
	}

	h.Handle(
		&RouteHandlerContext{
			W:    w,
			R:    r,
			Orch: o,
		},
	)
}
