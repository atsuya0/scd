package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/atsuya0/go-chooser"
)

type second struct {
	dataFile *os.File
	roots    []Root
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

func (s *second) remove(i int) error {
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
	if err := s.dataFile.Truncate(0); err != nil {
		return err
	}
	if _, err = s.dataFile.WriteAt(json, 0); err != nil {
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
	if _, name, err := nameChooser.SingleRun(); err != nil {
		return "", err
	} else {
		return name, nil
	}
}

func (s *second) chooseSubDir() (string, error) {
	_, root, err := s.getRoot()
	if err != nil {
		return "", err
	}
	pathChooser, err := chooser.NewChooser(root.Sub)
	if err != nil {
		return "", err
	}
	if _, dir, err := pathChooser.SingleRun(); err != nil {
		return "", err
	} else {
		return dir, nil
	}
}

func (s *second) removeSubDir() error {
	rootIndex, root, err := s.getRoot()
	if err != nil {
		return err
	}
	pathChooser, err := chooser.NewChooser(root.Sub)
	if err != nil {
		return err
	}
	index, _, err := pathChooser.SingleRun()
	if err != nil {
		return err
	} else if index < 0 {
		return nil
	}
	root.Sub = append(root.Sub[:index:index], root.Sub[index+1:]...)
	s.roots[rootIndex] = root
	return nil
}

func (s *second) getRoot() (int, Root, error) {
	wd, err := os.Getwd()
	if err != nil {
		return -1, Root{}, err
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return -1, Root{}, err
	}
	for i, root := range s.roots {
		if strings.HasPrefix(wd, strings.Replace(root.Path, "~", homeDir, 1)) {
			return i, root, nil
		}
	}
	return -1, Root{}, errors.New("This path is outside the scope.")
}

func (s *second) addSubDir() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	index, root, err := s.getRoot()
	if err != nil {
		return err
	}
	root.Sub = append(root.Sub, wd)
	s.roots[index] = root
	return nil
}

// Get an undeclared name error.
// second := second {roots: roots}
func newSecond(roots []Root) second {
	return second{roots: roots}
}
