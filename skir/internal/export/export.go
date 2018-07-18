// TODO(sbrow): Add xml format to skir/export.

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

type format struct {
	desc string
	f    func()
}

var formats = map[string]format{
	"csv": {
		desc: `csv formatted files to use as datasets in Photoshop.
		One file is generated for Deck Cards, and another is generated for Non-Deck Cards.`,
		f: func() {
			err := DataSet("nondeckcards", "cards.Leader IS NULL ORDER BY name ASC")
			if err != nil {
				base.Errorf("%s", err)
			}
			err = DataSet("deckcards", "cards.Leader IS NOT NULL ORDER BY name ASC")
			if err != nil {
				base.Errorf("%s", err)
			}
		},
	},
	"ue": {
		desc: `a collection of JSON files for importing into Unreal Engine.
		Deck cards are grouped by deck, Non-Deck Cards are grouped together.`,
		f: UEJSON,
	},
}
var CmdExport = &base.Command{
	UsageLine: "export [format]",
	Short:     "compile cards from the database to a specific format",
	Long: `'Skir export' pulls information for all cards from the database and compiles them into the given format.

The valid formats are:`,
	Run: func(*base.Command, []string) { formats["ue"].f() },
}

// DataSet returns the cards as a Photoshop dataset formatted csv file.
func DataSet(name, query string) error {
	log.SetPrefix(fmt.Sprintf("[%s] ", name))
	log.Println(`Generating dataset`)
	cards, err := skirmish.LoadMany(query)
	if err != nil {
		return err
	}
	dat := [][]string{cards[0].Labels()}
	for _, card := range cards {
		dat = append(dat, card.CSV(false)...)
	}
	path := filepath.Join(skirmish.Cfg.DB.Dir, fmt.Sprintf("%s.csv", name))
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
func UEJSON() {
	if err := os.Mkdir("Unreal_JSONs", 0777); err != nil {
		base.Errorf(err.Error())
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
			path := filepath.Join("Unreal_JSONs", name+".json")
			f, err := os.Create(path)
			if err != nil {
				base.Fatalf("%s %s", path, err.Error())
			}
			defer f.Close()
			for _, card := range cards {
				data, err := card.UEJSON(true)
				if err != nil {
					base.Errorf(err.Error())
					continue
				}
				f.Write(data)
			}
		}(leader.Name)
		cards, err := skirmish.LoadMany(fmt.Sprintf("cards.Leader is NULL"))
		if err != nil {
			base.Fatalf(err.Error())
		}
		f, err := os.Create(filepath.Join("./", "Unreal_JSONs", "Non_deck.json"))
		if err != nil {
			base.Fatalf(err.Error())
		}
		defer f.Close()
		for _, card := range cards {
			data, err := card.UEJSON(true)
			if err != nil {
				base.Errorf(err.Error())
				continue
			}
			f.Write(data)
		}
	}

	wg.Wait()
}
