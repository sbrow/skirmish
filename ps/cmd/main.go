package main

import (
	"fmt"
	app "github.com/sbrow/ps"
	"github.com/sbrow/skirmish/ps"
	"github.com/sbrow/skirmish/sql"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetPrefix("[ps] ")
	args := os.Args[1:]
	switch args[0] {
	case "crop":
	case "undo":
		err := app.DoAction("DK", strings.Title(args[0]))
		if err != nil {
			log.Panic(err)
		}
	case "deck":
		cards, err := sql.LoadMany(fmt.Sprintf("cards.leader='%s'", args[1]))
		if err != nil {
			log.Panic(err)
		}
		d := ps.NewDeck(app.Normal)
		defer d.Doc.Dump()
		for _, card := range cards {
			d.ApplyDataset(card.Name() + "_1")
			d.PNG(false)
		}
	default:
		sql.GenData()
		app.Wait("$ Import the current dataset file into photoshop," +
			" then press enter to continue")
		d := ps.NewDeck(app.Normal)
		d.ApplyDataset(strings.Join(args, " "))
		d.PNG(false)
	}
}
