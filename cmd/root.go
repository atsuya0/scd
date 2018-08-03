package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type Pair struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Source struct {
	Pairs []Pair `json:"source"`
}

func (s *Source) match(name string) (int, string, error) {
	for i, pair := range s.Pairs {
		if pair.Name == name {
			return i, pair.Path, nil
		}
	}
	err := fmt.Errorf("%s is invalid name.", name)

	return 0, "", err
}

func (s *Source) isDuplicate(options RegisterOptions) (err error) {
	for _, pair := range s.Pairs {
		if pair.Name == options.name {
			err = fmt.Errorf("This name has already been registered.")
			return
		}
		if pair.Path == options.path {
			err = fmt.Errorf("This path has already been registered.")
			return
		}
	}
	return
}

func (s *Source) del(i int) {
	s.Pairs = append(s.Pairs[:i:i], s.Pairs[i+1:]...)
}

func newSourceFile(src string) error {
	file, err := os.Create(src)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonBytes, err := json.Marshal(Source{Pairs: []Pair{}})
	if err != nil {
		return err
	}

	_, err = file.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}

func loadSource(src string, flag int) (*os.File, Source) {
	if _, err := os.Stat(src); err != nil {
		if err := newSourceFile(src); err != nil {
			log.Fatalln(err)
		}
	}

	file, err := os.OpenFile(src, flag, 0600)
	if err != nil {
		log.Fatalln(err)
	}

	decoder := json.NewDecoder(file)
	source := Source{}
	if err = decoder.Decode(&source); err != nil {
		log.Fatalln(err)
	}

	return file, source
}

func createRootCmd() *cobra.Command {
	src := os.Getenv("SECOND_CMD_PATH")

	var cmd = &cobra.Command{
		Use:   "second",
		Short: "You can switch path with the second name.",
	}

	cmd.AddCommand(createRegisterCmd(src))
	cmd.AddCommand(createChangeCmd(src))
	cmd.AddCommand(createListCmd(src))
	cmd.AddCommand(createDeleteCmd(src))
	cmd.AddCommand(createInitCmd(src))

	return cmd
}

func Execute() {
	cmd := createRootCmd()
	cmd.SetOutput(os.Stdout)

	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}
