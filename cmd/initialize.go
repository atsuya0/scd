package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func initialize(src string) error {
	const yes = "yes"
	const no = "no"

	fmt.Printf("Can I initialize the data? ('%s' or '%s')\n>>> ", yes, no)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for scanner.Text() != yes && scanner.Text() != no {
		fmt.Print(">>> ")
		scanner.Scan()
	}

	if scanner.Text() == yes {
		if err := newSourceFile(src); err != nil {
			return err
		}
		fmt.Println("Processing was successful.")
	}

	return nil

}

func createInitializeCmd(src string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Initialization of data.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := initialize(src); err != nil {
				log.Fatalln("init:", err)
			}
		},
	}

	return cmd
}
