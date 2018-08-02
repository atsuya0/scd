package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	var cmd = &cobra.Command{
		Use:   "change",
		Short: "Change the working directory with the second name",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatalln("change:", fmt.Errorf("At least one argument is required."))
			}
			if err := change(src, args[0]); err != nil {
				log.Fatalln("change:", err)
			}
		},
	}

	return cmd
}
