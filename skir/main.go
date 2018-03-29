// Command skir is the primary skirmish command.
//
// Build
//
// Generates photoshop dataset (csv) files with data pulled from the PostgresSQL
// database.
//
//		skir build
//		skir build regex / skirmish build -r
//		skir build bold  / skirmish build -b
//		skir build data  / skirmish build -d
//
// PS
//
// Fills out a document (psd) template with data from a dataset (csv) file and
//
// 		skir ps
//		skir ps -card=$CARDNAME // Builds the psd for the given card
//		skir ps -deck=$DECKNAME
//		skir ps action $ACTIONNAME
//
// Card
//
// loads card information from the database and outputs it to STDOUT.
//
//		skir card $CARDNAME
//
// TODO: config file for non-programmers
package main

import (
	"flag"
	"fmt"
	app "github.com/sbrow/ps"
	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/build"
	"github.com/sbrow/skirmish/ps"
	"github.com/sbrow/skirmish/sql"
	"log"
	"os"
	"strings"
)

func main() {
	card := flag.String("card", "", "card get info on a card.")
	flag.Parse()
	args := []string{}
	cmd := ""
	switch {
	case len(os.Args) > 2:
		args = os.Args[2:]
		fallthrough
	case len(os.Args) > 1:
		cmd = os.Args[1]
	}
	switch {
	case cmd == "ps":
		switch args[0] {
		case "crop":
		case "undo":
			err := app.DoAction("DK", strings.Title(args[0]))
			if err != nil {
				panic(err)
			}
		case "save":
			// ps.Save(true, args...)
		default:
			ps.Format()
		}
	case cmd == "card" || *card != "":
		name := *card
		if name == "" {
			name = strings.Join(args, " ")
		}
		card, err := sql.Load(name)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(card)
	case cmd == "build":
		log.SetPrefix("[main] ")
		// if !*fast {
		log.Println("Generating cards")
		build.Data()
		// }
		// log.SetPrefix("[photoshop] ")
		// build.PSDs()
		log.Println("Cards successfully generated!")
	case cmd == "db":
		opt := args[0]
		if opt == "save" {
			sql.Dump(skirmish.DataDir)
		} else if opt == "load" {
			sql.Recover(skirmish.DataDir)
		}
	}
}
