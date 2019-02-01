package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func del(_ *cobra.Command, args []string) error {
	file, source, err := loadSource(os.O_RDWR)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return fmt.Errorf("delete: %v", err)
	}

	if err = file.Truncate(0); err != nil {
		return fmt.Errorf("delete: %v", err)
	}

	num, _, err := source.match(args[0])
	if err != nil {
		return fmt.Errorf("delete: %v", err)
	}
	if err := source.del(num); err != nil {
		return fmt.Errorf("delete: %v", err)
	}

	jsonBytes, err := json.Marshal(source)
	if err != nil {
		return fmt.Errorf("delete: %v", err)
	}
	_, err = file.WriteAt(jsonBytes, 0)
	if err != nil {
		return fmt.Errorf("delete: %v", err)
	}

	return nil
}

func cmdDelete() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "delete the second name",
		Args:  cobra.MinimumNArgs(1),
		RunE:  del,
	}

	return cmd
}
