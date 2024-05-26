package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type ReposResp struct {
	Name string `json:"name"`
}

func GetRepoList(config Config) []string {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", config.Username)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Token))
	client := new(http.Client)
	resp, err := client.Do(request)
	LogError(err, fmt.Sprintf("An error occurred while retrieving a list of repositories via GitHub's API. %v", err))
	defer resp.Body.Close()
	slog.Info(fmt.Sprintf("GitHub repos | Status code: %v | %v", resp.StatusCode, resp.Request.URL))
	if resp.StatusCode != http.StatusOK {
		slog.Error(fmt.Sprintf("A status code other than http.StatusOK is returned from the GitHub API. %v", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	LogError(err, fmt.Sprintf("An error occurred while reading the Body of the response. %v", err))

	var result []ReposResp
	err = json.Unmarshal(body, &result)
	LogError(err, fmt.Sprintf("An error occurred while converting Body to Json. %v", err))

	var repoList []string
	for _, item := range result {
		repoList = append(repoList, item.Name)
	}
	return repoList
}
