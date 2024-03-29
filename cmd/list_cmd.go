package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	name bool
	path bool
}

func list(options *ListOptions, out io.Writer) error {
	second, err := newSecond(false)
	if err != nil {
		return err
	}

	if (options.name && options.path) || (!options.name && !options.path) {
		bytes, err := json.MarshalIndent(second.roots, "", strings.Repeat(" ", 2))
		if err != nil {
			return err
		}
		fmt.Fprintln(out, string(bytes))
	} else if options.name {
		for _, root := range second.roots {
			fmt.Fprintln(out, root.Name)
		}
	} else if options.path {
		for _, root := range second.roots {
			fmt.Fprintln(out, root.Path)
		}
	}

	return nil
}

func listCmd() *cobra.Command {
	options := &ListOptions{}

	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List the second name and the target path in JSON format.",
		RunE: func(_ *cobra.Command, _ []string) error {
			return list(options, os.Stdout)
		},
	}

	cmd.Flags().BoolVarP(&options.name, "name", "n", false, "the second name")
	cmd.Flags().BoolVarP(&options.path, "path", "p", false, "the target path")

	return cmd
}
