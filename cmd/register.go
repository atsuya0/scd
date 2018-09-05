package cmd

import (
	"encoding/json"
	"fmt"
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
	file, source, err := loadSource(os.O_RDWR)
	if err != nil {
		return fmt.Errorf("register: %v", err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err = source.isDuplicate(*options); err != nil {
		return fmt.Errorf("register: %v", err)
	}

	source.Pairs = append(source.Pairs, Pair{Name: options.name, Path: options.path})
	jsonBytes, err := json.Marshal(source)
	if err != nil {
		return fmt.Errorf("register: %v", err)
	}
	_, err = file.WriteAt(jsonBytes, 0)
	if err != nil {
		return fmt.Errorf("register: %v", err)
	}

	return nil
}

func cmdRegister() *cobra.Command {
	options := &RegisterOptions{}

	var cmd = &cobra.Command{
		Use:   "register",
		Short: "Attach the second name to path",
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
		&options.name, "name", "n", filepath.Base(wd), "Second name")
	cmd.Flags().StringVarP(
		&options.path, "path", "p", strings.Replace(wd, user.HomeDir, "~", 1),
		"Target path")

	return cmd
}
