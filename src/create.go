package main

import (
	"fmt"
	"os"
)

func create(world string) {
	fmt.Println(world)
	fmt.Println("Creating server...")
	os.Mkdir(world, 0755)
}
