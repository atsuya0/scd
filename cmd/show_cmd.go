package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func show(cmd *cobra.Command, args []string) error {
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
