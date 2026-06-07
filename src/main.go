package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mcserver <command> [flags]")
		os.Exit(1)
	}
	command := os.Args[1]
	arguments := os.Args[2:]
	flags := map[string]string{
		"name":    "",
		"version": "latest",
	}

	for i, arg := range arguments {
		if arg == "--name" {
			fmt.Println("Creating server...")
			if !strings.HasPrefix(arguments[i+1], "--") {
				flags["name"] = arguments[i+1]
			}
		} else if arg == "--version" {
			if !strings.HasPrefix(arguments[i+1], "--") {
				flags["version"] = arguments[i+1]
			}
		}
	}

	switch command {
	case "create":
		if flags["name"] == "" {
			fmt.Println("--name needed to create server")
			os.Exit(1)
		}

		os.Mkdir(flags["name"], 0755)
	case "start":
		fmt.Println(command, arguments)
		os.Exit(0)
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
	os.Exit(0)
}
