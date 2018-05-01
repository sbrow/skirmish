package main

import (
	"flag"
	"fmt"
	app "github.com/sbrow/ps"
	sk "github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/ps"
	"github.com/sbrow/skirmish/sql"
	"github.com/sbrow/update"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	// "sync"
	"time"
)

var flagCpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func init() {
	update.Update()
}

func main() {
	start := time.Now()
	fast := flag.Bool("f", false, " fast mode- skip dataset generation.")
	flag.Parse()
	if *flagCpuprofile != "" {
		f, err := os.Create(*flagCpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	log.SetPrefix("[ps] ")
	log.Println("Opening Photoshop")
	app.Open(sk.Template)
	args := flag.Args()
	var leaders []string
	var condition string

	// var wg sync.WaitGroup
	if !*fast {
		go sql.GenData()
		// wg.Add(1)
		// go func() {
		// defer wg.Done()
		// sql.GenData()
		// }()
	}

	switch args[0] {
	case "crop":
	case "undo":
		// wg.Wait()
		err := app.DoAction("DK", strings.Title(args[0]))
		if err != nil {
			log.Panic(err)
		}
	case "all":
		leaders = make([]string, len(sk.Leaders))
		for i, ldr := range sk.Leaders {
			leaders[i] = ldr.Name
		}
		condition = "NOT cards.leader is NULL"
		fallthrough
	case "deck":
		if len(leaders) == 0 {
			leaders = []string{args[1]}
			condition = fmt.Sprintf("cards.leader='%s'", args[1])
		}
		condition += " AND NOT EXISTS(SELECT name FROM completed WHERE name=cards.name)"
		order := "cards.leader, cards.supertypes, cards.type, char_length(name) ASC"
		cards, err := sql.LoadMany(fmt.Sprintf("%s ORDER BY %s", condition, order))
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("Cards", len(cards))
		d := ps.NewDeck(app.Normal)
		defer d.Doc.Dump()
		defer app.Close(app.PSSaveChanges)
		app.Wait("$ Import the current dataset file into Photoshop," +
			" then press enter to continue")
		// wg.Wait()
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
	default:
		app.Wait("$ Import the current dataset file into photoshop," +
			" then press enter to continue")
		// wg.Wait()
		d := ps.NewDeck(app.Fast)
		defer d.Doc.Dump()
		name := strings.Join(args, " ")
		if !strings.HasSuffix(name[:len(name)-1], "_") {
			name += "_1"
		}
		d.ApplyDataset(name)
		d.PNG(false)
	}
	fmt.Println(time.Since(start))
}
