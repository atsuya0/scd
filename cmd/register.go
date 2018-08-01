package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type Options struct {
	name string
	path string
}

func isDuplicate(pairs []Pair, options Options) (err error) {
	for _, pair := range pairs {
		if pair.Name == options.name {
			err = fmt.Errorf("This name has already been registered.")
			return
		}
		if pair.Path == options.path {
			err = fmt.Errorf("This path has already been registered.")
			return
		}
	}
	return nil
}

func register(src string, options Options) (err error) {
	file, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(file)
	var source Source
	if err = decoder.Decode(&source); err == nil {
		if err = isDuplicate(source.Pairs, options); err != nil {
			return
		}
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

	return nil
}

func createRegisterCmd(src string) *cobra.Command {
	options := &Options{}

	var cmd = &cobra.Command{
		Use:   "register",
		Short: "Attach the second name to path",
		Run: func(cmd *cobra.Command, args []string) {
			if err := register(src, *options); err != nil {
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
