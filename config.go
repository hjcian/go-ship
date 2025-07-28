package main

import (
	tagfetcher "go-ship/tag_fetcher"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Images []ImageConfig `yaml:"images"`
}

type ImageConfig struct {
	Name       string                  `yaml:"name"`
	TagPattern string                  `yaml:"tag_pattern"`
	Registry   tagfetcher.RegistryType `yaml:"registry"`
}

func LoadConfig(filePath string) (*Config, error) {
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
