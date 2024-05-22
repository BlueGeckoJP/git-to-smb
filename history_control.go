package main

import (
	"fmt"
	"os"
	"strings"
)

func CheckIncludeHistory(id string) bool {
	historyFile, err := os.ReadFile("history.txt")
	LogError(err, fmt.Sprintf("An error occurred while reviewing the history file. %v", err))
	history := string(historyFile)
	return strings.Contains(history, id)
}

func AddToHistory(id string) {
	historyFile, err := os.OpenFile("history.txt", os.O_WRONLY|os.O_APPEND, 0666)
	LogError(err, fmt.Sprintf("An error occurred while reviewing the history file. %v", err))
	defer historyFile.Close()

	id = id + "\n"

	data := []byte(id)
	_, err = historyFile.Write(data)
	LogError(err, fmt.Sprintf("An error occurred while writing the history file. %v", err))
}
