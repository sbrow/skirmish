package ps

import (
	// "fmt"
	"github.com/sbrow/ps"
	// "github.com/sbrow/skirmish"
	"testing"
)

/*
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
	ApplyDataset("Combust_1")
	Save(true)
	ApplyDataset("Savage Melee_1")
	Save(true)
	ApplyDataset("Anger_1")
	Save(true)
	// ApplyDataset("High Ground_1")
	// ApplyDataset("Paranoia_1")
}

func TestDeckTemplate(t *testing.T) {
}

*/
func BenchmarkDeckTemplate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d := NewDeck(ps.Normal)
		d.ApplyDataset("Combust_1")
	}
}

// func BenchmarkNoTemplate(b *testing.B) {
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		ApplyDataset("Combust_1")
// 	}
// }
