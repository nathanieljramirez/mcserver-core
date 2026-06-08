package main

import (
	"fmt"
	"os"
	"os/exec"
)

func create(world string) {
	fmt.Println(world)
	fmt.Println("Creating server...")
	exec.Command("java", "-jar", "paper.jar")
	os.Mkdir(world, 0755)
}
