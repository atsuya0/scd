package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func display(cmd *cobra.Command, args []string) error {
	list, file, err := getListAndFile(os.O_RDONLY)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("%v\n", err)
		}
	}()

	_, path, err := list.match(args[0])
	if err != nil {
		return err
	}
	cmd.Println(strings.Replace(path, "~", os.Getenv("HOME"), 1))

	return nil
}

func displayCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "display",
		Short: "Display the target path by the second name.",
		Args:  cobra.MinimumNArgs(1),
		RunE:  display,
	}

	return cmd
}
