package main

import (
	"fmt"
	"github.com/sbrow/debug"
	"github.com/sbrow/skirmish/deck"
	"io"
	"io/ioutil"
	"os"
	"runtime/trace"
)

const FOLDER = "F:\\GitLab\\dreamkeepers-psd\\card_jsons"

func main() {
	trace.Start(os.Stdout)
	defer trace.Stop()
	f, err := os.Create("data.txt")
	debug.Check(err)
	defer f.Close()

	fmt.Fprint(f, deck.Labels())
	dir, err := ioutil.ReadDir(FOLDER)
	debug.Check(err)

	go func(w io.Writer) {
		for _, file := range dir {
			if isDeck(file.Name()) {
				d := deck.New(FOLDER + "\\" + file.Name())
				fmt.Fprintln(w, d.String())
			}
		}
	}(f)
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
