package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	WWW struct {
		Address string `yaml:"addr"`
	} `yaml:"www"`

	Logging struct {
		File  string `yaml:"file"`
		Level string `yaml:"level"`
	} `yaml:"logging"`
}

func LoadDefault() *Config {
	var cfg = &Config{}

	cfg.WWW.Address = "10.0.0.1:8080"
	cfg.Logging.File = ""
	cfg.Logging.Level = ""

	return cfg
}

func Load(file *os.File) (*Config, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err1 := yaml.Unmarshal(data, &cfg)
	if err1 != nil {
		return nil, err1
	}

	return &cfg, nil
}
