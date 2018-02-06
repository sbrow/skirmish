package gen

import (
	"fmt"
	"github.com/sbrow/ps"
	"github.com/sbrow/skirmish/deck"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

const Template = "F:\\GitLab\\dreamkeepers-psd\\Template009.1.psd"

var DataDir string
var ImgDir string

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
				srcDir = path.Join(os.Getenv("HOME"), "Downloads")
			}
			os.Setenv(env, srcDir)
			log.Printf("Environment variable \"%s\" not found. "+
				"Will attempt to run with \"%[1]s=%s\"\n", env, os.Getenv(env))
		}
		DataDir = os.Getenv("SK_SRC")
		ImgDir = os.Getenv("SK_IMG")
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

func Dataset() {
	log.SetPrefix("[dataset] ")
	log.Println(`Generating "dataset.csv"`)
	f, err := os.Create("data.txt")
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
			d := deck.New(path.Join(DataDir, file.Name()))
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
	ps.Js("F:\\GitHub\\Code\\javascript\\src\\Photoshop\\Skirmish\\bin\\syncCards.jsx", "C:/")
	log.Println("Closing Template")
	ps.Close()
	log.Println("Closing Other open files")
	log.Println("Quitting Photoshop")
	ps.Quit(2)
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
