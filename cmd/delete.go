package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func del(_ *cobra.Command, args []string) error {
	file, source, err := loadSource(os.O_RDWR)
	if err != nil {
		return fmt.Errorf("delete: %v", err)
	}
	defer file.Close()
	if err = file.Truncate(0); err != nil {
		return fmt.Errorf("delete: %v", err)
	}

	num, _, err := source.match(args[0])
	if err != nil {
		return fmt.Errorf("delete: %v", err)
	}
	source.del(num)

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
