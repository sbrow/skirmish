package ps

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/sbrow/skirmish"
)

func TestTemplate(t *testing.T) {
	want := filepath.Join(os.Getenv("SK_PS"), "Template009.1.psd")
	got := filepath.Join(skirmish.Cfg.PS.Dir, skirmish.Cfg.PS.Deck)
	if got != want {
		t.Fatalf("Wanted: \"%s\"\nGot: \"%s\"\n", want, got)
	}
}

func TestHeroTemplate(t *testing.T) {
	want := filepath.Join(os.Getenv("SK_PS"), "Template009.1h.psd")
	got := filepath.Join(skirmish.Cfg.PS.Dir, skirmish.Cfg.PS.NonDeck)
	if got != want {
		t.Fatalf("Wanted: \"%s\"\nGot: \"%s\"\n", want, got)
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"HelloWorld", errors.New("Hello, World"),
			fmt.Sprintf(" error at %s:44 Hello, World",
				filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "sbrow", "skirmish", "ps", "ps_test.go"),
			),
		},
		// TODO(sbrow): Add error to ps.TestError that sources a different file.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.err)
			got := Errors[len(Errors)-1].String()
			if got != tt.want {
				t.Errorf("wanted: \"%s\"\ngot: \"%s\"\n", tt.want, got)
			}
		})
	}
}

func TestGetTolerances(t *testing.T) {
	tests := []struct {
		name       string
		tolerances map[string]int
	}{
		{"", map[string]int{"title": 55, "short": 17, "long": 13, "flavor": 80, "bottom": 65}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTolerances()
			if !reflect.DeepEqual(Tolerances, tt.tolerances) {
				t.Errorf("wanted: %+v\ngot: %+v", Tolerances, tt.tolerances)
			}
		})
	}
}
