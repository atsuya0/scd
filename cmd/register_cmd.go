package cmd

import (
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
	sub  bool
}

func register(options *RegisterOptions) error {
	second, err := newSecond()
	defer func() {
		if err = second.dataFile.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return err
	}

	if options.sub {
		err := second.addSubDir()
		if err != nil {
			return err
		}
	} else {
		if err = second.isDuplicate(*options); err != nil {
			return err
		}
		second.roots = append(second.roots, Root{Name: options.name, Path: options.path})
	}
	if err = second.flush(); err != nil {
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
		log.Fatalln(err)
	}
	user, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	cmd.Flags().StringVarP(
		&options.name, "name", "n", filepath.Base(wd), "the second name")
	cmd.Flags().StringVarP(
		&options.path, "path", "p", strings.Replace(wd, user.HomeDir, "~", 1),
		"the target path")
	cmd.Flags().BoolVarP(
		&options.sub, "sub", "s", false, "sub directory")

	return cmd
}
