package config

import (
	"os"

	"github.com/NessibeliY/binance-ticker-parser/internal/values"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Symbols    []string `yaml:"symbols"`
	MaxWorkers int      `yaml:"max_workers"`
}

func Load() (*Config, error) {
	config := &Config{}
	rawYaml, err := os.ReadFile(values.ConfigFile)
	if err != nil {
		return nil, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYaml, &config)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshaling yaml")
	}
	return config, nil
}
