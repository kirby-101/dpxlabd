package main

import (
	"os"

	_config "github.com/kirby-101/dpxlabd/config"
	logging "github.com/kirby-101/logging-go"
)

var (
	defaultConfigFilePaths = []string{"./config.yml", "/usr/local/etc/dpxlabd/config.yml"}
	rwForOwnerOnlyPerm     = 0o600
)

func initConfig() (*_config.Config, error) {
	var configFile *os.File

	for _, path := range defaultConfigFilePaths {

		// #nosec G304 -- Zugriff nur auf bekannte Log- und Config-Dateien
		file, err := os.OpenFile(
			path,
			os.O_RDONLY,
			os.FileMode(rwForOwnerOnlyPerm),
		)

		if err == nil {
			configFile = file
		}
	}

	return _config.LoadFile(configFile)
}

func initLogger(filepath, level string) (*logging.Logger, error) {
	var logger *logging.Logger

	if filepath == "" {
		logger = logging.NewLogger(os.Stdout)
	} else {
		// #nosec G304 -- Zugriff nur auf bekannte Log- und Config-Dateien
		logFile, err := os.OpenFile(
			filepath,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			os.FileMode(rwForOwnerOnlyPerm),
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
