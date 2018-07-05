package ps

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/sbrow/skirmish"
)

func TestTemplate(t *testing.T) {
	want := `F:\GitLab\dreamkeepers-psd\Template009.1.psd`
	got := filepath.Join(skirmish.Config.PSD.Dir, skirmish.Config.PSD.Deck)
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
		{"HelloWorld", errors.New("Hello, World"), " error at C:/Users/Spencer/go/src/github.com/sbrow/skirmish/ps/ps_test.go:75 Hello, World"},
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

/*
func TestHeroTemplate(t *testing.T) {
	want := `F:\GitLab\dreamkeepers-psd\Template009.1h.psd`
	got := HeroTemplate
	if got != want {
		t.Fatalf("Wanted: \"%s\"\nGot: \"%s\"\n", want, got)
	}
}

func TestSetLeader(t *testing.T) {
	tmp2.SetLeader("Bast")
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
*/
