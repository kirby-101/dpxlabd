package dpxlabd

import (
	"net/http"

	"github.com/kirby-101/dpxlab/dpxlabd/routes"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/adm/auth"
)

func (o *Orchestrator) StartWWW() error {
	o.Logger.Info("WWW: Listening Socket", o.cfg.WWW.Address)
	o.www.mux.HandleFunc("/", o.HandleRoute)

	var err = http.ListenAndServe(o.cfg.WWW.Address, o.www.mux)

	return err
}

// Handle(*RouteHandlerContext) error
func (o *Orchestrator) RegisterRoute(route string, handler RouteHandler) {
	o.www.routes[route] = handler
	o.Logger.Debug("registered", route)
}

func (o *Orchestrator) HandleRoute(w http.ResponseWriter, r *http.Request) {
	// Default RouteHandler
	var h = o.www.errR

	// exact match?
	if _h, oke := o.www.routes[r.URL.Path]; oke {
		// 200
		h = _h
		o.Logger.Info(r.RemoteAddr, trunc(r.UserAgent()), r.Method, r.URL.Path)
	} else {
		// 404
		o.Logger.Debug(r.RemoteAddr, trunc(r.UserAgent()), "unknown handler", r.Method, r.URL.Path)
	}

	// is authorized to handle?
	if auth.GetClientToken().Role != h.AccessLevel() {
		routes.Error(ctx, http.StatusForbidden)

		return
	}

	// serve
	h.Handle(
		&RouteHandlerContext{
			W:    w,
			R:    r,
			Orch: o,
		},
	)
}
