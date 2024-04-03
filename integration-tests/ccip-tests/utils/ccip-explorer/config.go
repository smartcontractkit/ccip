package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Networks map[string]struct {
		URL    string            `toml:"url"`
		Tokens map[string]string `toml:"tokens"`
	} `toml:"networks"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}
