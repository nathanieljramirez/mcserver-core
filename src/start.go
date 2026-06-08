package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func start(name string, memory string) {
	config := readConfig()
	path := filepath.Join("..", config.Jar)

	cmd := exec.Command("java", "-jar", path)
	if memory != "" {
		cmd = exec.Command("java", "-Xms"+memory+"G", "-Xmx"+memory+"G", "-jar", "--nogui")
	}
	cmd.Dir = name
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
