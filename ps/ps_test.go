package ps

import (
	"github.com/sbrow/ps"
	"github.com/sbrow/skirmish"
	"testing"
)

func TestSetLeader(t *testing.T) {
	SetLeader("Tinsel")
}

func TestFormatTitle(t *testing.T) {
	ps.Open(skirmish.Template)
	err := FormatTitle()
	if err != nil {
		t.Fatal(err)
	}
}

func TestFormatTextBox(t *testing.T) {
	FormatTextbox()
	doc.Dump()
}

func TestApplyDataset(t *testing.T) {
	ApplyDataset("Anger_1")
}

func BenchmarkManualLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Format()
	}
}
