package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type RegisterOptions struct {
	name string
	path string
}

func register(options RegisterOptions) (err error) {
	file, source := loadSource(os.O_RDWR)
	defer file.Close()
	if err = source.isDuplicate(options); err != nil {
		return
	}

	source.Pairs = append(source.Pairs, Pair{Name: options.name, Path: options.path})
	jsonBytes, err := json.Marshal(source)
	if err != nil {
		return
	}
	_, err = file.WriteAt(jsonBytes, 0)
	if err != nil {
		return
	}

	return
}

func createRegisterCmd() *cobra.Command {
	options := &RegisterOptions{}

	var cmd = &cobra.Command{
		Use:   "register",
		Short: "Attach the second name to path",
		Run: func(cmd *cobra.Command, args []string) {
			if err := register(*options); err != nil {
				log.Fatalln("register", err)
			}
		},
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	cmd.Flags().StringVarP(&options.name, "name", "n", filepath.Base(wd), "Second name")
	cmd.Flags().StringVarP(&options.path, "path", "p", wd, "Target path")

	return cmd
}
