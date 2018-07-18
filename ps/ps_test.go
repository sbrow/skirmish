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
			fmt.Sprintf(" error at %s:46 Hello, World",
				filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "sbrow", "skirmish", "ps", "ps_test.go"),
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

func TestGetTolerances(t *testing.T) {
	tests := []struct {
		name       string
		tolerances map[string]int
		wantErr    bool
	}{
		{"DB", map[string]int{
			"title":  55,
			"short":  17,
			"long":   13,
			"flavor": 80,
			"bottom": 65,
		}, false},

		{"WrongDB", map[string]int{}, true},
		// {"EmptyTable", map[string]int{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "WrongDB" {
				skirmish.Connect(skirmish.LocalDB.DBArgs())
			}
			if err := GetTolerances(); (err != nil) != tt.wantErr {
				t.Errorf("GetTolerances() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(Tolerances, tt.tolerances) {
				t.Errorf("wanted: %+v\ngot: %+v", Tolerances, tt.tolerances)
			}
		})
	}
}
