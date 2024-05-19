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
	for _, v := range commitList[1] {
		fmt.Println(v.Commit)
		fmt.Println(v.Commit.Message)
		fmt.Println(v.SHA)
		fmt.Println()
	}
}
