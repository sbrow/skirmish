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
	Template
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
	d := &DeckTemplate{Template: *New(mode, CardTemplate)}
	txt := d.Doc.MustExist("Text").(*ps.LayerSet)
	if txt == nil {
		log.Panic("LayerSet \"Text\" was not found!")
	}
	d.Banners = d.Doc.MustExist("Areas").(*ps.LayerSet).MustExist("TitleBackground").(*ps.LayerSet)
	if d.Banners == nil {
		log.Panic("LayerSet \"TitleBackground\" was not found!")
	}
	if ps.Mode == 2 {
		d.Dataset = d.ID.TextItem.Contents()
	}
	d.Cost = txt.MustExist("cost").(*ps.ArtLayer)
	if d.Cost == nil {
		log.Panic("ArtLayer \"cost\" was not found!")
	}
	d.Type = txt.MustExist("type").(*ps.ArtLayer)
	if d.Type == nil {
		log.Panic("ArtLayer \"type\" was not found!")
	}
	d.HeroLife = txt.MustExist("hero life").(*ps.ArtLayer)
	if d.HeroLife == nil {
		log.Panic("ArtLayer \"heroLife\" was not found!")
	}
	areas := d.Doc.MustExist("Areas").(*ps.LayerSet)
	if areas == nil {
		log.Panic("LayerSet \"Areas\" was not found!")
	}
	d.CostBG = areas.MustExist("CostBackground").(*ps.LayerSet).
		ArtLayer("cost_color")
	if d.CostBG == nil {
		log.Panic("cost_bg layer not found!")
	}
	ind := d.Doc.MustExist("Indicators").(*ps.LayerSet)
	if ind == nil {
		log.Panic("LayerSet \"Indicators\" was not found!")
	}
	bottom := areas.MustExist("Bottom").(*ps.LayerSet)
	if bottom == nil {
		log.Panic("LayerSet \"Bottom\" was not found!")
	}
	d.LBar = bottom.ArtLayer("L Bar")
	if d.LBar == nil {
		log.Panic("ArtLayer \"LBar\" was not found!")
	}
	d.HeroLifeBG = ind.ArtLayer("hero_life_background")
	if d.HeroLifeBG == nil {
		log.Panic("LayerSet \"hero_life_background\" was not found!")
	}
	d.DamageBG = ind.ArtLayer("damage_background")
	if d.DamageBG == nil {
		log.Panic("LayerSet \"damage_background\" was not found!")
	}
	d.LifeBG = ind.ArtLayer("life_background")
	if d.LifeBG == nil {
		log.Panic("LayerSet \"life_background\" was not found!")
	}
	d.RarityInd = ind.MustExist("Rarity").(*ps.LayerSet)
	if d.RarityInd == nil {
		log.Panic("Rarity layers not found!")
	}
	d.TypeInd = ind.MustExist("Type").(*ps.LayerSet)
	if d.TypeInd == nil {
		log.Panic("Rarity layers not found!")
	}
	return d
}

// ApplyDataset applies the given dataset,
// selects card data from the sql server,
// checks all its values against the active document,
// updates any fields that were changed,
// and then calls any necessary formatting functions.
func (d *DeckTemplate) ApplyDataset(id string) {
	name := strings.TrimRight(id, `_123`)

	// Skip if dataset already applied.
	if ps.Mode == ps.Fast && d.Dataset == id && d.Card != nil {
		if d.Card.Name() == name {
			return
		}
	}

	card, err := skirmish.Load(name)
	if err != nil {
		d.Doc.Dump()
		log.Panic(fmt.Sprintf("Card '%s' not found. Check your spelling.", name))
	}
	d.Card = card

	// TODO(sbrow): Skip SetLeader when appropriate.
	d.SetLeader(d.Card.Leader())

	// TODO(sbrow): run this as a go routine?
	d.Template.ApplyDataset(id, name)
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

// TODO(sbrow): Fix DeckTemplate.SetLeader skip
func (d *DeckTemplate) SetLeader(name string) {
	// if ps.Mode == 2 && name == d.Card.Leader() {
	// 	return
	// }
	banner, ind := d.Template.SetLeader(name)

	rarity := ps.Compare(banner, ind)
	barStroke := ps.Stroke{Size: 4, Color: banner}
	counterStroke := ps.Stroke{Size: 4, Color: ind}
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

	d.Cost.SetStroke(ps.Stroke{4, rarity}, ps.ColorWhite)
	d.HeroLife.SetStroke(barStroke, ps.ColorWhite)
}

// TODO(sbrow): (3) Make type font smaller when 2 or more supertypes.
func (d *DeckTemplate) FormatTextbox() {

	/*
		if len(d.Card.STypes()) > 1 {
			d.Type.TextItem.SetSize(9.0)
		} else {
			d.Type.TextItem.SetSize(10.0)
		}
	*/
	d.Template.FormatTextbox()
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

// PNG saves a copy the produced card image as a .png file in the appropriate
// subfolder of  "SK_OUT".
// If crop is true, the bleed area around the card is cropped out of the image
// before saving.
func (d *DeckTemplate) PNG(crop bool) {
	log.Println("Saving copy as PNG")
	if !crop {
		err := ps.SaveAs(filepath.Join(os.Getenv("SK_PS"), "Decks", d.Card.Leader(),
			d.ID.TextItem.Contents()))
		if err != nil {
			d.Doc.Dump()
			log.Panic(err)
		}
		return
	}
	err := ps.DoAction("DK", "Crop")
	if err != nil {
		d.Doc.Dump()
		log.Panic(err)
	}
	err = ps.SaveAs(filepath.Join(os.Getenv("SK_PS"), "Decks", d.Card.Leader(),
		d.ID.TextItem.Contents()))
	if err != nil {
		d.Doc.Dump()
		log.Panic(err)
	}
	err = ps.DoAction("DK", "Undo")
	if err != nil {
		d.Doc.Dump()
		log.Panic(err)
	}
}
