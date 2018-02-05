// TODO: config file for non-programmers
package main

import (
	"flag"
	"fmt"
	"github.com/sbrow/skirmish/cmd/gen"
	"os"
	"regexp"
)

func init() {
}

func main() {
	flagSet := flag.NewFlagSet("", flag.ExitOnError)
	fast := flagSet.Bool("f", false, "fast mode: skips dataset generation.")
	flagSet.Parse(os.Args[2:])

	args := []string{}
	cmd := ""
	switch {
	case len(args) > 0:
		cmd = os.Args[1]
	case len(args) > 1:
		args = os.Args[2:]
	}
	// TODO: So kludgey
	if cmd == "gen" || cmd == "" {
		fmt.Println(*fast)
		if *fast {
			gen.GenPSDs(*fast)
		} else {
			gen.Main(args...)
		}
	}
}

func ReplaceText(text string) {
	// First, find the resolve text.
	reg, err := regexp.Compile("{[1-9]}")
	if err != nil {
		panic(err)
		return
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
