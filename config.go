package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

type RegistryType int

const (
	DockerHub RegistryType = iota
	AWS
)

func (r *RegistryType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	switch s {
	case "dockerhub":
		*r = DockerHub
	case "aws":
		*r = AWS
	default:
		return errors.New("invalid registry type")
	}

	return nil
}

type Config struct {
	Images []ImageConfig `yaml:"images"`
}

type ImageConfig struct {
	Name       string       `yaml:"name"`
	TagPattern string       `yaml:"tag_pattern"`
	Registry   RegistryType `yaml:"registry"`
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
