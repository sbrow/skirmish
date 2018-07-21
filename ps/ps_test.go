package ps

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/sbrow/skirmish"
)

// TODO(sbrow): Deck and NonDeck Template tests are broken.
func TestDeckTemplate(t *testing.T) {
	want := filepath.Join(skirmish.Cfg.PS.Dir, skirmish.Cfg.PS.Deck)
	got := CardTemplate
	if got != want {
		t.Fatalf("Wanted: \"%s\"\nGot: \"%s\"\n", want, got)
	}
}

func TestNonDeckTemplate(t *testing.T) {
	want := filepath.Join(skirmish.Cfg.PS.Dir, skirmish.Cfg.PS.NonDeck)
	got := HeroTemplate
	if got != want {
		t.Fatalf("Wanted: \"%s\"\nGot: \"%s\"\n", want, got)
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
