package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	WWW struct {
		Address string `yaml:"addr"`
	} `yaml:"www"`
}

func Load(file *os.File) (*Config, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("io ReadAll: %s", err.Error())
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
