package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func createRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "second",
		Short: "You can switch path with the second name.",
	}

	cmd.AddCommand(createRegisterCmd())
	cmd.AddCommand(createChangeCmd())
	cmd.AddCommand(createListCmd())
	cmd.AddCommand(createDeleteCmd())
	cmd.AddCommand(createInitializeCmd())

	return cmd
}

func Execute() {
	cmd := createRootCmd()
	cmd.SetOutput(os.Stdout)

	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}
