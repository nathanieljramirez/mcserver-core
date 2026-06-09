package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Parse(url string, target any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("%s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(target)

	if err != nil {
		return err
	}

	return nil
}

// https://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
func DownloadFile(filepath string, url string) error {
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
