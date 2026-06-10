package mcserverkit

import (
	"os"
	"os/exec"
	"path/filepath"

	"mcserverkit.github.io/internal"
)

func Install(version string) error {
	if version == "latest" {
		var VersionResponse struct {
			Versions map[string][]string `json:"versions"`
		}
		err := internal.Parse("https://fill.papermc.io/v3/projects/paper", &VersionResponse)
		if err != nil {
			return err
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

	err := internal.Parse(PaperAPI, &PaperResponse)
	if err != nil {
		return err
	}

	err = internal.DownloadFile(PaperResponse.Downloads.ServerDefault.Name, PaperResponse.Downloads.ServerDefault.Url)
	if err != nil {
		return err
	}

	err = internal.CreateConfig(internal.Config{Jar: PaperResponse.Downloads.ServerDefault.Name})
	if err != nil {
		return err
	}
	return nil
}

func Create(name string) error {
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

func Start(name string, memory ...string) error {
	config, err := internal.ReadConfig()
	if err != nil {
		return err
	}

	path := filepath.Join("..", config.Jar)
	cmd := exec.Command("java", "-jar", path, "--nogui")
	if len(memory) > 0 {
		cmd = exec.Command("java", "-Xms"+memory[0], "-Xmx"+memory[0], "-jar", path, "--nogui")
	}
	cmd.Dir = name
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
