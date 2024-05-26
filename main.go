package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Token       string `yaml:"token"`
	Username    string `yaml:"username"`
	MountedPath string `yaml:"mountedpath"`
}

func main() {
	CheckAndCreateLogJSON()
	logJson, err := os.OpenFile("log.json", os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(logJson, nil))
	slog.SetDefault(logger)

	data, err := os.ReadFile("config.yaml")
	LogError(err, fmt.Sprintf("An error occurred while loading the file. %v", err))

	var config Config
	err = yaml.Unmarshal(data, &config)
	LogError(err, fmt.Sprintf("An error occurred during deserializaton of the yaml file. %v", err))

	repoList := GetRepoList(config)
	time.Sleep(1 * time.Second)
	slog.Info(fmt.Sprintf("Repo List: %v", repoList))

	commitList := GetCommitList(config, repoList)
	time.Sleep(1 * time.Second)
	for _, v := range commitList {
		for _, c := range v.Commits {
			id := fmt.Sprintf("%s@%s", v.ProjectName, c.SHA)
			if !CheckIncludeHistory(id) {
				DownloadCommit(config, v.ProjectName, c.SHA)
				AddToHistory(id)
			}
		}
	}

	CopyToMountedPath(config)
}
