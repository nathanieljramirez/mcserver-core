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
		"--name": "",
	}

	for i, arg := range arguments {
		if _, ok := flags[arg]; ok {
			if !strings.HasPrefix(arguments[i+1], "--") {
				flags[arg] = arguments[i+1]
			}
		}
	}

	switch command {
	case "install":
		if len(arguments) == 0 || strings.HasPrefix(arguments[0], "--") {
			fmt.Println("Install command needs a version")
			os.Exit(1)
		}

		install(arguments[0])
	case "create":
		if flags["--name"] == "" {
			fmt.Println("--name needed to create server")
			os.Exit(1)
		}

		create(flags["--name"])

		fmt.Println("Creating server...")
		os.Mkdir(flags["--name"], 0755)
	case "start":
		if flags["--name"] == "" {
			fmt.Println("--name needed to start server")
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
	os.Exit(0)
}
