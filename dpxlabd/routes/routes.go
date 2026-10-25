package routes

import (
	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/adm"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/assets"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/md"
)

var RouteTable = map[string]dpxlabd.RouteHandler{}

func init() {

	// hoffentlich alles dabei
	var (
		_indexRoute    = NewIndexRoute()
		_9309Route     = NewRobotsTxtRoute()
		_markdownRoute = md.NewMarkdownRoute(RouteTable)
		_assetsRoute   = assets.NewAssetsRoute(RouteTable)
		_admRoute      = adm.NewAdmRoute(RouteTable)
	)

	RouteTable = map[string]dpxlabd.RouteHandler{
		_indexRoute.Route():    _indexRoute,
		_9309Route.Route():     _9309Route,
		_markdownRoute.Route(): _markdownRoute,
		_assetsRoute.Route():   _assetsRoute,
		_admRoute.Route():      _admRoute,
	}
}

func NewAllRoutes() map[string]dpxlabd.RouteHandler {
	return RouteTable
}
