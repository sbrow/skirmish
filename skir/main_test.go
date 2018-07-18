package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/sbrow/skirmish/skir/internal/version"
)

// TODO(sbrow): Figure out how to test internal packages.
func Test_runVersion(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"noArgs", []string{"version"}, "skir version " + version.Version},
		{"oneArg", []string{"version", "arg"}, `usage: version
Run 'skir help version' for details.
exit status 2
`},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			var out, stErr bytes.Buffer
			args := []string{
				"run",
				filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "sbrow", "skirmish",
					"skir", "main.go"),
			}
			args = append(args, tt.args...)
			cmd := exec.Command("go", args...)
			cmd.Stdout = &out
			cmd.Stderr = &stErr
			var got string
			if err := cmd.Run(); err != nil {
				got = stErr.String()
			} else {
				got = out.String()
			}
			if got != tt.want {
				t.Errorf("wanted: \"%s\"\ngot \"%s\"\n", tt.want, got)
			}
		})
	}
}
