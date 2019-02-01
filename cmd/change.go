package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	// "syscall"

	"github.com/spf13/cobra"
)

func change(cmd *cobra.Command, args []string) error {
	file, source, err := loadSource(os.O_RDONLY)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return fmt.Errorf("change: %v", err)
	}

	_, path, err := source.match(args[0])
	if err != nil {
		return fmt.Errorf("change: %v", err)
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
		Short: "Change the working directory with the second name",
		Args:  cobra.MinimumNArgs(1),
		RunE:  change,
	}

	return cmd
}
