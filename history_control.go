package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func CheckIncludeHistory(id string) bool {
	CheckAndCreateHistory()
	historyFile, err := os.ReadFile("history.txt")
	LogError(err, fmt.Sprintf("An error occurred while reviewing the history file. %v", err))
	history := string(historyFile)
	return strings.Contains(history, id)
}

func AddToHistory(id string) {
	CheckAndCreateHistory()
	historyFile, err := os.OpenFile("history.txt", os.O_WRONLY|os.O_APPEND, 0666)
	LogError(err, fmt.Sprintf("An error occurred while reviewing the history file. %v", err))
	defer historyFile.Close()

	id = id + "\n"

	data := []byte(id)
	_, err = historyFile.Write(data)
	LogError(err, fmt.Sprintf("An error occurred while writing the history file. %v", err))
}

func CheckAndCreateHistory() {
	_, err := os.Stat("history.txt")
	if err != nil {
		slog.Info(fmt.Sprintf("The history file could not be found. Create a new one. %v", err))
		f, err := os.Create("history.txt")
		LogError(err, fmt.Sprintf("Failed to create history file. %v", err))
		f.Close()
	}
}
