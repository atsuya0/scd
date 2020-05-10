package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func getEnvPath() (string, error) {
	path := os.Getenv("SECOND_LIST_PATH")
	if path == "" {
		return "", fmt.Errorf("Cannot get path what use defined env")
	}
	return path, nil
}

func getXdgPath() (string, error) {
	conf := os.Getenv("XDG_CONFIG_HOME")
	if conf == "" {
		return "", fmt.Errorf("Cannot get path what use XDG_CONFIG_HOME")
	}
	path := filepath.Join(conf, "second")
	if err := os.MkdirAll(path, 0700); err != nil {
		return "", err
	}

	return filepath.Join(path, "list.json"), nil
}

func getConfPath() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	path := filepath.Join(user.HomeDir, ".config", "second")
	if err := os.MkdirAll(path, 0700); err != nil {
		return "", err
	}

	return filepath.Join(path, "list.json"), nil
}

func getListPath() (string, error) {
	if path, err := getEnvPath(); err == nil {
		return path, nil
	}
	if path, err := getXdgPath(); err == nil {
		return path, nil
	}
	if path, err := getConfPath(); err == nil {
		return path, nil
	} else {
		return "", fmt.
			Errorf("Cannot get the path from the user infomation: %w", err)
	}
}

func formatFile() error {
	path, err := getListPath()
	if err != nil {
		return fmt.Errorf("Cannot get path of the list: %w", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	jsonBytes, err := json.Marshal(List{})
	if err != nil {
		return err
	}

	_, err = file.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}

func getListAndFile(flag int) (List, *os.File, error) {
	path, err := getListPath()
	if err != nil {
		return List{}, &os.File{}, fmt.Errorf("Cannot get path of the list: %w", err)
	}

	if _, err := os.Stat(path); err != nil {
		if err := formatFile(); err != nil {
			return List{}, &os.File{}, fmt.Errorf("Cannot format data file: %w", err)
		}
	}

	file, err := os.OpenFile(path, flag, 0600)
	if err != nil {
		return List{}, &os.File{}, err
	}

	decoder := json.NewDecoder(file)
	var list List
	if err = decoder.Decode(&list); err != nil {
		return List{}, &os.File{}, err
	}

	return list, file, nil
}
