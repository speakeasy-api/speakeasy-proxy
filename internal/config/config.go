package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v9"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DownstreamBaseURL string   `yaml:"downstreamBaseURL" env:"DOWNSTREAM_BASE_URL" validate:"required"`
	Port              string   `yaml:"port" env:"PORT" validate:"required" default:"3333" `
	APIKey            string   `env:"SPEAKEASY_API_KEY" validate:"required"`
	ApiID             string   `yaml:"apiID" env:"SPEAKEASY_API_ID" validate:"required"`
	VersionID         string   `yaml:"versionID" env:"SPEAKEASY_VERSION_ID" validate:"required"`
	OpenAPIDocs       []string `yaml:"openAPIDocs" env:"OPENAPI_DOCS" validate:"min=1" default:"[\"./openapi.yaml\"]"`
	ConfigLocation    string   `yaml:"configLocation" env:"CONFIG_LOCATION" validate:"required" default:"./config.yaml"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := defaults.Set(&cfg); err != nil {
		return nil, fmt.Errorf("failed to set defaults: %w", err)
	}

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse env: %w", err)
	}

	if _, err := os.Stat(cfg.ConfigLocation); err == nil {
		configData, err := os.ReadFile(cfg.ConfigLocation)
		if err != nil {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}

		if err := yaml.Unmarshal(configData, &cfg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
		}
	}

	v := validator.New()

	// TODO maybe parse errors and try to give better details on resolving requirements
	if err := v.Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return &cfg, nil
}
