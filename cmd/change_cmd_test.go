package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestChange(t *testing.T) {
	patterns := []struct {
		cmd    string
		output string
	}{
		{cmd: "scd change bin", output: "/usr/bin"},
		{cmd: "scd change xd", output: "/etc/X11/xorg.conf.d"},
	}

	path, err := filepath.Abs("./testdata")
	if err != nil {
		t.Error(err)
	}
	if err := os.Setenv("SCD_DATA_PATH", path); err != nil {
		t.Error(err)
	}

	cmd := NewCmd()

	for _, p := range patterns {
		buffer := bytes.NewBuffer(nil)
		cmd.SetOutput(buffer)

		args := strings.Split(p.cmd, " ")
		cmd.SetArgs(args[1:])
		cmd.Execute()

		result := buffer.String()
		if p.output != result {
			t.Errorf("'%s' is not '%s'", result, p.output)
		}
	}
}
