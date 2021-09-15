package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func remove(args []string, sub bool) error {
	second, err := getSecond()
	defer func() {
		if err = second.dataFile.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return err
	}

	if sub {
		if err := second.removeSubDir(); err != nil {
			return err
		}
		if err = second.flush(); err != nil {
			return err
		}
		return nil
	}

	var name string
	if len(args) != 0 {
		name = args[0]
	} else {
		name, err = second.choose()
		if err != nil {
			return err
		} else if name == "" {
			return nil
		}
	}

	num, _, err := second.match(name)
	if err != nil {
		return err
	}
	if err := second.remove(num); err != nil {
		return err
	}
	if err = second.flush(); err != nil {
		return err
	}

	return nil
}

func removeCmd() *cobra.Command {
	var sub bool
	var cmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove the second name.",
		RunE: func(_ *cobra.Command, args []string) error {
			return remove(args, sub)
		},
	}
	cmd.Flags().BoolVarP(&sub, "sub", "s", false, "sub directory")

	return cmd
}
