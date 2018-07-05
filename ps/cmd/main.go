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
	fast := flag.Bool("f", false, " fast mode- skip dataset generation.")
	flag.Parse()
	log.SetPrefix("[ps] ")
	log.Println("Opening Photoshop")
	app.Open(sk.Template)
	args := os.Args[1:]
	var leaders []string
	var condition string

	var wg sync.WaitGroup
	if !*fast {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sql.GenData()
		}()
	}

	switch args[0] {
	case "crop":
	case "undo":
		wg.Wait()
		err := app.DoAction("DK", strings.Title(args[0]))
		if err != nil {
			log.Panic(err)
		}
	case "all":
		leaders = make([]string, len(sk.Leaders))
		for i, ldr := range sk.Leaders {
			leaders[i] = ldr.Name
		}
		condition = "NOT cards.leader=NULL"
		fallthrough
	case "deck":
		if len(leaders) == 0 {
			leaders = []string{args[1]}
			condition = fmt.Sprintf("cards.leader='%s'", args[1])
		}
		order := "cards.leader, cards.supertypes, cards.type, char_length(name) ASC"
		cards, err := sql.LoadMany(fmt.Sprintf("%s ORDER BY %s", condition, order))
		if err != nil {
			log.Panic(err)
		}
		d := ps.NewDeck(app.Fast)
		defer d.Doc.Dump()
		defer app.Close(app.PSSaveChanges)
		app.Wait("$ Import the current dataset file into photoshop," +
			" then press enter to continue")
		wg.Wait()
		for _, ldr := range leaders {
			d.SetLeader(ldr)
			for _, card := range cards {
				imgs, err := card.Images()
				if err != nil {
					log.Println("ERROR:", err)
					imgs = []string{"1"}
				}
				for i := range imgs {
					d.ApplyDataset(fmt.Sprintf("%s_%d", card.Name(), i+1))
					d.PNG(false)
				}
			}
		}
	default:
		app.Wait("$ Import the current dataset file into photoshop," +
			" then press enter to continue")
		wg.Wait()
		d := ps.NewDeck(app.Normal)
		name := strings.Join(args, " ")
		if !strings.HasSuffix(name[:len(name)-1], "_") {
			name += "_1"
		}
		d.ApplyDataset(name)
		d.PNG(false)
	}
	fmt.Println(time.Since(start))
}