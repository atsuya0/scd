package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
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

func Execute() {
	cmd := rootCmd()
	cmd.SetOutput(os.Stdout)
	cmd.Execute()
}
