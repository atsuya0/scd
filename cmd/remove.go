package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func remove(_ *cobra.Command, args []string) error {
	list, file, err := getListAndFile(os.O_RDWR)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("%v\n", err)
		}
	}()

	if err = file.Truncate(0); err != nil {
		return err
	}

	num, _, err := list.match(args[0])
	if err != nil {
		return err
	}
	if err := list.del(num); err != nil {
		return err
	}

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

func removeCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove the second name.",
		Args:  cobra.MinimumNArgs(1),
		RunE:  remove,
	}

	return cmd
}
