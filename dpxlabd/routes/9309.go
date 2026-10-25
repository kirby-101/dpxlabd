package routes

import (
	"net/http"

	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/adm/auth"
)

const robotsTxtPagePath = "html/robots.txt"

// * RFC 9309: /robots.txt
type RobotsTxtRoute struct {
	AllowedMethods []string
}

func NewRobotsTxtRoute() *RobotsTxtRoute {
	return &RobotsTxtRoute{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
		},
	}
}

func (r *RobotsTxtRoute) Route() string {
	return "/robots.txt"
}

// public
func (r *RobotsTxtRoute) AccessLevel() auth.Role {
	return auth.RoleInvalid
}

func (r *RobotsTxtRoute) Handle(ctx *dpxlabd.RouteHandlerContext) {
	if !AssertHttpMethod(ctx, r.AllowedMethods) {
		return
	}

	ctx.W.Header().Set("Content-Type", "text/plain; charset=utf-8")
	ctx.W.WriteHeader(http.StatusOK)

	// HEAD /robots.txt
	if ctx.R.Method == http.MethodHead {
		return
	}

	// serve: GET /robots.txt
	CatServeFS(ctx, htmlFS, robotsTxtPagePath)
}
