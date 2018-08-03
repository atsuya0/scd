package cmd

import (
	"fmt"
	"log"
	"os"
	// "syscall"

	"github.com/spf13/cobra"
)

func change(src string, name string) (err error) {
	file, source := loadSource(src, os.O_RDONLY)
	defer file.Close()

	_, path, err := source.match(name)
	if err != nil {
		return
	}
	fmt.Println(path)
	// if err = os.Chdir(path); err != nil {
	// 	return
	// }
	// shell := os.Getenv("SHELL")
	// if err = syscall.Exec(shell, []string{shell}, syscall.Environ()); err != nil {
	// 	return
	// }

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
