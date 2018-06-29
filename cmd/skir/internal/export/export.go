// TODO(sbrow): Add other file formats.
package export

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/cmd/skir/internal/base"
)

var formats = map[string]func(){
	"csv": func() {
		DataSet("nondeckcards", "cards.Leader IS NULL ORDER BY name ASC")
		DataSet("deckcards", "cards.Leader IS NOT NULL ORDER BY name ASC")
	},
	"ue": UEJSON,
}
var CmdExport = &base.Command{
	UsageLine: "export",
	Short:     "compile cards from database to csv",
	Long: `'Skir export' pulls cards from a database and compiles them into a csv
file to be used as a dataset in Photoshop.`,
	Run: func(*base.Command, []string) { formats["ue"]() },
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
	path := filepath.Join(skirmish.DataDir, fmt.Sprintf("%s.csv", name))
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
	if err := os.Mkdir("Unreal_JSONs", os.ModeDir); err != nil {
		base.Errorf(err.Error())
	}
	wg := &sync.WaitGroup{}
	wg.Add(len(skirmish.Leaders))
	for _, leader := range skirmish.Leaders {
		go func(name string) {
			defer wg.Done()
			cards, err := skirmish.LoadMany(fmt.Sprintf("cards.Leader = '%s'", name))
			f, err := os.Create(filepath.Join("./", "Unreal_JSONs", name+".json"))
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
		}(leader.Name)
		cards, err := skirmish.LoadMany(fmt.Sprintf("cards.Leader is NULL"))
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
