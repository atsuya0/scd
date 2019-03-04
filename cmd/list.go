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
	list, file, err := getListAndFile(os.O_RDONLY)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("%+v\n", err)
		}
	}()
	if err != nil {
		return err
	}

	if (options.name && options.path) || (!options.name && !options.path) {
		bytes, err := json.MarshalIndent(list.Pairs, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(out, string(bytes))
	} else if options.name {
		for _, pair := range list.Pairs {
			fmt.Fprintln(out, pair.Name)
		}
	} else if options.path {
		for _, pair := range list.Pairs {
			fmt.Fprintln(out, pair.Path)
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
