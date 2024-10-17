package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string
		Port int
	}
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	file, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
