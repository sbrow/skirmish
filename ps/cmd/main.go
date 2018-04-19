package main

import (
	"flag"
	"fmt"
	app "github.com/sbrow/ps"
	sk "github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/ps"
	"github.com/sbrow/skirmish/sql"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	fast := flag.Bool("f", false, "fast skip dataset generation.")
	flag.Parse()
	log.SetPrefix("[ps] ")
	log.Println("Opening Photoshop")
	app.Open(sk.Template)
	args := os.Args[1:]
	switch args[0] {
	case "crop":
	case "undo":
		err := app.DoAction("DK", strings.Title(args[0]))
		if err != nil {
			log.Panic(err)
		}
	case "deck":
		cards, err := sql.LoadMany(fmt.Sprintf(
			"cards.leader='%s' ORDER BY name ASC", args[1]))
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
		// TODO: run GenData in a separate Goroutine, syncing with a WaitGroup.
		var wg sync.WaitGroup
		if !*fast {
			wg.Add(1)
			go func() {
				defer wg.Done()
				sql.GenData()
			}()
		}
		app.Wait("$ Import the current dataset file into photoshop," +
			" then press enter to continue\r\n")
		wg.Wait()
		d := ps.NewDeck(app.Normal)
		d.ApplyDataset(strings.Join(args, " "))
		d.PNG(false)
	}
	fmt.Println(time.Since(start))
}
