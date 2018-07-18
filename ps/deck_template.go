package ps

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sbrow/ps"
	"github.com/sbrow/skirmish"
)

// CardTemplate holds the path to the Photoshop Template for Deck Cards.
var CardTemplate = filepath.Join(os.Getenv("SK_PS"), "Template009.1.psd")

type DeckTemplate struct {
	template
	Banners    *ps.LayerSet
	Cost       *ps.ArtLayer
	Type       *ps.ArtLayer
	HeroLife   *ps.ArtLayer
	CostBG     *ps.ArtLayer
	TypeInd    *ps.LayerSet
	RarityInd  *ps.LayerSet
	HeroLifeBG *ps.ArtLayer
	DamageBG   *ps.ArtLayer
	LBar       *ps.ArtLayer
}

func NewDeck(mode ps.ModeEnum) *DeckTemplate {
	d := &DeckTemplate{template: *New(mode, CardTemplate)}
	txt := d.Doc.MustExist("Text").(*ps.LayerSet)
	d.Banners = d.Doc.MustExist("Areas").(*ps.LayerSet).MustExist("TitleBackground").(*ps.LayerSet)
	if ps.Mode == 2 {
		d.Dataset = d.ID.TextItem.Contents()
	}
	d.Cost = txt.MustExist("cost").(*ps.ArtLayer)
	d.Type = txt.MustExist("type").(*ps.ArtLayer)
	d.HeroLife = txt.MustExist("hero life").(*ps.ArtLayer)
	areas := d.Doc.MustExist("Areas").(*ps.LayerSet)
	d.CostBG = areas.MustExist("CostBackground").(*ps.LayerSet).MustExist("cost_color").(*ps.ArtLayer)
	ind := d.Doc.MustExist("Indicators").(*ps.LayerSet)
	bottom := areas.MustExist("Bottom").(*ps.LayerSet)
	d.LBar = bottom.MustExist("L Bar").(*ps.ArtLayer)
	d.HeroLifeBG = ind.MustExist("hero_life_background").(*ps.ArtLayer)
	d.DamageBG = ind.MustExist("damage_background").(*ps.ArtLayer)
	d.LifeBG = ind.MustExist("life_background").(*ps.ArtLayer)
	d.RarityInd = ind.MustExist("Rarity").(*ps.LayerSet)
	d.TypeInd = ind.MustExist("Type").(*ps.LayerSet)
	return d
}

// ApplyDataset applies the given dataset,
// selects card data from the sql server,
// checks all its values against the active document,
// updates any fields that were changed,
// and then calls any necessary formatting functions.
func (d *DeckTemplate) ApplyDataset(id string) {

	// Skip if dataset already applied.
	if ps.Mode == ps.Fast && d.Dataset == id && d.Card != nil {
		if d.Card.Name() == id {
			return
		}
	}
	name := strings.TrimRight(id, `_123`)

	card, err := skirmish.Load(name)
	if err != nil {
		d.Doc.Dump()
		log.Panic(fmt.Sprintf("Card '%s' not found. Check your spelling.", name))
	}
	d.Card = card

	// TODO(sbrow): Skip SetLeader when appropriate.
	d.SetLeader(d.Card.Leader())

	// TODO(sbrow): run d.Template.ApplyDataset as a go routine?
	d.template.ApplyDataset(id)
	d.Type.Refresh()

	// Update layer data
	d.Cost.Refresh()
	d.HeroLife.Refresh()
	d.RarityInd.Refresh()
	d.HeroLifeBG.Refresh()
	d.DamageBG.Refresh()
	d.LifeBG.Refresh()
	d.TypeInd.Refresh()

	// doc.LayerSet("Border").Refresh() // TODO(sbrow): Fix Border.Refresh()
	d.FormatTitle()
	d.FormatTextbox()
}

func (d *DeckTemplate) GetDoc() *ps.Document {
	return d.Doc
}

// TODO(sbrow): Fix DeckTemplate.SetLeader skip
func (d *DeckTemplate) SetLeader(name string) {
	banner, ind, barStroke, err := d.template.SetLeader(name)
	if err != nil {
		log.Panic(err) // TODO(sbrow): Remove panic in DeckTemplate.SetLeader
	}
	counterStroke := ps.Stroke{Size: 4, Color: ind}
	rarity := ps.Compare(banner, ind)
	rarities := d.RarityInd.ArtLayers()

	if d.Card != nil &&
		strings.Contains(strings.Join(d.Card.STypes(), ","), "Channeled") {
		d.CostBG.SetColor(rarity)
	} else {
		d.CostBG.SetColor(ps.ColorGray)
	}
	for _, lyr := range d.TypeInd.ArtLayers() {
		lyr.SetColor(ind)
	}
	// Use indices instead of range because the bottom layer is the
	// rarity_background and we want it to stay black.
	for i := 0; i < 3; i++ {
		rarities[i].SetColor(rarity)
	}
	d.Resolve.SetStroke(counterStroke, ps.ColorWhite)
	d.Life.SetStroke(barStroke, ps.ColorWhite)

	d.LBar.SetColor(banner)
	d.HeroLifeBG.SetColor(ind)
	d.DamageBG.SetColor(ind)
	d.LifeBG.SetColor(ind)

	d.Cost.SetStroke(ps.Stroke{Size: 4, Color: rarity}, ps.ColorWhite)
	d.HeroLife.SetStroke(barStroke, ps.ColorWhite)
}

// TODO(sbrow): (3) Make type font smaller when 2 or more supertypes.
func (d *DeckTemplate) FormatTextbox() {
	if len(d.Card.STypes()) > 1 {
		d.Type.TextItem.SetSize(9.0)
	} else {
		d.Type.TextItem.SetSize(10.0)
	}
	d.template.FormatTextbox()
}

// FormatTitle finds the correct length background for the card's title, makes
// it visible, and hides the rest. Returns an error if the title was longer than
// the longest background.
func (d *DeckTemplate) FormatTitle() error {
	tol := Tolerances["title"]
	found := false
	for _, lyr := range d.Banners.ArtLayers() {
		if !found && d.Name.Bounds()[1][0]+tol <= lyr.Bounds()[1][0] {
			found = true
			lyr.SetVisible(true)
		} else {
			lyr.SetVisible(false)
		}
	}
	if !found {
		d.Doc.Dump()
		return errors.New("Title too long.")
	}
	return nil
}
