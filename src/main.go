package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mcserver <command> [flags]")
		os.Exit(1)
	}
	Command(os.Args[1], os.Args[2:])
	os.Exit(0)
}

func Command(command string, args []string) {
	switch command {
	case "create":
		for _, arg := range args {
			if arg == "--name" {
				fmt.Println("Creating server...")
				return
			} else {
				fmt.Println("--name needed to create server")
			}
		}
	case "start":
		fmt.Println(command, args)
		os.Exit(0)
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}
