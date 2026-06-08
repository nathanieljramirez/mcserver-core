package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func parse(url string, target any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("%s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func install(version string) {
	PaperAPI := "https://fill.papermc.io/v3/projects/paper/versions/" + version + "/builds/latest"

	var PaperResponse struct {
		Downloads struct {
			ServerDefault struct {
				Url string `json:"url"`
			} `json:"server:default"`
		} `json:"downloads"`
	}

	err := parse(PaperAPI, &PaperResponse)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println(PaperResponse.Downloads.ServerDefault.Url)
}
