package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type CommitsWithProjectName struct {
	Commits     []CommitsResp
	ProjectName string
}

type CommitsResp struct {
	SHA    string `json:"sha"`
	Commit struct {
		Message string `json:"message"`
		Author  struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commit"`
}

func GetCommitList(config Config, repos []string) []CommitsWithProjectName {
	var commitsList []CommitsWithProjectName
	for _, v := range repos {
		time.Sleep(1 * time.Second)
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", config.Username, v)
		request, _ := http.NewRequest("GET", url, nil)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Token))
		request.Header.Add("Accept", "application/vnd.github.v3+json")
		client := new(http.Client)
		resp, err := client.Do(request)
		LogError(err, fmt.Sprintf("Error while retrieving commit history from GitHub API. %v", err))
		defer resp.Body.Close()
		slog.Info(fmt.Sprintf("GitHub commits | Status Code: %v | %v", resp.StatusCode, resp.Request.URL))
		if resp.StatusCode != http.StatusOK {
			slog.Error(fmt.Sprintf("A status code other than http.StatusOK is returned from the GitHub API. %v", resp.StatusCode))
		}

		body, err := io.ReadAll(resp.Body)
		LogError(err, fmt.Sprintf("An error occurred while reading the Body of the response. %v", err))

		var result []CommitsResp
		err = json.Unmarshal(body, &result)
		LogError(err, fmt.Sprintf("An error occurred while converting Body to Json. %v", err))

		var cwpn CommitsWithProjectName
		cwpn.Commits = result
		cwpn.ProjectName = v
		commitsList = append(commitsList, cwpn)
	}
	return commitsList
}
