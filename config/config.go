package config

import (
	"io"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server  OptionsServer  `yaml:"server"`
	WWW     OptionsWWW     `yaml:"www"`
	Sched   OptionsSched   `yaml:"sched"`
	Logging OptionsLogging `yaml:"logging"`
}

type OptionsServer struct {
	ShutdownTimeout time.Duration
}

type OptionsWWW struct {
	Enabled bool   `yaml:"enabled"`
	Address string `yaml:"address"`
}

type OptionsSched struct {
	Enabled bool           `yaml:"enabled"`
	Storage OptionsStorage `yaml:"storage"`
}

type OptionsStorage struct {
	EndpointURL     string `yaml:"endpoint"`          // https://%s.r2.cloudflarestorage.com
	AccountID       string `yaml:"account-id"`        // Provide your Cloudflare account ID#
	AccessKeyID     string `yaml:"access-key-id"`     // Retrieve your S3 API credentials for your R2 bucket via API tokens
	AccessKeySecret string `yaml:"access-key-secret"` // (see: https://developers.cloudflare.com/r2/api/tokens)
}

type OptionsLogging struct {
	File  string `yaml:"file"`
	Level string `yaml:"level"`
}

func LoadFile(file *os.File) (*Config, error) {
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
