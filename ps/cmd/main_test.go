package main

import (
	app "github.com/sbrow/ps"
	"github.com/sbrow/skirmish/ps"
	// "github.com/sbrow/skirmish/sql"
	// "sync"
	"testing"
)

func TestRefresh(t *testing.T) {
	d := ps.NewDeck(app.Normal)
	d.DeckInd.Refresh()
}

func TestText(t *testing.T) {
	d := ps.NewDeck(app.Normal)
	// d.ApplyDataset("Chaotic Blast_1")
	d.FormatTextbox()
}

/*
func BenchmarkGoRoutine(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		sql.GenData()
	}()
	ps.Wait("$ Import the current dataset file into photoshop," +
		" then press enter to continue")
	wg.Wait()
}

func BenchmarkSTD(b *testing.B) {
	sql.GenData()
	ps.Wait("$ Import the current dataset file into photoshop," +
		" then press enter to continue")
}
*/
