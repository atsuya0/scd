package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func del(src string, name string) (err error) {
	file, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(file)
	var source Source
	if err = decoder.Decode(&source); err != nil {
		return
	}

	num, _, err := source.match(name)
	if err != nil {
		return
	}
	source.del(num)

	jsonBytes, err := json.Marshal(source)
	if err != nil {
		return
	}
	if err = file.Truncate(0); err != nil {
		return
	}
	_, err = file.WriteAt(jsonBytes, 0)
	if err != nil {
		return
	}

	return
}

func createDeleteCmd(src string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete the second name",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatalln("delete:", fmt.Errorf("At least one argument is required."))
			}
			if err := del(src, args[0]); err != nil {
				log.Fatalln("delete:", err)
			}
		},
	}

	return cmd
}