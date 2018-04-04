package build

import (
	"fmt"
	"github.com/sbrow/ps"
	sk "github.com/sbrow/skirmish"
	"log"
	"regexp"
)

func PSDs() {
	log.SetPrefix("[photoshop] ")
	log.Println("Opening photoshop")
	ps.Start()
	log.Println("Opening Template")
	ps.Open(sk.Template)
	ps.Wait("$ Import the current dataset file into photoshop, then press enter to continue")
	ps.DoJs("F:\\GitHub\\Code\\javascript\\src\\Photoshop\\Skirmish\\bin\\syncCards.jsx", "C:/")
	log.Println("Closing Template")
	ps.Close(2)
	log.Println("Closing Other open files")
	log.Println("Quitting Photoshop")
	ps.Quit(2)
}

func Regexp() {}

// TODO: Function to Replace '{1}'  with resolve crystals.
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

func isDeck(filename string) bool {
	switch filename {
	case "data.txt":
		fallthrough
	case "Formatting.json":
		fallthrough
	case "Heroes.json":
		fallthrough
	case "old":
		return false
	default:
		return true
	}
}
