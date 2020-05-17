package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/tayusa/go-chooser"
)

type second struct {
	file  *os.File
	pairs []Pair
}

type Pair struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (s *second) match(name string) (int, string, error) {
	for i, pair := range s.pairs {
		if pair.Name == name {
			return i, pair.Path, nil
		}
	}
	err := errors.New(name + " invalid name")

	return 0, "", err
}

func (s *second) isDuplicate(options RegisterOptions) (err error) {
	for _, pair := range s.pairs {
		if pair.Name == options.name {
			err = errors.New("This name has already been registered.")
			return
		}
		if pair.Path == options.path {
			err = errors.New("This path has already been registered.")
			return
		}
	}
	return
}

func (s *second) del(i int) error {
	if 0 <= i && i < len(s.pairs) {
		s.pairs = append(s.pairs[:i:i], s.pairs[i+1:]...)
		return nil
	}
	return errors.New("out of range.")
}

func (s *second) update() error {
	json, err := json.MarshalIndent(s.pairs, "", strings.Repeat(" ", 2))
	if err != nil {
		return err
	}
	if err := s.file.Truncate(0); err != nil {
		return err
	}
	if _, err = s.file.WriteAt(json, 0); err != nil {
		return err
	}

	return nil
}

func (s *second) choose() (string, error) {
	var names []string
	for _, pair := range s.pairs {
		names = append(names, pair.Name)
	}
	nameChooser, err := chooser.NewChooser(names)
	if err != nil {
		return "", err
	}
	if names := nameChooser.Run(); len(names) != 0 {
		return names[0], nil
	}
	return "", nil
}

// Get an undeclared name error.
// second := second {pairs: pairs}
func newSecond(pairs []Pair) second {
	return second{pairs: pairs}
}
