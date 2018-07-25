// TODO(sbrow): Add xml format to skir/export. [Issue: https://github.com/sbrow/skirmish/issues/30] [Issue](https://github.com/sbrow/skirmish/issues/50)

package export

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/skir/internal/base"
)

type format struct {
	desc string
	f    func() error
}

var formats = map[string]format{
	"csv": {
		desc: `csv formatted files to use as datasets in Photoshop.
		One file is generated for Deck Cards, and another is generated for Non-Deck Cards.
		The files are generated in the top level of the "dreamkeepers-psd" repository.`,
		f: func() error {
			err := DataSet("nondeckcards", "cards.leader IS NULL")
			if err != nil {
				base.Errorf("%s", err)
			}
			err = DataSet("deckcards", "cards.leader IS NOT NULL")
			if err != nil {
				base.Errorf("%s", err)
			}
			return err
		},
	},
	"ue": {
		desc: `a collection of JSON files for importing into Unreal Engine.
		Deck cards are grouped by deck, Non-Deck Cards are grouped together.
		The files can be found in the "skirmish/data/Unreal_JSONs/" folder.`,
		f: UEJSON,
	},
}

var CmdExport = &base.Command{
	UsageLine: "export [format]",
	Short:     "compile cards from the database to a specific format",
	Long: `'Skir export' pulls information for all cards from the database and compiles them into the given format.

The valid formats are:`,
	Run: Run,
}

func init() {
	fmts := make([]string, len(formats))
	i := 0
	for name := range formats {
		fmts[i] = name
		i++
	}
	sort.Strings(fmts)
	for _, name := range fmts {
		desc := formats[name].desc
		CmdExport.Long += fmt.Sprintf("\n\t%s\t%s", name, desc)
	}
}
func Run(cmd *base.Command, args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stdout, cmd.UsageLine)
		fmt.Fprintln(os.Stdout, cmd.Short)
		fmt.Fprintln(os.Stdout, cmd.Long)
		return
	}
	format, ok := formats[args[0]]
	log.SetOutput(os.Stdout)
	if !ok {
		base.Errorf("format %s was not found", args[0])
		return
	}
	if err := format.f(); err != nil {
		base.Errorf("%s", err)
	}
}

// DataSet returns the cards as a Photoshop dataset formatted csv file.
func DataSet(name, query string) error {
	log.SetPrefix(fmt.Sprintf("[%s] ", name))
	log.SetOutput(os.Stdout)
	log.Println(`Generating dataset`)
	cards, err := skirmish.LoadMany(query)
	if err != nil {
		return err
	}
	dat := [][]string{cards[0].Labels()}
	for _, card := range cards {
		dat = append(dat, card.CSV(false)...)
	}
	path := filepath.Join(skirmish.Cfg.PS.Dir, fmt.Sprintf("%s.csv", name))
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(dat)
	log.Println(path, "generated!")
	return nil
}

// UEJSON generates JSON files for import into Unreal Engine.
func UEJSON() error {
	log.SetOutput(os.Stdout)
	log.Println("Generating JSON files for Unreal Engine...")
	pkg := filepath.Join("github.com", "sbrow", "skirmish")
	path := filepath.Join(os.Getenv("GOPATH"), "src", pkg, "data", "Unreal_JSONs")
	if err := os.MkdirAll(path, 0700); err != nil {
		log.Println(err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(len(skirmish.Leaders))
	for _, leader := range skirmish.Leaders {
		go func(name string) {
			defer wg.Done()
			cards, err := skirmish.LoadMany(fmt.Sprintf("cards.Leader = '%s'", name))
			if err != nil {
				base.Fatalf(err.Error())
			}
			path := filepath.Join(path, name+".json")
			f, err := os.Create(path)
			if err != nil {
				base.Fatalf(err.Error())
			}
			defer f.Close()
			f.Write([]byte("[\n"))
			for i, card := range cards {
				data, err := card.UEJSON(true)
				if err != nil {
					base.Errorf(card.Name(), err.Error())
					continue
				}
				f.Write(data)
				if i+1 != len(cards) {
					f.Write([]byte{',', '\n'})
				}
			}
			f.Write([]byte("]"))
		}(leader.Name)
		cards, err := skirmish.LoadMany(fmt.Sprintf("cards.Leader is NULL"))
		if err != nil {
			return err
		}
		f, err := os.Create(filepath.Join(path, "Non_deck.json"))
		if err != nil {
			return err
		}
		defer f.Close()
		f.Write([]byte("[\n"))
		for i, card := range cards {
			data, err := card.UEJSON(true)
			if err != nil {
				base.Errorf(err.Error())
				continue
			}
			f.Write(data)
			if i+1 != len(cards) {
				f.Write([]byte{',', '\n'})
			}
		}
		f.Write([]byte("]"))
	}
	wg.Wait()
	return nil
}
