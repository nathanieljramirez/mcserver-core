package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Jar string `jar:"version"`
}

func createConfig(config Config) error {
	data, err := json.MarshalIndent(config, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile("config.json", data, 0644)
}
