package ps

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sbrow/ps"
	"github.com/sbrow/skirmish"
)

// HeroTemplate holds the path to the NonDeckCard Photoshop Template.
var HeroTemplate = filepath.Join(os.Getenv("SK_PS"), "Template009.1h.psd")

type NonDeckTemplate struct {
	template
	Plus     *ps.ArtLayer
	HaloInd  *ps.LayerSet
	HeroInd  *ps.ArtLayer
	Factions *ps.LayerSet
	LBar     *ps.LayerSet
	BtmBG    *ps.ArtLayer
}

func NewNonDeck(mode ps.ModeEnum) *NonDeckTemplate {
	log.SetPrefix("[ps.NewNonDeck] ")
	n := &NonDeckTemplate{template: *New(mode, HeroTemplate)}
	areas := n.Doc.MustExist("Areas").(*ps.LayerSet)
	if areas == nil {
		log.Panic("LayerSet \"Areas\" was not found!")
	}
	txt := n.Doc.MustExist("Text").(*ps.LayerSet)
	if txt == nil {
		log.Panic("LayerSet \"Text\" was not found!")
	}
	n.Plus = txt.MustExist("+").(*ps.ArtLayer)
	if n.Plus == nil {
		log.Panic("ArtLayer \"Text\" was not found!")
	}
	n.BtmBG = areas.ArtLayer("bottom_color")
	if n.BtmBG == nil {
		log.Panic("ArtLayer \"bottom_color\" was not found!")
	}
	n.LBar = areas.MustExist("LeaderBar").(*ps.LayerSet)
	if n.LBar == nil {
		log.Panic("LayerSet \"LeaderBar\" was not found!")
	}
	ind := n.Doc.MustExist("Indicators").(*ps.LayerSet)
	if ind == nil {
		log.Panic("LayerSet \"Indicators\" was not found!")
	}
	n.HeroInd = ind.ArtLayer("HeroIcon")
	if n.HeroInd == nil {
		log.Panic("ArtLayer \"HeroIcon\" was not found!")
	}
	n.HaloInd = ind.MustExist("Halo").(*ps.LayerSet)
	if n.HaloInd == nil {
		log.Panic("LayerSet \"Halo\" was not found!")
	}
	n.Factions = ind.MustExist("Faction").(*ps.LayerSet)
	if n.Factions == nil {
		log.Panic("LayerSet \"Faction\" was not found!")
	}
	return n
}

func (n *NonDeckTemplate) ApplyDataset(name string) {
	// Skip if dataset already applied.
	if ps.Mode == ps.Fast && n.Dataset == name && n.Card != nil {
		return
	}
	card, err := skirmish.Load(name)
	if err != nil {
		n.Doc.Dump()
		log.Panic(fmt.Sprintf("Card '%s' not found. Check your spelling.", name))
	}
	n.Card = card

	n.SetLeader(n.Card.Leader())
	id := name
	if strings.Contains(id, "(Halo)") {
		tmp := strings.Split(id, " ")
		id = tmp[0]
	}
	n.template.ApplyDataset(id)
}

func (n *NonDeckTemplate) GetDoc() *ps.Document {
	return n.Doc
}

func (n *NonDeckTemplate) SetLeader(name string) {
	banner, ind, barStroke, err := n.template.SetLeader(name)
	if err != nil {
		log.Fatal(err) // TODO(sbrow): Remove fatal err from NonDeckTemplate.SetLeader
	}
	for _, lyr := range n.LBar.ArtLayers() {
		if lyr.Name() != "LeaderBar" {
			lyr.SetColor(ind)
		} else {
			lyr.SetColor(banner)
		}
	}
	for _, lyr := range n.Factions.ArtLayers() {
		log.Println(lyr.Name(), n.Card.Faction())
		if lyr.Name() == n.Card.Faction() {
			lyr.SetStroke(barStroke, ind)
		}
	}
	halo := n.HaloInd.ArtLayers()
	halo[0].SetColor(ind)
	halo[1].SetColor(banner)
	n.HeroInd.SetColor(ind)
	n.BtmBG.SetColor(banner)
	n.Plus.SetStroke(barStroke, ps.ColorWhite)
	n.Resolve.SetStroke(barStroke, ps.ColorWhite)
	n.Life.SetStroke(ps.Stroke{Size: 0, Color: ind}, ps.ColorBlack)
}
