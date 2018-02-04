package gen

import (
	"fmt"
	"github.com/sbrow/skirmish/deck"
	"github.com/sbrow/skirmish/ps"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

const Template = "F:\\GitLab\\dreamkeepers-psd\\Template009.1.psd"

var Folder = os.Getenv("SK_SRC")

func Main() {
	// trace.Start(os.Stdout)
	// defer trace.Stop()
	// log.SetPrefix("[starting] ")
	fmt.Println("Generating cards")
	genDataset()
	genPSDs()
	log.SetPrefix("[photoshop] ")

	// log.SetPrefix("\n[complete] ")
	fmt.Println("Cards successfully generated!")
}

func genDataset() {
	log.SetPrefix("[dataset] ")
	log.Println(`Generating "dataset.csv"`)
	f, _ := os.Create("data.txt") // TODO: Fix.
	defer f.Close()

	dir, err := ioutil.ReadDir(Folder) // TODO: Fix.
	if err != nil {
		panic(err)
	}
	labels := false
	for _, file := range dir {
		if isDeck(file.Name()) {
			d := deck.New(path.Join(Folder, file.Name()))
			// log.Println("Generating", strings.TrimRight(file.Name(), ".json"))
			if !labels {
				fmt.Fprint(f, d.Labels())
				labels = true
			}
			fmt.Fprintln(f, d.String())
		}
	}
	log.Println(`..."dataset.csv" generated!`)
}

func genPSDs() {
	defer func() {
		if r := recover(); strings.Contains(r.(string), "windows") {
			log.Printf("skipping genPSDs(): %s (%s)",
				"Not supported on this OS.",
				runtime.GOOS)
		}
	}()
	log.Println("Opening photoshop")
	if runtime.GOOS != "windows" {
		log.Panic("Non-windows OS detected. skir only interfaces with " +
			"64 bit Windows.")
	}
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
