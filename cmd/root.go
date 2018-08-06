package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func rootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "second",
		Short: "You can switch path with the second name.",
	}

	cmd.AddCommand(cmdRegister())
	cmd.AddCommand(cmdChange())
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
