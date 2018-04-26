package ps

import (
	"bufio"
	"encoding/csv"
	"github.com/sbrow/ps"
	sk "github.com/sbrow/skirmish"
	"log"
	"os"
	"path/filepath"
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
}*/

func TestFormatTextBox(t *testing.T) {
	d := NewDeck(ps.Normal)
	d.FormatTextbox()
	d.Doc.Dump()
}

/*
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
*/
func TestDeckTemplate(t *testing.T) {
	d := NewDeck(ps.Normal)
	d.ApplyDataset("Rumour_1")
	d.PNG(true)
	// d.ApplyDataset("Paranoia_1")
	// d.Doc.Dump()
	// d.PNG(true)
	// d.FormatTextbox()
	// d.ApplyDataset("Anger_1")
	// d.PNG(true)
}

func TestNonDeckTemplate(t *testing.T) {
	n := NewNonDeck(ps.Normal)
	n.ApplyDataset("Lilith")
	n.FormatTextbox()
	n.Doc.Dump()
}

func TestEntireDeck(t *testing.T) {
	// Load Data
	f, err := os.Open(filepath.Join(sk.DataDir, "deckcards.csv"))
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
	cards, err := r.ReadAll()
	if err != nil {
		log.Panic(err)
	}

	//Parse Data
	id := -1
	ldr := -1
	for i, lbl := range cards[0] {
		if lbl == "id" {
			id = i
		}
		if lbl == "Bast" {
			ldr = i
			break
		}
	}

	d := NewDeck(ps.Normal)
	defer d.Doc.Dump()

	// Use Data
	for _, row := range cards[1:] {
		if row[ldr] == "true" {
			d.ApplyDataset(row[id])
			d.PNG(true)
		}
	}
}

//39, 45, 39
func TestText(t *testing.T) {
	n := NewDeck(ps.Normal)
	n.ApplyDataset("Rumour_1")
	n.AddSymbols()
	n.Doc.Dump()
}

func TestSize(t *testing.T) {
	v := 9.0
	log.Printf("%f\n", v)
	d := NewDeck(ps.Normal)
	lyr := d.Doc.LayerSet("Text").ArtLayer("type")
	lyr.TextItem.SetSize(v)
	log.Println(lyr.TextItem.Size())
}

// func BenchmarkDeckTemplate(b *testing.B) {
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		d := NewDeck(ps.Normal)
// 		d.ApplyDataset("Combust_1")
// 	}
// }

// 40s
func BenchmarkDeckInd(b *testing.B) {
	d := NewDeck(ps.Normal)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.DeckInd.Refresh()
	}
}

// 17s
func BenchmarkDeckIndRefresh(b *testing.B) {
	d := NewDeck(ps.Normal)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, lyr := range d.DeckInd.ArtLayers() {
			lyr.Refresh()
		}
	}
}
