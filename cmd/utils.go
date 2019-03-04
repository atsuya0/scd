package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type Pair struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type List struct {
	Pairs []Pair `json:"list"`
}

func (l *List) match(name string) (int, string, error) {
	for i, pair := range l.Pairs {
		if pair.Name == name {
			return i, pair.Path, nil
		}
	}
	err := fmt.Errorf("%s is invalid name.", name)

	return 0, "", err
}

func (l *List) isDuplicate(options RegisterOptions) (err error) {
	for _, pair := range l.Pairs {
		if pair.Name == options.name {
			err = fmt.Errorf("This name has already been registered.")
			return
		}
		if pair.Path == options.path {
			err = fmt.Errorf("This path has already been registered.")
			return
		}
	}
	return
}

func (l *List) del(i int) error {
	if 0 <= i && i < len(l.Pairs) {
		l.Pairs = append(l.Pairs[:i:i], l.Pairs[i+1:]...)
		return nil
	}
	return fmt.Errorf("out of range.")
}

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
		return "", err
	}
}

func formatFile() error {
	path, err := getListPath()
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("%+v\n", err)
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
		return List{}, &os.File{}, err
	}

	if _, err := os.Stat(path); err != nil {
		if err := formatFile(); err != nil {
			return List{}, &os.File{}, err
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
