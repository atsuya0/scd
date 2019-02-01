package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "second",
		Short: "Change the working directory with the second name.",
	}

	cmd.AddCommand(cmdRegister())
	cmd.AddCommand(cmdChange())
	cmd.AddCommand(cmdShow())
	cmd.AddCommand(cmdList())
	cmd.AddCommand(cmdDelete())
	cmd.AddCommand(cmdInitialize())

	return cmd
}

func Execute() {
	cmd := rootCmd()
	cmd.SetOutput(os.Stdout)
	cmd.Execute()
}
