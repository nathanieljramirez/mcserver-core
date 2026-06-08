package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
)

func parse(url string, target any) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func install(version string) {
	url := "https://api.papermc.io/v2/projects/paper"

	var PaperResponse struct {
		Versions []string `json:"versions"`
	}

	parse(url, &PaperResponse)

	if version == "latest" {
		version = PaperResponse.Versions[len(PaperResponse.Versions)-1]
	} else if !slices.Contains(PaperResponse.Versions, version) {
		fmt.Println("Available versions:", PaperResponse.Versions)
		os.Exit(1)
	}
}
