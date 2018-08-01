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

func (s *Source) match(name string) (string, error) {
	for _, pair := range s.Pairs {
		if pair.Name == name {
			return pair.Path, nil
		}
	}
	err := fmt.Errorf("%s is invalid name.", name)

	return "", err
}

func createRootCmd(src string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "second",
		Short: "You can switch path with the second name.",
	}

	cmd.AddCommand(createRegisterCmd(src))
	cmd.AddCommand(createChangeCmd(src))

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
