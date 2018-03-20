package ps

import (
	"github.com/sbrow/ps"
	"github.com/sbrow/skirmish"
	"testing"
)

func TestFormatTitle(t *testing.T) {
	ps.Open(skirmish.Template)
	err := FormatTitle()
	if err != nil {
		t.Fatal(err)
	}
}

func TestFormatTextBox(t *testing.T) {
	FormatTextbox()
}

func TestFormatSpeed(t *testing.T) {
	FormatSpeed()
}

func TestSetLeader(t *testing.T) {
	SetLeader("Lilith")
}
