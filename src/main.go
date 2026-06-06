package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	Command(os.Args[1], os.Args[2:])
	os.Exit(0)
}

func Command(command string, args []string) {
	fmt.Println(command, args)
}
