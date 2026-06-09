package core

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
