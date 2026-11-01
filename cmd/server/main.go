//go:build freebsd

package main

import (
	"fmt"

	"github.com/kirby-101/dpxlabd"
	_config "github.com/kirby-101/dpxlabd/config"
	logging "github.com/kirby-101/logging-go"
)

var (
	// $(git describe --tags --always --dirty)
	Version string

	config *_config.Config
	logger *logging.Logger
)

func init() {
	var err error

	config, err = initConfig()
	if err != nil {
		panic(
			fmt.Errorf("load config file: %s", err.Error()),
		)
	}

	logger, err = initLogger(config.Logging.File, config.Logging.Level)
	if err != nil {
		panic(err)
	}
}

func main() {
	logger.Info("Version", Version)

	srv := dpxlabd.New(config, logger)

	registerAllSubroutes(srv)
	srv.Start()
}
