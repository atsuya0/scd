package cmd

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

//go:embed scd.zsh
var zshScriptBytes []byte

func script(cmd *cobra.Command, _ []string) error {
	fmt.Println(string(zshScriptBytes))
	return nil
}

func scriptCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "script",
		Short: "Show the zsh script to be loaded in advance.",
		RunE:  script,
	}

	return cmd
}
