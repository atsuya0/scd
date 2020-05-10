package cmd

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func show(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("At least one argument is required.")
	}
	list, file, err := getListAndFile(os.O_RDONLY)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	_, path, err := list.match(args[0])
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
		Args:  cobra.MinimumNArgs(1),
		RunE:  show,
	}

	return cmd
}
