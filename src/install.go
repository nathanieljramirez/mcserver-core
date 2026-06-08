package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
)

func install(version string) {
	url := "https://api.papermc.io/v2/projects/paper"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	var paper struct {
		Versions []string `json:"versions"`
	}

	err = json.NewDecoder(resp.Body).Decode(&paper)

	if err != nil {
		fmt.Println(err)
		return
	}

	if version == "latest" {
		version = paper.Versions[len(paper.Versions)-1]
	} else if !slices.Contains(paper.Versions, version) {
		fmt.Println("Available versions:", paper.Versions)
		os.Exit(1)
	}
}
