package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (s *Source) del(i int) {
	s.Pairs = append(s.Pairs[:i:i], s.Pairs[i+1:]...)
}

func getSrc() (string, error) {
	path := os.Getenv("SECOND_CMD_PATH")

	if path != "" {
		return path, nil
	} else {
		user, err := user.Current()
		if err != nil {
			return "", err
		}
		return filepath.Join(user.HomeDir, ".second"), nil
	}
}

func newSourceFile() error {
	src, err := getSrc()
	if err != nil {
		return err
	}
	file, err := os.Create(src)
	if err != nil {
		return err
	}
	defer file.Close()

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
	tmpFile, err := ioutil.TempFile(os.TempDir(), "second")
	if err != nil {
		return tmpFile, Source{}, err
	}

	src, err := getSrc()
	if err != nil {
		return tmpFile, Source{}, err
	}

	if _, err := os.Stat(src); err != nil {
		if err := newSourceFile(); err != nil {
			return tmpFile, Source{}, err
		}
	}

	file, err := os.OpenFile(src, flag, 0600)
	if err != nil {
		return tmpFile, Source{}, err
	}

	decoder := json.NewDecoder(file)
	source := Source{}
	if err = decoder.Decode(&source); err != nil {
		return tmpFile, Source{}, err
	}

	return file, source, nil
}