package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func parse(url string, target any) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error requesting URL:", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
	}

	json.NewDecoder(resp.Body).Decode(target)
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

func Install(version string) {
	if version == "latest" {

		var VersionResponse struct {
			Versions map[string][]string `json:"versions"`
		}

		parse("https://fill.papermc.io/v3/projects/paper", &VersionResponse)

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

	parse(PaperAPI, &PaperResponse)

	downloadFile(PaperResponse.Downloads.ServerDefault.Name, PaperResponse.Downloads.ServerDefault.Url)

	CreateConfig(Config{Jar: PaperResponse.Downloads.ServerDefault.Name})
}
