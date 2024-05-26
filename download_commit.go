package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func DownloadCommit(config Config, repo string, sha string) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/zipball/%s", config.Username, repo, sha)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Token))
	client := new(http.Client)
	resp, err := client.Do(request)
	LogError(err, fmt.Sprintf("Error while downloading commit zipfile from GitHub API. %v", err))
	slog.Info(fmt.Sprintf("GitHub download | Status Code: %v | %v", resp.StatusCode, resp.Request.URL))
	if resp.StatusCode != http.StatusOK {
		slog.Error(fmt.Sprintf("A status code other than http.StatusOK is returned from the GitHub API. %v", resp.StatusCode))
	}

	file, err := os.Create(fmt.Sprintf("commits/%s-%s", repo, sha))
	LogError(err, fmt.Sprintf("An error occurred while creating the file. %v", err))

	_, err = io.Copy(file, resp.Body)
	LogError(err, fmt.Sprintf("An error occurred while downloading the file. %v", err))

	resp.Body.Close()
	file.Close()
	time.Sleep(1 * time.Second)
}
