package core

import (
	"encoding/json"
	"fmt"
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

func readConfig() Config {
	var config Config
	data, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	return config
}
