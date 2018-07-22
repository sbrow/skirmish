package ps

import (
	"runtime"
	"testing"

	"github.com/sbrow/ps"
)

func TestDeckTemplate_SetLeader(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Photoshop is likely not installed")
	}
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
