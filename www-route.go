package dpxlabd

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route interface {
	Register(*mux.Router) *mux.Route
	Handle(w http.ResponseWriter, r *http.Request)
}
