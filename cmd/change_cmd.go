package cmd

import (
	// "syscall"

	"github.com/spf13/cobra"
)

func change(cmd *cobra.Command, args []string, sub bool) error {
	second, err := getSecond()
	if err != nil {
		return err
	}
	if err = second.dataFile.Close(); err != nil {
		return err
	}

	if sub {
		path, err := second.chooseSubDir()
		if err != nil {
			return err
		} else if path == "" {
			return nil
		}
		cmd.Print(path)
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

	_, path, err := second.match(name)
	if err != nil {
		return err
	}
	cmd.Print(path)
	// if err := os.Chdir(path); err != nil {
	// 	return err
	// }
	// shell := os.Getenv("SHELL")
	// if err := syscall.Exec(shell, []string{shell}, syscall.Environ()); err != nil {
	// 	return err
	// }

	return nil
}

func changeCmd() *cobra.Command {
	var sub bool
	var cmd = &cobra.Command{
		Use:   "change",
		Short: "Change the current working directory with the second name.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return change(cmd, args, sub)
		},
	}

	cmd.Flags().BoolVarP(&sub, "sub", "s", false, "sub directory")

	return cmd
}
