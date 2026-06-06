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
	switch os.Args[1] {
	case "create":
		var name string = ""
		for i, arg := range os.Args[2:] {
			if arg == "--name" {
				fmt.Println("Creating server...")
				if strings.HasPrefix(os.Args[i+1], "--") {
					name = os.Args[i+1]
				} else {
					fmt.Println("--name requires a value")
					os.Exit(1)
				}
				return
			} else {
				fmt.Println("--name needed to create server")
				os.Exit(1)
			}
		}
	case "start":
		fmt.Println(os.Args[1], os.Args[2:])
		os.Exit(0)
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
	os.Exit(0)
}
