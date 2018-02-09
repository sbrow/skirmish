// TODO: config file for non-programmers
// TODO: command - "card" display info for a card
// TODO: separate commands for each ps operation.
package main

import (
	// "flag"
	// "fmt"
	app "github.com/sbrow/ps"
	"github.com/sbrow/skirmish/build"
	"github.com/sbrow/skirmish/ps"
	"log"
	"os"
	"strings"
)

func main() {
	// flagSet := flag.NewFlagSet("", flag.ExitOnError)
	// fast := flagSet.Bool("f", false, "fast mode: skips dataset generation.")
	// flagSet.Parse(os.Args[2:])

	args := []string{}
	cmd := ""
	switch {
	case len(os.Args) > 2:
		args = os.Args[2:]
		fallthrough
	case len(os.Args) > 1:
		cmd = os.Args[1]
	}
	switch cmd {
	case "ps":
		switch args[0] {
		case "crop":
		case "undo":
			err := app.DoAction("DK", strings.Title(args[0]))
			if err != nil {
				panic(err)
			}
		case "save":
			ps.Save(true, args...)
		}
	case "gen":
		fallthrough
	case "":
		log.SetPrefix("[main] ")
		// if !*fast {
		log.Println("Generating cards")
		build.Data()
		// }
		log.SetPrefix("[photoshop] ")
		build.PSDs()
		log.Println("Cards successfully generated!")
	}
}
