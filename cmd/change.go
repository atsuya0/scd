package cmd

import (
	"log"
	"os"
	"strings"

	// "syscall"

	"github.com/spf13/cobra"
)

func change(cmd *cobra.Command, args []string) error {
	list, file, err := getListAndFile(os.O_RDONLY)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("%+v\n", err)
		}
	}()

	_, path, err := list.match(args[0])
	if err != nil {
		return err
	}
	cmd.Print(strings.Replace(path, "~", os.Getenv("HOME"), 1))
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
