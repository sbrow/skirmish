package ps

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// TestError is in its own file due to it's reliance on
// line numbers.
func TestError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"HelloWorld", errors.New("Hello, World"),
			fmt.Sprintf(" error at %s:29 Hello, World",
				filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "sbrow", "skirmish", "ps", "error_test.go"),
			),
		},
		{"nil", nil, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errors = make([]psError, 0)
			Error(tt.err)
			var got string
			if len(Errors) > 0 {
				got = Errors[len(Errors)-1].String()
			}
			if got != tt.want {
				t.Errorf("wanted: \"%s\"\ngot: \"%s\"\n", tt.want, got)
			}
		})
	}
}
