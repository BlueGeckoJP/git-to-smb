package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Token    string `yaml:"token"`
	Username string `yaml:"username"`
}

func main() {
	data, err := os.ReadFile("config.yaml")
	LogError(err, fmt.Sprintf("An error occurred while loading the file. %v", err))

	var config Config
	err = yaml.Unmarshal(data, &config)
	LogError(err, fmt.Sprintf("An error occurred during deserializaton of the yaml file. %v", err))

	repoList := GetRepoList(config)
	fmt.Println(repoList)

	commitList := GetCommitList(config, repoList)
	for _, v := range commitList {
		for _, c := range v.Commits {
			DownloadCommit(config, v.ProjectName, c.SHA)
			AddToHistory(fmt.Sprintf("%s@%s", v.ProjectName, c.SHA))
		}
	}
}
