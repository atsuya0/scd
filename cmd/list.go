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

func list(options ListOptions) error {
	file, source := loadSource(os.O_RDWR)
	defer file.Close()

	if (options.name && options.path) || (!options.name && !options.path) {
		bytes, err := json.MarshalIndent(source.Pairs, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(bytes))
	} else if options.name {
		for _, pair := range source.Pairs {
			fmt.Println(pair.Name)
		}
	} else if options.path {
		for _, pair := range source.Pairs {
			fmt.Println(pair.Path)
		}
	}

	return nil
}

func createListCmd() *cobra.Command {
	options := &ListOptions{}

	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List of second name and target path.",
		Run: func(cmd *cobra.Command, _ []string) {
			if err := list(*options); err != nil {
				log.Fatalln("list:", err)
			}
		},
	}

	cmd.Flags().BoolVarP(&options.name, "name", "n", false, "Second name")
	cmd.Flags().BoolVarP(&options.path, "path", "p", false, "Target path")

	return cmd
}
