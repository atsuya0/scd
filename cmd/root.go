package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "second",
		Short: "Change the current working directory with the second name.",
	}

	cmd.AddCommand(registerCmd())
	cmd.AddCommand(changeCmd())
	cmd.AddCommand(displayCmd())
	cmd.AddCommand(listCmd())
	cmd.AddCommand(removeCmd())
	cmd.AddCommand(initializeCmd())

	return cmd
}
