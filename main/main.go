package main

import (
	"fmt"

	"github.com/kirby-101/dpxlab/config"
	"github.com/kirby-101/dpxlab/db"
	"github.com/kirby-101/dpxlab/dpxlabd"
	"github.com/kirby-101/dpxlab/dpxlabd/routes"
	"github.com/kirby-101/logging-go"
)

// $(git describe --tags --always --dirty)
var Version string

const (
	defaultConfigFilePath = "/usr/local/etc/dpxlabd.yml"
	rwForOwnerOnlyPerm    = 0o600
)

var (
	cfg      *config.Config
	database *db.Database
	logger   *logging.Logger
)

func init() {
	_cfg, err := setupConfig(defaultConfigFilePath)
	if err != nil {
		panic(err)
	}

	_db, err := setupDatabase()
	if err != nil {
		panic(err)
	}

	_logger, err := setupLogger(_cfg.Logging.File, _cfg.Logging.Level)
	if err != nil {
		panic(err)
	}

	cfg = _cfg
	database = _db
	logger = _logger
}

func main() {
	logger.Info("Version", Version)

	orch := dpxlabd.New(
		cfg,
		database,
		logger,
		setupRoutes(),
		routes.NewErrorRouteHandler(),
	)

	if err := orch.Start(); err != nil {
		logger.Fatal(fmt.Sprintf("%#v", err))
	}
}
