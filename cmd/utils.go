package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func getEnvPath() (string, error) {
	path := os.Getenv("SECOND_LIST_PATH")
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
	path := filepath.Join(conf, "second")
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
	path := filepath.Join(homeDir, ".config", "second")
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

func getSecond() (second, error) {
	path, err := getSecondPath()
	if err != nil {
		return second{}, err
	}

	var file *os.File
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err = os.Create(path)
		return second{file: file, pairs: make([]Pair, 0)}, err
	} else if err != nil {
		return second{}, err
	} else {
		file, err = os.OpenFile(path, os.O_RDWR, 0600)
		if err != nil {
			return second{}, err
		}
	}

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		return second{}, err
	}
	var pairs []Pair
	if err = json.Unmarshal(buffer.Bytes(), &pairs); err != nil {
		return second{}, err
	}

	return second{file: file, pairs: pairs}, nil
}

func getPairs() ([]Pair, error) {
	path, err := getSecondPath()
	if err != nil {
		return make([]Pair, 0), err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return make([]Pair, 0), nil
	} else if err != nil {
		return make([]Pair, 0), err
	}
	file, err := os.Open(path)
	if err != nil {
		return make([]Pair, 0), err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		return make([]Pair, 0), err
	}
	var pairs []Pair
	if err = json.Unmarshal(buffer.Bytes(), &pairs); err != nil {
		return make([]Pair, 0), err
	}

	return pairs, nil
}
