package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func createInitCmd(src string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Initialization of data.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := newSourceFile(src); err != nil {
				log.Fatalln("init:", err)
			}
		},
	}

	return cmd
}
