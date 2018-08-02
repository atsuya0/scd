package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	name bool
	path bool
}

func list(src string, options ListOptions) (err error) {
	file, err := os.Open(src)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(file)
	var source Source
	if err = decoder.Decode(&source); err != nil {
		return
	}

	if (options.name && options.path) || (!options.name && !options.path) {
		for _, pair := range source.Pairs {
			fmt.Println(pair)
		}
	} else if options.name {
		for _, pair := range source.Pairs {
			fmt.Println(pair.Name)
		}
	} else if options.path {
		for _, pair := range source.Pairs {
			fmt.Println(pair.Path)
		}
	}

	return
}

func createListCmd(src string) *cobra.Command {
	options := &ListOptions{}

	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List of second name.",
		Run: func(cmd *cobra.Command, _ []string) {
			if err := list(src, *options); err != nil {
				log.Fatalln("list:", err)
			}
		},
	}

	cmd.Flags().BoolVarP(&options.name, "name", "n", false, "Second name")
	cmd.Flags().BoolVarP(&options.path, "path", "p", false, "Target path")

	return cmd
}
