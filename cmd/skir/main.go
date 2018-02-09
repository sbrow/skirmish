// TODO: config file for non-programmers
// TODO: command - "card" display info for a card
// TODO: separate commands for each ps operation.
package main

import (
	// "flag"
	"fmt"
	"github.com/sbrow/ps"
	"github.com/sbrow/skirmish/gen"
	"github.com/sbrow/skirmish/photoshop"
	"log"
	"os"
	"regexp"
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
			err := ps.DoAction("DK", strings.Title(args[0]))
			if err != nil {
				panic(err)
			}
		case "save":
			photoshop.Save(true, args...)
		}
	case "gen":
		fallthrough
	case "":
		log.SetPrefix("[main] ")
		// if !*fast {
		log.Println("Generating cards")
		gen.Dataset()
		// }
		log.SetPrefix("[photoshop] ")
		gen.PSDs()
		log.Println("Cards successfully generated!")
	}
}

func ReplaceText(text string) {
	// First, find the resolve text.
	reg, err := regexp.Compile("{[1-9]}")
	if err != nil {
		panic(err)
	}
	temp := reg.FindStringIndex(text)
	resolve := text[temp[0]:temp[1]]

	// Prevents compiler errors. Remove eventually.
	fmt.Println(resolve)

	// Next, find the lower bounds of the text
	// +	Get the BR x value by stripping away all other lines,
	// 		and all text to the right of the symbol.
	// +	Get the BR y value by stripping away all lines/text after it.
	// layer.textItem.contents = text[:temp[1]]
	// x1, y1, x2, y2 = layer.textItem.bounds
	//
	// Place the circle there
	// resolveCircle = placeFile(x2, y2, filename, "bottom right")
	//
	// Color it
	// colorlayer(resolveCircle, color)
	//
	// Place and color the number.
	//
	// Scrub away the old text, add space as necessary.
	//
	return
}
