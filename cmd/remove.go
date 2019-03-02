package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func remove(_ *cobra.Command, args []string) error {
	file, list, err := loadList(os.O_RDWR)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return fmt.Errorf("remove: %v", err)
	}

	if err = file.Truncate(0); err != nil {
		return fmt.Errorf("remove: %v", err)
	}

	num, _, err := list.match(args[0])
	if err != nil {
		return fmt.Errorf("remove: %v", err)
	}
	if err := list.del(num); err != nil {
		return fmt.Errorf("remove: %v", err)
	}

	jsonBytes, err := json.Marshal(list)
	if err != nil {
		return fmt.Errorf("remove: %v", err)
	}
	_, err = file.WriteAt(jsonBytes, 0)
	if err != nil {
		return fmt.Errorf("remove: %v", err)
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
