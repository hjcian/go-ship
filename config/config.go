package config

import (
	"go-ship/register"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Images []ImageConfig `yaml:"images"`
}

type ImageConfig struct {
	Name        string                `yaml:"name"`
	TagPatterns []string              `yaml:"tag_patterns"`
	Registry    register.RegistryType `yaml:"registry"`
}

func Load(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
