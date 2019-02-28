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

type Source struct {
	Pairs []Pair `json:"source"`
}

func (s *Source) match(name string) (int, string, error) {
	for i, pair := range s.Pairs {
		if pair.Name == name {
			return i, pair.Path, nil
		}
	}
	err := fmt.Errorf("%s is invalid name.", name)

	return 0, "", err
}

func (s *Source) isDuplicate(options RegisterOptions) (err error) {
	for _, pair := range s.Pairs {
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

func (s *Source) del(i int) error {
	if 0 <= i && i < len(s.Pairs) {
		s.Pairs = append(s.Pairs[:i:i], s.Pairs[i+1:]...)
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

func newSourceFile() error {
	src, err := getListPath()
	if err != nil {
		return err
	}
	file, err := os.Create(src)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	jsonBytes, err := json.Marshal(Source{Pairs: []Pair{}})
	if err != nil {
		return err
	}

	_, err = file.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}

func loadSource(flag int) (*os.File, Source, error) {
	path, err := getListPath()
	if err != nil {
		return &os.File{}, Source{}, err
	}

	if _, err := os.Stat(path); err != nil {
		if err := newSourceFile(); err != nil {
			return &os.File{}, Source{}, err
		}
	}

	file, err := os.OpenFile(path, flag, 0600)
	if err != nil {
		return &os.File{}, Source{}, err
	}

	decoder := json.NewDecoder(file)
	source := Source{}
	if err = decoder.Decode(&source); err != nil {
		return &os.File{}, Source{}, err
	}

	return file, source, nil
}
