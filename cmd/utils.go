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

func getListPath() (string, error) {
	path := os.Getenv("SECOND_LIST_PATH")

	if path != "" {
		return path, nil
	} else {
		user, err := user.Current()
		if err != nil {
			return "", err
		}
		return filepath.Join(user.HomeDir, ".second_list"), nil
	}
}

func newListFile() error {
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
			log.Fatalln(err)
		}
	}()

	jsonBytes, err := json.Marshal(List{Pairs: []Pair{}})
	if err != nil {
		return err
	}

	_, err = file.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}

func loadList(flag int) (*os.File, List, error) {
	path, err := getListPath()
	if err != nil {
		return &os.File{}, List{}, err
	}

	if _, err := os.Stat(path); err != nil {
		if err := newListFile(); err != nil {
			return &os.File{}, List{}, err
		}
	}

	file, err := os.OpenFile(path, flag, 0600)
	if err != nil {
		return &os.File{}, List{}, err
	}

	decoder := json.NewDecoder(file)
	var list List
	if err = decoder.Decode(&list); err != nil {
		return &os.File{}, List{}, err
	}

	return file, list, nil
}
