package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
	return nil
}

func (s *Source) del(i int) {
	s.Pairs = append(s.Pairs[:i:i], s.Pairs[i+1:]...)
}

func createRootCmd(src string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "second",
		Short: "You can switch path with the second name.",
	}

	cmd.AddCommand(createRegisterCmd(src))
	cmd.AddCommand(createChangeCmd(src))
	cmd.AddCommand(createListCmd(src))
	cmd.AddCommand(createDeleteCmd(src))

	return cmd
}

func Execute() {
	src := os.Getenv("HOME") + "/.second_names"

	cmd := createRootCmd(src)
	cmd.SetOutput(os.Stdout)

	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}
