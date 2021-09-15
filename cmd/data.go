package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func getDataDir() (string, error) {
	errMsg := "Failed to get the data directory"

	if path := os.Getenv("SCD_DATA_PATH"); path != "" {
		return path, nil
	}
	if dataHome := os.Getenv("XDG_DATA_HOME"); dataHome != "" {
		path := filepath.Join(dataHome, "scd")
		if err := os.MkdirAll(path, 0700); err != nil {
			return "", fmt.Errorf(errMsg+": %w", err)
		}
		return path, nil
	}
	if homeDir, err := os.UserHomeDir(); err != nil {
		path := filepath.Join(homeDir, ".local", "share", "scd")
		if err := os.MkdirAll(path, 0700); err != nil {
			return "", fmt.Errorf(errMsg+": %w", err)
		}
		return path, nil
	}
	return "", errors.New(errMsg)
}

func getDataPath() (string, error) {
	path, err := getDataDir()
	if err != nil {
		return "", errors.New("Failed to get the data path")
	}
	return filepath.Join(path, "list.json"), nil
}
