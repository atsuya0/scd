package cmd

import "errors"

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
	err := errors.New(name + " invalid name")

	return 0, "", err
}

func (l *List) isDuplicate(options RegisterOptions) (err error) {
	for _, pair := range l.Pairs {
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

func (l *List) del(i int) error {
	if 0 <= i && i < len(l.Pairs) {
		l.Pairs = append(l.Pairs[:i:i], l.Pairs[i+1:]...)
		return nil
	}
	return errors.New("out of range.")
}
