package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
)

func install(version string) {
	resp, err := http.Get("https://api.papermc.io/v2/projects/paper")
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
		fmt.Println(paper.Versions[len(paper.Versions)-1])
	} else if slices.Contains(paper.Versions, version) {
		fmt.Println(version)
	} else {
		fmt.Println("Available versions:", paper.Versions)
	}
}
