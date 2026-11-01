package dpxlabd

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type WWW struct {
	Router *mux.Router
	srv    *http.Server
}

func (srv *Server) StartWWW(errChannel chan error) {
	if !srv.Config.WWW.Enabled {
		srv.Logger.Info("WWW disabled")
		return
	}

	srv.Logger.Info("trying to listen on", srv.Config.WWW.Address)

	errChannel <- srv.WWW.srv.ListenAndServe()
}

func (srv *Server) StopWWW(ctx context.Context) error {
	// is running?
	return srv.WWW.srv.Shutdown(ctx)
}

/*
func (srv *Server) configureRouter() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", initFileserver(srv.Config.DocumentServerHost)))

	r.HandleFunc("/", srv.ServerAPI.Index).Methods(http.MethodGet)
	r.HandleFunc("/editor", srv.ServerAPI.Editor).Methods(http.MethodGet)
	r.HandleFunc("/callback", srv.ServerAPI.Callback).Methods(http.MethodPost)
	r.HandleFunc("/remove", srv.ServerAPI.Remove).Methods(http.MethodGet)
	r.HandleFunc("/remove", srv.ServerAPI.Remove).Methods(http.MethodDelete)
	r.HandleFunc("/upload", srv.ServerAPI.Upload).Methods(http.MethodPost)
	r.HandleFunc("/convert", srv.ServerAPI.Convert).Methods(http.MethodPost)
	r.HandleFunc("/download", srv.ServerAPI.Download).Methods(http.MethodGet)
	r.HandleFunc("/history", srv.ServerAPI.History).Methods(http.MethodGet)
	r.HandleFunc("/create", srv.ServerAPI.Create).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/reference", srv.ServerAPI.Reference).Methods(http.MethodPost)
	r.HandleFunc("/files", srv.ServerAPI.Files).Methods(http.MethodGet)
	r.HandleFunc("/rename", srv.ServerAPI.Rename).Methods(http.MethodPost)
	r.HandleFunc("/historyObj", srv.ServerAPI.HistoryObj).Methods(http.MethodPost)
	r.HandleFunc("/restore", srv.ServerAPI.Restore).Methods(http.MethodPut)
	r.HandleFunc("/formats", srv.ServerAPI.Formats).Methods(http.MethodGet)
	r.HandleFunc("/config", srv.ServerAPI.Config).Methods(http.MethodGet)
	r.HandleFunc("/forgotten", srv.ServerAPI.Forgotten).Methods(http.MethodGet, http.MethodDelete)

	return r
}
*/
