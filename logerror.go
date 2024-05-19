package main

import "log/slog"

func LogError(err error, message string) {
	if err != nil {
		slog.Error(message)
	}
}
