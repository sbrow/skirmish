package ps

import (
	"testing"

	"github.com/sbrow/ps"
)

var tmp *Template

func init() {
	tmp = New(ps.Normal, CardTemplate)
}

func TestTemplate(t *testing.T) {
	want := `F:\GitLab\dreamkeepers-psd\Template009.1.psd`
	got := CardTemplate
	if got != want {
		t.Fatalf("Wanted: \"%s\"\nGot: \"%s\"\n", want, got)
	}
}

func TestHeroTemplate(t *testing.T) {
	want := `F:\GitLab\dreamkeepers-psd\Template009.1h.psd`
	got := HeroTemplate
	if got != want {
		t.Fatalf("Wanted: \"%s\"\nGot: \"%s\"\n", want, got)
	}
}
func TestSetLeader(t *testing.T) {
	tmp.SetLeader("Tinsel")
}

func TestFormatTitle(t *testing.T) {
	d := NewDeck(ps.Normal)
	d.FormatTitle()
	d.Doc.Dump()
}

func TestFormatTextBox(t *testing.T) {
	d := NewDeck(ps.Normal)
	defer d.Doc.Dump()
	d.ApplyDataset("Combust_1")
	d.FormatTextbox()
}

func TestApplyDataset(t *testing.T) {
	d := NewDeck(ps.Normal)
	defer d.Doc.Dump()
	tests := []string{"Combust_1"}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			d.ApplyDataset(tt)
			d.PNG(true)
		})
	}
}
