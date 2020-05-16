package cmd

import (
	"os"
	"strings"

	// "syscall"

	"github.com/spf13/cobra"
)

func change(cmd *cobra.Command, args []string) error {
	pairs, err := getPairs()
	if err != nil {
		return err
	}
	second := newSecond(pairs)

	_, path, err := second.match(args[0])
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
		Args:  cobra.MinimumNArgs(1),
		RunE:  change,
	}

	return cmd
}
