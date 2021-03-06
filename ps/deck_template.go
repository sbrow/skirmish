package ps

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	ps "github.com/sbrow/ps/v2"
	"github.com/sbrow/skirmish"
)

// CardTemplate holds the path to the Photoshop Template for Deck Cards.
var CardTemplate = filepath.Join(skirmish.Cfg.PS.Dir, skirmish.Cfg.PS.Deck)

// DeckTemplate is the Template for leader and partner hero cards.
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

// NewDeck returns a new DeckTemplate object, pulling values
// from the .psd file.
//
// NewDeck will open Photoshop and the corresponding Template .psd if
// it is not currently open.
func NewDeck(mode ps.ModeEnum) *DeckTemplate {
	d := &DeckTemplate{template: *newTemplate(mode, CardTemplate)}
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

// ApplyDataset performs the following actions:
// 		1. Applies the given dataset.
// 		2. Selects card data from the sql server.
//		3. Checks all its values against the active document.
// 		4. Updates any fields that were changed.
// 		5. Calls any necessary formatting functions.
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

	// TODO(sbrow): Skip SetLeader when appropriate. [Issue](https://github.com/sbrow/skirmish/issues/38)
	d.SetLeader(d.Card.Leader())

	// TODO(sbrow): run d.Template.ApplyDataset as a go routine? [Issue](https://github.com/sbrow/skirmish/issues/40)
	d.template.ApplyDataset(id)
	Error(d.Type.Refresh())

	// Update layer data
	Error(d.Cost.Refresh())
	Error(d.HeroLife.Refresh())
	Error(d.RarityInd.Refresh())
	Error(d.HeroLifeBG.Refresh())
	Error(d.DamageBG.Refresh())
	Error(d.LifeBG.Refresh())
	Error(d.TypeInd.Refresh())

	if d.Mode == UEMode {
		d.Cost.SetVisible(false)
		d.HeroLife.SetVisible(false)
	}
	// doc.LayerSet("Border").Refresh() // TODO(sbrow): Fix Border.Refresh() [Issue](https://github.com/sbrow/skirmish/issues/41)
	Error(d.FormatTitle())
	d.FormatTextbox()
}

// GetDoc returns the Document associated with this template. It implements the
// Template interface.
func (d *DeckTemplate) GetDoc() *ps.Document {
	return d.Doc
}

// SetLeader changes fill layers that contain a leader color,
// and sets them to the colors of the given leader.
//
// TODO(sbrow): Fix DeckTemplate.SetLeader skip [Issue](https://github.com/sbrow/skirmish/issues/39)
func (d *DeckTemplate) SetLeader(name string) {
	banner, ind, barStroke, err := d.template.SetLeader(name)
	if err != nil {
		Error(err)
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

// FormatTextbox sets appropriate tolerances for the text layers in the
// textbox, and hides or resizes elements that are too large.
//
// TODO(sbrow): (3) Make type font smaller when 2 or more supertypes. [Issue](https://github.com/sbrow/skirmish/issues/42)
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
			Error(lyr.SetVisible(true))
		} else {
			Error(lyr.SetVisible(false))
		}
	}
	if !found {
		d.Doc.Dump()
		return errors.New("given title is too long")
	}
	return nil
}
