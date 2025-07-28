package register

import "errors"

type RegistryType int

const (
	Unknown RegistryType = iota
	DockerHub
	AWS
)

func (r RegistryType) ToString() string {
	switch r {
	case DockerHub:
		return "dockerhub"
	case AWS:
		return "aws"
	default:
		return "unknown"
	}
}

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
