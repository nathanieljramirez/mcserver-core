package main

import (
	"fmt"
	"os"
)

func create(name string) {
	fmt.Println("Creating server...")
	os.Mkdir(name, 0755)

	err := os.WriteFile(name+"/eula.txt", []byte("eula=true"), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
