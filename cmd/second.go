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
	roots []Root
}

type Root struct {
	Name string   `json:"name"`
	Path string   `json:"path"`
	Sub  []string `json:"sub"`
}

func (s *second) match(name string) (int, string, error) {
	for i, root := range s.roots {
		if root.Name == name {
			return i, root.Path, nil
		}
	}
	err := errors.New(name + " invalid name")

	return 0, "", err
}

func (s *second) isDuplicate(options RegisterOptions) (err error) {
	for _, root := range s.roots {
		if root.Name == options.name {
			err = errors.New("This name has already been registered.")
			return
		}
		if root.Path == options.path {
			err = errors.New("This path has already been registered.")
			return
		}
	}
	return
}

func (s *second) del(i int) error {
	if 0 <= i && i < len(s.roots) {
		s.roots = append(s.roots[:i:i], s.roots[i+1:]...)
		return nil
	}
	return errors.New("out of range.")
}

func (s *second) flush() error {
	json, err := json.MarshalIndent(s.roots, "", strings.Repeat(" ", 2))
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
	for _, root := range s.roots {
		names = append(names, root.Name)
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
// second := second {roots: roots}
func newSecond(roots []Root) second {
	return second{roots: roots}
}
