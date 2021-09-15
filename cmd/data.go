package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func getEnvPath() (string, error) {
	path := os.Getenv("SCD_LIST_PATH")
	if path == "" {
		return "", errors.New("Cannot get path what use defined env")
	}
	return path, nil
}

func getXdgPath() (string, error) {
	conf := os.Getenv("XDG_CONFIG_HOME")
	if conf == "" {
		return "", errors.New("Cannot get path what use XDG_CONFIG_HOME")
	}
	path := filepath.Join(conf, "scd")
	if err := os.MkdirAll(path, 0700); err != nil {
		return "", err
	}

	return filepath.Join(path, "list.json"), nil
}

func getConfPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(homeDir, ".config", "scd")
	if err := os.MkdirAll(path, 0700); err != nil {
		return "", err
	}

	return filepath.Join(path, "list.json"), nil
}

func getSecondPath() (string, error) {
	if path, err := getEnvPath(); err == nil {
		return path, nil
	}
	if path, err := getXdgPath(); err == nil {
		return path, nil
	}
	if path, err := getConfPath(); err != nil {
		return path, nil
	} else {
		return "", fmt.Errorf("Cannot get the file path to save the data: %w", err)
	}
}
