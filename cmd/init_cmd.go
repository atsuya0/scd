package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func initialize(_ *cobra.Command, _ []string) error {
	const yes = "yes"
	const no = "no"

	fmt.Printf("Can I initialize the data? ('%s' or '%s')\n>>> ", yes, no)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for scanner.Text() != yes && scanner.Text() != no {
		fmt.Print(">>> ")
		scanner.Scan()
	}

	if scanner.Text() == no {
		return nil
	}

	path, err := getDataFile()
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := file.WriteString("[]"); err != nil {
		return err
	}

	fmt.Println("Processing was successful.")

	return nil

}

func initializeCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize the data.",
		RunE:  initialize,
	}

	return cmd
}
