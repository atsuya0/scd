package main

import (
	"log"
	"os"

	"github.com/atsuya0/scd/cmd"
)

func main() {
	command := cmd.NewCmd()
	command.SetOutput(os.Stdout)
	if err := command.Execute(); err != nil {
		log.Fatalf("%+v\n", err)
	}
}
