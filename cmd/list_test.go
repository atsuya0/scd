package cmd

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	options := []ListOptions{{name: true}, {path: true}}
	outputs := []string{"binxd", "/usr/bin/etc/X11/xorg.conf.d"}

	path, err := filepath.Abs("./testdata/test.json")
	if err != nil {
		log.Fatalln(err)
	}
	if err := os.Setenv("SECOND_CMD_PATH", path); err != nil {
		log.Fatalln(err)
	}

	for i, opt := range options {
		buffer := &bytes.Buffer{}
		list(&opt, buffer)

		result := strings.Replace(buffer.String(), "\n", "", -1)
		if outputs[i] != result {
			t.Errorf("%s is not %s", result, outputs[i])
		}
	}
}
