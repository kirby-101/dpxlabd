package main

import (
	"os"
	"sync"

	"github.com/kirby-101/dpxlab/config"
	"github.com/kirby-101/dpxlab/db"
	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/assets"
	"github.com/kirby-101/dpxlab/dpxlabd/routes/md"
	"github.com/kirby-101/logging-go"
)

func setupRoutes() map[string]dpxlabd.RouteHandler {
	var routes = map[string]dpxlabd.RouteHandler{
		"/":   routes.NewIndexRoute(),
		"/md": md.NewMarkdownRoute(),
	}

	routes["/assets"] = assets.NewAssetsRoute(routes)

	return routes
}

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
