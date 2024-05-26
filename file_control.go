package main

import (
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path"
	"path/filepath"
)

func CopyToMountedPath(config Config) {
	stat, err := os.Stat(config.MountedPath)
	LogError(err, fmt.Sprintf("An error occurred while checking the status of the mounted folder. %v", err))
	if !stat.IsDir() {
		slog.Error("Mount destination path was not a folder.")
	}
	savePath := path.Join(config.MountedPath, "git-to-smb")
	if _, err := os.Stat(savePath); err != nil {
		slog.Info("The destination folder could not be found. Create a new one. %v", err)
		os.MkdirAll(savePath, os.ModePerm)
	}

	err = filepath.Walk("commits", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if _, err := os.Stat(path); err == nil {
			if stat, _ := os.Stat(path); stat.IsDir() {
				return nil
			}

			src, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
			if err != nil {
				return err
			}

			dst, err := os.OpenFile(pathJoin(savePath, fmt.Sprintf("%s.zip", info.Name())), os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
			if err != nil {
				return err
			}

			_, err = io.Copy(dst, src)
			if err != nil {
				return err
			}

			slog.Info(fmt.Sprintf("Copied: %s", pathJoin(savePath, fmt.Sprintf("%s.zip", info.Name()))))
		}
		return nil
	})
	LogError(err, fmt.Sprintf("An error occurred while copying a file (filepath.Walk), %v", err))
}

func pathJoin(str1 string, str2 string) string {
	return path.Join(str1, str2)
}
