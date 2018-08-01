package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func createRootCmd(src string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "second",
		Short: "You can switch path with the second name.",
	}

	cmd.AddCommand(createRegisterCmd(src))
	// cmd.AddCommand(createListCmd(trashPath))

	return cmd
}

func Execute() {
	// src := os.Getenv("HOME") + "/second_names"
	src := os.Getenv("GOPATH") + "/src/cli/second/second_names"

	cmd := createRootCmd(src)
	cmd.SetOutput(os.Stdout)

	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}
