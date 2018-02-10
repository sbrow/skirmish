package build

import (
	"fmt"
	"github.com/sbrow/ps"
	"github.com/sbrow/skirmish/deck"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func init() {
	defer log.SetPrefix("")
	log.SetPrefix("[init] ")
	log.Print("Initializing")

	// Handle panics
	defer func() {
		env := ""
		if r := recover(); r != nil {
			switch r {
			case strings.Contains(r.(string), "SK_IMG"):
				env = "SK_IMG"
			case strings.Contains(r.(string), "SK_SRC"):
				env = "SK_SRC"
			}
			srcDir := ""
			switch runtime.GOOS {
			case "windows":
				srcDir = filepath.Join(os.Getenv("HOMEPATH"), "Downloads")
			default:
				srcDir = filepath.Join(os.Getenv("HOME"), "Downloads")
			}
			os.Setenv(env, srcDir)
			log.Printf("Environment variable \"%s\" not found. "+
				"Will attempt to run with \"%[1]s=%s\"\n", env, os.Getenv(env))
		}
		DataDir = os.Getenv("SK_SRC")
		ImageDir = os.Getenv("SK_IMG")
	}()

	envVars := []string{"SK_SRC", "SK_IMG"}
	for _, val := range envVars {
		_, ok := os.LookupEnv(val)
		if !ok {
			log.SetPrefix("[ERROR] ")
			log.Panicf("Environment variable \"%s\" does not exist - please create it!",
				val)
			log.SetPrefix("")
		}

	}
}

func Data() {
	log.SetPrefix("[dataset] ")
	log.Println(`Generating "dataset.csv"`)
	// f, err := os.Create(filepath.Join(os.Getenv("SK_SRC"), "data.txt"))
	f, err := os.Create(filepath.Join("data.txt"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dir, err := ioutil.ReadDir(DataDir)
	if err != nil {
		panic(fmt.Sprintf("%s (%s)", err, DataDir))
	}
	labels := false
	for _, file := range dir {
		if isDeck(file.Name()) {
			d := deck.New(DataDir)
			log.Println("Generating", strings.TrimRight(file.Name(), ".json"))
			if !labels {
				fmt.Fprint(f, d.Labels())
				labels = true
			}
			fmt.Fprintln(f, d.String())
		}
	}
	log.Println("\"dataset.csv\" generated!")
}

func PSDs() {
	log.SetPrefix("[photoshop] ")
	log.Println("Opening photoshop")
	ps.Start()
	log.Println("Opening Template")
	ps.Open(Template)
	ps.Wait("$ Import the current dataset file into photoshop, then press enter to continue")
	ps.DoJs("F:\\GitHub\\Code\\javascript\\src\\Photoshop\\Skirmish\\bin\\syncCards.jsx", "C:/")
	log.Println("Closing Template")
	ps.Close()
	log.Println("Closing Other open files")
	log.Println("Quitting Photoshop")
	ps.Quit(2)
}

func Regexp() {}

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
