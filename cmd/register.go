package cmd

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type RegisterOptions struct {
	name string
	path string
}

func register(options *RegisterOptions) error {
	list, file, err := getListAndFile(os.O_RDWR)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("%+v\n", err)
		}
	}()
	if err != nil {
		return err
	}

	if err = list.isDuplicate(*options); err != nil {
		return err
	}

	list.Pairs = append(list.Pairs, Pair{Name: options.name, Path: options.path})
	jsonBytes, err := json.Marshal(list)
	if err != nil {
		return err
	}
	_, err = file.WriteAt(jsonBytes, 0)
	if err != nil {
		return err
	}

	return nil
}

func registerCmd() *cobra.Command {
	options := &RegisterOptions{}

	var cmd = &cobra.Command{
		Use:   "register",
		Short: "Attach the second name to the target path",
		RunE: func(_ *cobra.Command, _ []string) error {
			return register(options)
		},
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	user, err := user.Current()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	cmd.Flags().StringVarP(
		&options.name, "name", "n", filepath.Base(wd), "the second name")
	cmd.Flags().StringVarP(
		&options.path, "path", "p", strings.Replace(wd, user.HomeDir, "~", 1),
		"the target path")

	return cmd
}
