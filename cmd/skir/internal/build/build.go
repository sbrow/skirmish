package build

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/cmd/skir/internal/base"
)

var CmdBuild = &base.Command{
	UsageLine: "build",
	Short:     "compile cards from database to csv",
	Long: `'Skir build' pulls cards from a database and compiles them into a csv
file to be used as a dataset in Photoshop.`,
}

func init() {
	CmdBuild.Run = buildRun
}

func buildRun(cmd *base.Command, args []string) {
	genDataSet("deckcards", "cards.Leader IS NOT NULL ORDER BY name ASC")
	genDataSet("nondeckcards", "cards.Leader IS NULL ORDER BY name ASC")
}

func genDataSet(name, query string) {
	log.SetPrefix(fmt.Sprintf("[%s] ", name))
	log.Println(`Generating Dataset`)
	cards, err := skirmish.LoadMany(query)
	if err != nil {
		log.Panic(err)
	}
	dat := [][]string{cards[0].Labels()}
	for _, card := range cards {
		dat = append(dat, card.CSV(false)...)
	}
	path := filepath.Join(skirmish.DataDir, fmt.Sprintf("%s.csv", name))
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(dat)
	log.Println(path, "generated!")
}
