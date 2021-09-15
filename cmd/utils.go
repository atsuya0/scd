package cmd

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func getSecond() (second, error) {
	path, err := getDataPath()
	if err != nil {
		return second{}, err
	}

	var file *os.File
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err = os.Create(path)
		return second{file: file, roots: make([]Root, 0)}, err
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
	var roots []Root
	if err = json.Unmarshal(buffer.Bytes(), &roots); err != nil {
		return second{}, err
	}

	return second{file: file, roots: roots}, nil
}

func getRoots() ([]Root, error) {
	path, err := getDataPath()
	if err != nil {
		return make([]Root, 0), err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return make([]Root, 0), nil
	} else if err != nil {
		return make([]Root, 0), err
	}
	file, err := os.Open(path)
	if err != nil {
		return make([]Root, 0), err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		return make([]Root, 0), err
	}
	var roots []Root
	if err = json.Unmarshal(buffer.Bytes(), &roots); err != nil {
		return make([]Root, 0), err
	}

	return roots, nil
}
