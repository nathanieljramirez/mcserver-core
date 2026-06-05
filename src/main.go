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
	Command(os.Args[1], os.Args[2])
	os.Exit(0)
}

func Command(arg1 string, arg2 string) {
	fmt.Println(arg1, arg2)
}
