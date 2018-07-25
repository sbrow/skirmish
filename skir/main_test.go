package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/sbrow/skirmish/skir/internal/version"
)

func init() {
	if os.Getenv("GOCMD") == "" {
		os.Setenv("GOCMD", "go")
	}
}
func TestSkir(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		/* 		{"main", []string{}, `Skir is a tool for developing the Skirmish card game.

		Usage:

			skir command [arguments]

		The commands are:

			card        show information about a specific card
			dump        save the current database to disk
			export      compile cards from the database to a specific format
			ps          fill out Photoshop templates
			recover     reload the database from disk
			sql         query the database
			version     print skir version

		Use "skir help [command]" for more information about a command.

		exit status 2
		`, true}, */
		{"version", []string{"version"}, "skir version " + version.Version, false},
		{"versionW/Arg", []string{"version", "arg"}, `usage: version
Run 'skir help version' for details.
exit status 2
`, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []string{
				"run",
				filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "sbrow", "skirmish",
					"skir", "main.go"),
			}
			args = append(args, tt.args...)
			out, err := exec.Command(os.Getenv("GOCMD"), args...).CombinedOutput()
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}
			if string(out) != tt.want {
				t.Errorf("wanted: \"%s\"\ngot \"%s\"\n", tt.want, string(out))
			}
		})
	}
}
