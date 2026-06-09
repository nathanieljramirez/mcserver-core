package core

import (
	"encoding/json"
	"os"
)

type Config struct {
	Jar string `jar:"version"`
}

func CreateConfig(config Config) error {
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("config.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadConfig() (Config, error) {
	var config Config
	data, err := os.ReadFile("config.json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
