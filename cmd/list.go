package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	name bool
	path bool
}

func list(cmd *cobra.Command, options *ListOptions) error {
	file, source, err := loadSource(os.O_RDWR)
	if err != nil {
		return fmt.Errorf("list: %v", err)
	}
	defer file.Close()

	if (options.name && options.path) || (!options.name && !options.path) {
		bytes, err := json.MarshalIndent(source.Pairs, "", "  ")
		if err != nil {
			return fmt.Errorf("list: %v", err)
		}
		cmd.Println(string(bytes))
	} else if options.name {
		for _, pair := range source.Pairs {
			cmd.Println(pair.Name)
		}
	} else if options.path {
		for _, pair := range source.Pairs {
			cmd.Println(pair.Path)
		}
	}

	return nil
}

func cmdList() *cobra.Command {
	options := &ListOptions{}

	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List of second name and target path.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return list(cmd, options)
		},
	}

	cmd.Flags().BoolVarP(&options.name, "name", "n", false, "Second name")
	cmd.Flags().BoolVarP(&options.path, "path", "p", false, "Target path")

	return cmd
}
