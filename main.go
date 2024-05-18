package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Token    string `yaml:"token"`
	Username string `yaml:"username"`
}

type ReposResp struct {
	Name string `json:"name"`
}

func main() {
	data, err := os.ReadFile("config.yaml")
	logError(err, fmt.Sprintf("An error occurred while loading the file. %v", err))

	var config Config
	err = yaml.Unmarshal(data, &config)
	logError(err, fmt.Sprintf("An error occurred during deserializaton of the yaml file. %v", err))

	fmt.Println(getRepoList(config))
}

func getRepoList(config Config) []string {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", config.Username)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Token))
	client := new(http.Client)
	resp, err := client.Do(request)
	logError(err, fmt.Sprintf("An error occurred while retrieving a list of repositories via GitHub's API. %v", err))
	defer resp.Body.Close()
	slog.Info(fmt.Sprintf("GitHub repos | Status code: %v", resp.StatusCode))
	if resp.StatusCode != http.StatusOK {
		slog.Error(fmt.Sprintf("A status code other than http.StatusOK is returned from the GitHub API. %v", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	logError(err, fmt.Sprintf("An error occurred while reading the Body of the response. %v", err))

	var result []ReposResp
	err = json.Unmarshal(body, &result)
	logError(err, fmt.Sprintf("An error occurred while converting Body to Json. %v", err))

	var repoList []string
	for _, item := range result {
		repoList = append(repoList, item.Name)
	}
	return repoList
}

func logError(err error, message string) {
	if err != nil {
		slog.Error(message)
	}
}
