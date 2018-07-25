package ps

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	app "github.com/sbrow/ps/v2"
	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/ps"
	"github.com/sbrow/skirmish/skir/internal/base"
	"github.com/sbrow/skirmish/skir/internal/export"
)

var CmdPS = &base.Command{
	Run:       Run,
	UsageLine: "ps [deck] [all] [card name]",
	Short:     "fill out Photoshop templates",
	Long: `'skir ps' fills out a Photoshop template file with information from the database.


See the skirmish/ps package for instructions on setting up this tool.

To start generating files, invoke the tool with the cards you want. Available options are:
	- 'skir ps all' to generate all cards.
	- 'skir ps deck [leader name]' to generate all cards lead by the named character.
	- 'skir ps [card id]' to generate one card.

Running the tool will open Photoshop and the necessary template .psd,
after which it will pause and ask you to load a dataset file.
Dataset files are csv formatted files that correspond to fields in the Photoshop template.
The tool generates them for you and puts them in the dreamkeepers data folder as:
	- 'deckcards.csv' for deck cards.
	- 'nondeckcards.csv' for non-deckcards

To load a dataset file, open Photoshop and navigate to 'Image > Variables > Data Sets...'.
Make sure Encoding is set to "Automatic", and "Use First Column For Data Set Names" and
"Replace Existing Data Sets" are selected, then click 'Import' on the right side of the pop-up menu.
It will take a minute to load, but once it does,
hit 'OK' and then return to the terminal where you ran the tool and hit enter to continue.
After this, the program should not require are further user interaction.

The dataset file will only need to be reloaded when the Template is opened or the data is changed.

Deck cards will be output to "[dreamkeepers-psd]/Decks/[leader name]/[card id].png", Nondeck cards will
be output to "[dreamkeepers-psd]/Decks/Heroes/[card_id].png".

Photoshop is very slow, generating every card could take 15+ minutes, so be ready to wait.
`,
}

func Run(cmd *base.Command, args []string) {
	log.SetPrefix("[ps] ")
	log.Println("Opening Photoshop")
	if err := app.Init(); err != nil {
		base.Fatalf("%s", err)
	}
	var leaders []string
	var condition string

	var wg sync.WaitGroup
	//    if !*fast {
	wg.Add(1)
	go func() {
		defer wg.Done()
		export.DataSet("nondeckcards", "cards.Leader IS NULL")
		export.DataSet("deckcards", "cards.Leader IS NOT NULL")
	}()
	//    }

	switch args[0] {
	case "deck":
		if len(leaders) == 0 {
			leaders = []string{args[1]}
			condition = fmt.Sprintf("cards.leader='%s'", args[1])
		}
		// condition += " AND NOT EXISTS(SELECT name FROM completed WHERE name=cards.name)"
		cards, err := skirmish.LoadMany(fmt.Sprintf("%s", condition))
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("Cards", len(cards))
		d := ps.NewDeck(app.Normal)
		defer func() { d.Doc.Dump() }()
		defer app.Close(app.SaveChanges)
		app.Wait("$ Import the current dataset file into Photoshop," +
			" then press enter to continue")
		wg.Wait()
		for _, card := range cards {
			images, err := card.Images()
			if err != nil {
				log.Println("ERROR:", err)
				images = []string{"1"}
			}
			for i := range images {
				d.ApplyDataset(fmt.Sprintf("%s_%d", card.Name(), i+1))
				d.PNG(false)
			}
		}
	default:
		name := strings.Join(args, " ")
		card, err := skirmish.Load(name)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		switch card.(type) {
		case *skirmish.DeckCard:
			app.Open(ps.CardTemplate)
		case *skirmish.NonDeckCard:
			app.Open(ps.HeroTemplate)
		}
		app.Wait("$ Import the current dataset file into photoshop," +
			" then press enter to continue")
		if !strings.HasSuffix(name[:len(name)-1], "_") {
			name += "_1"
		}
		var t ps.Template
		switch card.(type) {
		case *skirmish.DeckCard:
			t = ps.NewDeck(app.Normal)
		case *skirmish.NonDeckCard:
			t = ps.NewNonDeck(app.Normal)
		}
		defer func() {
			ps.Errors.Report()
			t.GetDoc().Dump()
		}()
		wg.Wait()
		t.ApplyDataset(name)
		t.PNG(false)
	}
}
