package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const configFile = "configs/config.yaml"

type Config struct {
	Symbols    []string `yaml:"symbols"`
	MaxWorkers int      `yaml:"max_workers"`
}

func Load() (*Config, error) {
	config := &Config{}
	rawYaml, err := os.ReadFile(configFile)
	if err != nil {
		return nil, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYaml, &config)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshaling yaml")
	}
	return config, nil
}
