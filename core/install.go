package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

// https://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("%s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func install(version string) {
	if version == "latest" {

		var VersionResponse struct {
			Versions map[string][]string `json:"versions"`
		}

		err := parse("https://fill.papermc.io/v3/projects/paper", &VersionResponse)

		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		version = VersionResponse.Versions["26.1"][0]
	}

	PaperAPI := "https://fill.papermc.io/v3/projects/paper/versions/" + version + "/builds/latest"

	var PaperResponse struct {
		Downloads struct {
			ServerDefault struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"server:default"`
		} `json:"downloads"`
	}

	err := parse(PaperAPI, &PaperResponse)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	downloadFile(PaperResponse.Downloads.ServerDefault.Name, PaperResponse.Downloads.ServerDefault.Url)

	createConfig(Config{Jar: PaperResponse.Downloads.ServerDefault.Name})
}
