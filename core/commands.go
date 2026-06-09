package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func Install(version string) {
	if version == "latest" {
		var VersionResponse struct {
			Versions map[string][]string `json:"versions"`
		}

		Parse("https://fill.papermc.io/v3/projects/paper", &VersionResponse)
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

	Parse(PaperAPI, &PaperResponse)
	DownloadFile(PaperResponse.Downloads.ServerDefault.Name, PaperResponse.Downloads.ServerDefault.Url)
	CreateConfig(Config{Jar: PaperResponse.Downloads.ServerDefault.Name})
}

func Create(name string) error {
	fmt.Println("Creating server...")
	err := os.Mkdir(name, 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(name, "eula.txt"), []byte("eula=true"), 0644)
	if err != nil {
		return err
	}

	return nil
}
