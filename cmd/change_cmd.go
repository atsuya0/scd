package cmd

import (
	"os"
	"strings"

	// "syscall"

	"github.com/spf13/cobra"
)

func change(cmd *cobra.Command, args []string) error {
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
	cmd.Print(strings.Replace(path, "~", home, 1))
	// if err := os.Chdir(path); err != nil {
	// 	return err
	// }
	// shell := os.Getenv("SHELL")
	// if err := syscall.Exec(shell, []string{shell}, syscall.Environ()); err != nil {
	// 	return err
	// }

	return nil
}

func changeCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "change",
		Short: "Change the current working directory with the second name.",
		RunE:  change,
	}

	return cmd
}
