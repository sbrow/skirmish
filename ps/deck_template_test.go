package ps

import (
	"testing"

	ps "github.com/sbrow/ps/v2"
)

func TestDeckTemplate_SetLeader(t *testing.T) {
	Skip(t)
	tests := []struct {
		name string
	}{
		{"Bast"},
		{"Igrath"},
		{"Lilith"},
		{"Scuttler"},
	}
	_ = NewDeck(ps.Safe)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDeck(ps.Fast)
			d.SetLeader(tt.name)
		})
	}
}
