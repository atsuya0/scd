package cmd

import (
	"bytes"
	"encoding/json"
	"os"
)

func getSecond() (second, error) {
	path, err := getDataPath()
	if err != nil {
		return second{}, err
	}

	var dataFile *os.File
	if _, err := os.Stat(path); os.IsNotExist(err) {
		dataFile, err = os.Create(path)
		return second{dataFile: dataFile, roots: make([]Root, 0)}, err
	} else if err != nil {
		return second{}, err
	} else {
		dataFile, err = os.OpenFile(path, os.O_RDWR, 0600)
		if err != nil {
			return second{}, err
		}
	}

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(dataFile); err != nil {
		return second{}, err
	}
	var roots []Root
	if err = json.Unmarshal(buffer.Bytes(), &roots); err != nil {
		return second{}, err
	}

	return second{dataFile: dataFile, roots: roots}, nil
}
