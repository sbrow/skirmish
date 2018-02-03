package gen

import (
	"fmt"
	"github.com/sbrow/skirmish/deck"
	"github.com/sbrow/skirmish/ps"
	"io/ioutil"
	"log"
	"os"
	// "path"
	"strings"
)

const Template = "F:\\GitLab\\dreamkeepers-psd\\Template009.1.psd"

const Folder = "F:\\GitLab\\dreamkeepers-psd\\card_jsons"

func Main() {
	// trace.Start(os.Stdout)
	// defer trace.Stop()
	// log.SetPrefix("[starting] ")
	fmt.Println("Generating cards")
	loadData()
	log.SetPrefix("[photoshop] ")
	log.Println("Opening photoshop")
	ps.Start()
	log.Println("Opening Template")
	ps.Open(Template)
	ps.Wait("$ Import the current dataset file into photoshop, then press enter to continue")
	// ps.Js(path.Join(ps.Folder, ("syncCards.jsx")), ps.Folder)
	ps.Js("F:\\GitHub\\Code\\javascript\\src\\Photoshop\\Skirmish\\bin\\syncCards.jsx", "C:/")
	log.Println("Closing Template")
	ps.Close()
	log.Println("Quitting Photoshop")
	ps.Quit()
	// log.SetPrefix("\n[complete] ")
	fmt.Println("Cards successfully generated!")
}

func loadData() {
	log.SetPrefix("[dataset] ")
	log.Println(`Generating "dataset.csv"`)
	f, _ := os.Create("data.txt") // TODO: Fix.
	defer f.Close()

	dir, _ := ioutil.ReadDir(Folder) // TODO: Fix.
	labels := false
	for _, file := range dir {
		if isDeck(file.Name()) {
			log.Println("Generating", strings.TrimRight(file.Name(), ".json"))
			d := deck.New(Folder + "\\" + file.Name())
			if !labels {
				fmt.Fprint(f, d.Labels())
				labels = true
			}
			fmt.Fprintln(f, d.String())
		}
	}
	log.Println(`..."dataset.csv" generated!`)
}

func isDeck(filename string) bool {
	switch filename {
	case "Formatting.json":
		fallthrough
	case "Heroes.json":
		fallthrough
	case "old":
		fallthrough
	case "Wisp.json": // TODO: Hack.
		return false
	default:
		return true
	}
}
