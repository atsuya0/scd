package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func show(cmd *cobra.Command, args []string) error {
	file, source, err := loadSource(os.O_RDONLY)
	if err != nil {
		return fmt.Errorf("show: %v", err)
	}
	defer file.Close()

	_, path, err := source.match(args[0])
	if err != nil {
		return fmt.Errorf("show: %v", err)
	}
	cmd.Println(strings.Replace(path, "~", os.Getenv("HOME"), 1))

	return nil
}

func cmdShow() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "show",
		Short: "Show the second name",
		Args:  cobra.MinimumNArgs(1),
		RunE:  show,
	}

	return cmd
}
