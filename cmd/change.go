package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
)

func change(src string, name string) (err error) {
	file, err := os.Open(src)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(file)
	var source Source
	if err = decoder.Decode(&source); err != nil {
		return
	}

	shell := os.Getenv("SHELL")
	path, err := source.match(name)
	if err != nil {
		return
	}
	if err = os.Chdir(path); err != nil {
		return
	}
	if err = syscall.Exec(shell, []string{shell}, syscall.Environ()); err != nil {
		return
	}

	return
}

func createChangeCmd(src string) *cobra.Command {
	var name string

	var cmd = &cobra.Command{
		Use:   "change",
		Short: "Change the working directory with the second name",
		Run: func(cmd *cobra.Command, args []string) {
			if err := change(src, name); err != nil {
				log.Fatalln("change", err)
			}
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", filepath.Base(os.Getenv("HOME")), "Second name")

	return cmd
}
