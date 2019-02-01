package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	name bool
	path bool
}

func list(options *ListOptions, out io.Writer) error {
	file, source, err := loadSource(os.O_RDWR)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return fmt.Errorf("list: %v", err)
	}

	if (options.name && options.path) || (!options.name && !options.path) {
		bytes, err := json.MarshalIndent(source.Pairs, "", "  ")
		if err != nil {
			return fmt.Errorf("list: %v", err)
		}
		fmt.Fprintln(out, string(bytes))
	} else if options.name {
		for _, pair := range source.Pairs {
			fmt.Fprintln(out, pair.Name)
		}
	} else if options.path {
		for _, pair := range source.Pairs {
			fmt.Fprintln(out, pair.Path)
		}
	}

	return nil
}

func cmdList() *cobra.Command {
	options := &ListOptions{}

	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List of second name and target path.",
		RunE: func(_ *cobra.Command, _ []string) error {
			return list(options, os.Stdout)
		},
	}

	cmd.Flags().BoolVarP(&options.name, "name", "n", false, "Second name")
	cmd.Flags().BoolVarP(&options.path, "path", "p", false, "Target path")

	return cmd
}
