package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func show(cmd *cobra.Command, args []string) error {
	roots, err := getRoots()
	if err != nil {
		return err
	}
	second := newSecond(roots)

	var name string
	if len(args) != 0 {
		name = args[0]
	} else {
		name, err = second.choose()
		if err != nil {
			return err
		} else if name == "" {
			return nil
		}
	}

	_, path, err := second.match(name)
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	cmd.Println(strings.Replace(path, "~", home, 1))

	return nil
}

func showCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "show",
		Short: "Show the target path by the second name.",
		RunE:  show,
	}

	return cmd
}
