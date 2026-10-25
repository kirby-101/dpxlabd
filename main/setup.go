package main

import (
	"os"
	"sync"

	"github.com/kirby-101/dpxlab/config"
	"github.com/kirby-101/dpxlab/db"
	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/adm"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/assets"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/md"
	"github.com/kirby-101/logging-go"
)

func setupRoutes() map[string]dpxlabd.RouteHandler {
	var _routes = map[string]dpxlabd.RouteHandler{
		"/": routes.NewIndexRoute(),
	}

	_routes["/adm"] = adm.NewAdmRoute(_routes)
	_routes["/assets"] = assets.NewAssetsRoute(_routes)
	_routes["/pub"] = md.NewMarkdownRoute(_routes)

	return _routes
}

/*
package dpxlabd

import (
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/adm"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/assets"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/md"
)

var (
	_routes map[string]RouteHandler

	// index,robots.txt,md,assets,adm:safe,proxy,board,auth
	_indexRoute    = routes.NewIndexRoute()
	_9309Route     = routes.NewRobotsTxtRoute()
	_markdownRoute = md.NewMarkdownRoute()
	_assetsRoute   = assets.NewAssetsRoute()
	_admRoute      = adm.NewAdmRoute()
)

func NewAllRoutes() map[string]RouteHandler {

}

*/

func setupLogger(filepath, level string) (*logging.Logger, error) {
	var logger *logging.Logger

	if filepath == "" {
		logger = logging.NewLogger(os.Stdout)
	} else {
		// #nosec G304 -- Zugriff nur auf bekannte Log- und Config-Dateien
		logFile, err := os.OpenFile(
			filepath,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			rwForOwnerOnlyPerm,
		)
		if err != nil {
			return nil, err
		}

		logger = logging.NewLogger(logFile)
	}

	switch level {
	case "debug":
		logger.Level = logging.LogDebug

	case "info":
		logger.Level = logging.LogInfo

	case "error":
		logger.Level = logging.LogError

	case "fatal":
		logger.Level = logging.LogFatal

	default:
		logger.Level = logging.Level(1)
	}

	return logger, nil
}

// !
func setupDatabase() (*db.Database, error) {
	return &db.Database{
		Mutex: sync.Mutex{},
	}, nil
}

func setupConfig(filepath string) (*config.Config, error) {
	// #nosec G304 -- Zugriff nur auf bekannte Log- und Config-Dateien
	file, err := os.OpenFile(
		filepath,
		os.O_RDONLY,
		rwForOwnerOnlyPerm,
	)
	if err != nil {
		return nil, err
	}

	cfg, err := config.Load(file)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
