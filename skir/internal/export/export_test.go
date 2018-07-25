package export

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/sbrow/skirmish/skir/internal/base"
)

func TestRun(t *testing.T) {
	type args struct {
		cmd  *base.Command
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"noArgs", args{CmdExport, []string{}}, `export \[format\]`},
		{"csv", args{CmdExport, []string{"csv"}}, `Generating dataset`},
		{"ue", args{CmdExport, []string{"ue"}}, `JSON`},
		{"twoArgs", args{CmdExport, []string{"csv", "ue"}}, `Generating dataset`},

		{"badArg", args{CmdExport, []string{"xml"}}, `format xml was not found`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.SetOutput(os.Stdout)
			// Capture stdout.
			stdCopy := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			os.Stdout = w
			outC := make(chan string)
			go func() {
				var buf bytes.Buffer
				_, err := io.Copy(&buf, r)
				r.Close()
				if err != nil {
					fmt.Fprintf(os.Stderr, "testing: copying pipe: %v\n", err)
					os.Exit(1)
				}
				outC <- buf.String()
			}()
			defer func() {
				w.Close()
				os.Stdout = stdCopy
				got := <-outC
				reg, err := regexp.Compile(tt.want)
				if err != nil {
					t.Fatal(err)
				}
				if !reg.MatchString(got) {
					t.Errorf("wanted: \n\"%s\"\ngot:\n\"%s\"", tt.want, got)
				}
			}()
			Run(tt.args.cmd, tt.args.args)
		})
	}
}
