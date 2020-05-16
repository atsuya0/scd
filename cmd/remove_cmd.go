package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func remove(_ *cobra.Command, args []string) error {
	second, err := getSecond()
	defer func() {
		if err = second.file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return err
	}

	num, _, err := second.match(args[0])
	if err != nil {
		return err
	}
	if err := second.del(num); err != nil {
		return err
	}
	if err = second.update(); err != nil {
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
