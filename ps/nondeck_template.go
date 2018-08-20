package ps

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	ps "github.com/sbrow/ps/v2"
	"github.com/sbrow/skirmish"
)

// HeroTemplate holds the path to the NonDeckCard Photoshop Template.
var HeroTemplate = filepath.Join(skirmish.Cfg.PS.Dir, skirmish.Cfg.PS.NonDeck)

// NonDeckTemplate is the Template to use for leader and partner hero cards.
type NonDeckTemplate struct {
	template
	Plus     *ps.ArtLayer
	HaloInd  *ps.LayerSet
	HeroInd  *ps.ArtLayer
	Factions *ps.LayerSet
	LBar     *ps.LayerSet
	BtmBG    *ps.ArtLayer
}

// NewNonDeck returns a new DeckTemplate object, pulling values
// from the .psd file.
//
// NewNonDeck will open Photoshop and the corresponding Template .psd if
// it is not currently open.
func NewNonDeck(mode ps.ModeEnum) *NonDeckTemplate {
	log.SetPrefix("[ps.NewNonDeck] ")
	n := &NonDeckTemplate{template: *new(mode, HeroTemplate)}
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

// ApplyDataset performs the following actions:
// 		1. Applies the given dataset.
// 		2. Selects card data from the sql server.
//		3. Checks all its values against the active document.
// 		4. Updates any fields that were changed.
// 		5. Calls any necessary formatting functions.
func (n *NonDeckTemplate) ApplyDataset(id string) {
	// Skip if dataset already applied.
	if ps.Mode == ps.Fast && n.Dataset == id && n.Card != nil {
		return
	}
	name := id
	if strings.Contains(name, " (Halo)") {
		name = strings.TrimSuffix(name, " (Halo)")
	}
	card, err := skirmish.Load(name)
	if err != nil {
		n.Doc.Dump()
		log.Panic(fmt.Sprintf("Card '%s' not found. Check your spelling.", id))
	}
	n.Card = card

	n.template.ApplyDataset(id)
	n.SetLeader(n.Card.Name())

	if n.Mode == UEMode {
		n.Plus.SetVisible(false)
	}
	n.template.FormatTextbox()
}

// GetDoc returns the Document associated with this Template.
func (n *NonDeckTemplate) GetDoc() *ps.Document {
	return n.Doc
}

// SetLeader changes fill layers that contain a leader color,
// and sets them to the colors of the given leader.
func (n *NonDeckTemplate) SetLeader(name string) {
	banner, ind, barStroke, err := n.template.SetLeader(name)
	if err != nil {
		Error(err)
		return
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
	n.SpeedBG.SetColor(banner)
	n.BtmBG.SetColor(banner)
	n.Plus.SetStroke(barStroke, ps.ColorWhite)
	n.Resolve.SetStroke(barStroke, ps.ColorWhite)
	n.Life.SetStroke(ps.Stroke{Size: 0, Color: ind}, ps.ColorBlack)
}
