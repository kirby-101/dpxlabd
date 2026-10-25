package auth

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
)

/*
 * server gibt token mit signatur zurück
 */

// login.html braucht cookie banner bzw akzeptieren
type AuthRoute struct {
	activeChallenges map[uuid.UUID]*PoWChallenge
}

func NewAdmAuthRoute() *AuthRoute {}

func (r *AuthRoute) Route() string {
	return "/adm/auth"
}

// public
func (r *AuthRoute) AccessLevel() Role {
	return RoleInvalid
}

// server site? ; serve pow-challenge ; POST:login=>token
func (r *AuthRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	// assert correct route
	if ctx.R.URL.Path != "/adm/auth" {
		routes.Error(ctx, http.StatusNotFound)

		return
	}

	// assert correct http method
	switch ctx.R.Method {
	case http.MethodGet:
		r.HandleGet(ctx)
		return

	case http.MethodPost:
		r.HandlePostLogin(ctx)
		return

	default:
		routes.Error(ctx, http.StatusMethodNotAllowed)
		return
	}
}

func (r *AuthRoute) HandleGet(ctx *dpxlabd.RouteHandlerContext) {
	switch {
	case ctx.R.URL.Query().Has("pow"):
		r.HandleGetPoWChallenge(ctx)
		return

	case ctx.R.URL.Query().Has("logout"):
		r.HandleGetLogout(ctx)
	}

	// serve HTML page
	routes.CatServeHTMLFS(ctx, routes.LoginPagePath)
}
