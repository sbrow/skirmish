package main

import (
	"fmt"
	"github.com/sbrow/skirmish/deck"
	// "io"
	"io/ioutil"
	"os"
)

const Folder = "F:\\GitLab\\dreamkeepers-psd\\card_jsons"

func main() {
	// trace.Start(os.Stdout)
	// defer trace.Stop()
	f, _ := os.Create("data.txt") // TODO: Fix.
	defer f.Close()

	fmt.Fprint(f, deck.Labels())
	dir, _ := ioutil.ReadDir(Folder) // TODO: Fix.

	for _, file := range dir {
		if isDeck(file.Name()) {
			d := deck.New(Folder + "\\" + file.Name())
			fmt.Fprintln(f, d.String())
		}
	}
}

func isDeck(filename string) bool {
	switch filename {
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
