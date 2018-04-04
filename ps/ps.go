package ps

import (
	"errors"
	"github.com/sbrow/ps"
	sk "github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Template struct {
	Doc       *ps.Document
	DataSet   string
	Resolve   *ps.ArtLayer
	Banners   *ps.LayerSet
	Speed     *ps.ArtLayer
	Life      *ps.ArtLayer
	Damage    *ps.ArtLayer
	Short     *ps.ArtLayer
	Long      *ps.ArtLayer
	Flavor    *ps.ArtLayer
	ShortBG   *ps.ArtLayer
	LBar      *ps.ArtLayer
	ResolveBG *ps.ArtLayer
	DeckInd   *ps.LayerSet
}

func New(mode ps.ModeEnum, file string) *Template {
	ps.Open(file)
	ps.Mode = mode
	t := &Template{}
	log.Printf("Creating new template with mode %d", mode)
	doc, err := ps.ActiveDocument()
	if err != nil {
		log.Fatal(err)
	}
	t.Doc = doc
	txt := doc.LayerSet("Text")
	if txt == nil {
		log.Panic("LayerSet \"Text\" was not found!")
	}
	t.Banners = doc.LayerSet("Areas").LayerSet("TitleBackground")
	if t.Banners == nil {
		log.Panic("LayerSet \"TitleBackground\" was not found!")
	}
	t.Resolve = txt.ArtLayer("resolve")
	if t.Resolve == nil {
		log.Panic("ArtLayer \"resolve\" was not found!")
	}
	t.Speed = txt.ArtLayer("speed")
	if t.Speed == nil {
		log.Panic("ArtLayer \"speed\" was not found!")
	}
	t.Life = txt.ArtLayer("life")
	if t.Life == nil {
		log.Panic("ArtLayer \"life\" was not found!")
	}
	t.Damage = txt.ArtLayer("damage")
	if t.Damage == nil {
		log.Panic("ArtLayer \"damage\" was not found!")
	}
	t.Short = txt.ArtLayer("short")
	if t.Short == nil {
		log.Panic("ArtLayer \"short\" was not found!")
	}
	t.Long = txt.ArtLayer("long")
	if t.Long == nil {
		log.Panic("ArtLayer \"long\" was not found!")
	}
	t.Flavor = txt.ArtLayer("flavor")
	if t.Flavor == nil {
		log.Panic("ArtLayer \"flav\" was not found!")
	}
	areas := doc.LayerSet("Areas")
	if areas == nil {
		log.Panic("LayerSet \"Areas\" was not found!")
	}
	bottom := areas.LayerSet("Bottom")
	if bottom == nil {
		log.Panic("LayerSet \"Bottom\" was not found!")
	}
	t.ShortBG = bottom.ArtLayer("short_text_box")
	if t.ShortBG == nil {
		log.Panic("ArtLayer \"short_bg\" was not found!")
	}
	t.LBar = bottom.ArtLayer("L Bar")
	if t.LBar == nil {
		log.Panic("ArtLayer \"L Bar\" not found!")
	}
	t.ResolveBG = areas.LayerSet("ResolveBackground").
		ArtLayer("resolve_color")
	if t.ResolveBG == nil {
		log.Panic("resolve_bg layer not found!")
	}
	ind := doc.LayerSet("Indicators")
	if ind == nil {
		log.Panic("LayerSet \"Indicators\" was not found!")
	}
	t.DeckInd = ind.LayerSet("Deck")
	if t.DeckInd == nil {
		log.Panic("LayerSet \"Deck\" was not found!")
	}
	return t
}

type DeckTemplate struct {
	Template
	Name       *ps.ArtLayer
	ID         *ps.ArtLayer
	Cost       *ps.ArtLayer
	Type       *ps.ArtLayer
	HeroLife   *ps.ArtLayer
	CostBG     *ps.ArtLayer
	TypeInd    *ps.LayerSet
	RarityInd  *ps.LayerSet
	HeroLifeBG *ps.ArtLayer
	DamageBG   *ps.ArtLayer
	LifeBG     *ps.ArtLayer
	SpeedBG    *ps.ArtLayer
}

func NewDeck(mode ps.ModeEnum) *DeckTemplate {
	d := &DeckTemplate{Template: *New(mode, sk.Template)}
	txt := d.Doc.LayerSet("Text")
	if txt == nil {
		log.Panic("LayerSet \"Text\" was not found!")
	}
	d.Name = txt.ArtLayer("name")
	if d.Name == nil {
		log.Panic("ArtLayer \"name\" was not found!")
	}
	d.ID = txt.ArtLayer("id")
	if d.ID == nil {
		log.Panic("ArtLayer \"id\" was not found!")
	}
	if ps.Mode == 2 {
		d.DataSet = *d.ID.Text
	}
	d.Cost = txt.ArtLayer("cost")
	if d.Cost == nil {
		log.Panic("ArtLayer \"cost\" was not found!")
	}
	d.Type = txt.ArtLayer("type")
	if d.Type == nil {
		log.Panic("ArtLayer \"type\" was not found!")
	}
	d.HeroLife = txt.ArtLayer("hero life")
	if d.HeroLife == nil {
		log.Panic("ArtLayer \"heroLife\" was not found!")
	}
	areas := d.Doc.LayerSet("Areas")
	if areas == nil {
		log.Panic("LayerSet \"Areas\" was not found!")
	}
	d.CostBG = areas.LayerSet("CostBackground").
		ArtLayer("cost_color")
	if d.CostBG == nil {
		log.Panic("cost_bg layer not found!")
	}
	ind := d.Doc.LayerSet("Indicators")
	if ind == nil {
		log.Panic("LayerSet \"Indicators\" was not found!")
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
	d.SpeedBG = ind.ArtLayer("speed_background")
	if d.SpeedBG == nil {
		log.Panic("LayerSet \"speed_background\" was not found!")
	}
	d.RarityInd = ind.LayerSet("Rarity")
	if d.RarityInd == nil {
		log.Panic("Rarity layers not found!")
	}
	d.TypeInd = ind.LayerSet("Type")
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
	if d.DataSet == id {
		return
	}
	defer d.Doc.Dump()
	log.Printf("Applying dataset %s\n", id)
	card, err := sql.Load(strings.TrimRight(id, `_123`))
	if err != nil {
		log.Println(card)
		log.Panic(err)
	}
	d.SetLeader(card.Leader())
	ps.ApplyDataset(id)

	// Update layer data
	d.Name.Refresh()
	d.ID.Refresh()
	d.Cost.Refresh()
	d.Resolve.Refresh()
	d.Type.Refresh()
	d.Speed.Refresh()
	d.Life.Refresh()
	d.Damage.Refresh()
	d.Speed.Refresh()
	d.Short.Refresh()
	d.Long.Refresh()
	d.Flavor.Refresh()
	d.HeroLife.Refresh()
	d.RarityInd.Refresh()
	d.HeroLifeBG.Refresh()
	d.DamageBG.Refresh()
	d.LifeBG.Refresh()
	d.SpeedBG.Refresh()
	d.DeckInd.Refresh()
	d.TypeInd.Refresh()

	// TODO: Fix Border.Refresh()
	// doc.LayerSet("Border").Refresh()
	d.ResolveBG.Refresh()
	for _, lyr := range d.Doc.ArtLayers() {
		if lyr.Name() == "card_image" {
			lyr.Refresh()
		}
	}
	d.Format()
}

func (d *DeckTemplate) SetLeader(name string) {
	var banner ps.Hex
	var ind ps.Hex

	for _, ldr := range sk.Leaders {
		if ldr.Name == name {
			banner = ldr.Banner
			ind = ldr.Indicator
		}
	}
	if banner == nil || ind == nil {
		log.Panic("Leader not found!")
	}
	rarity := ps.Compare(banner, ind)
	barStroke := ps.Stroke{Size: 4, Color: banner}
	counterStroke := ps.Stroke{Size: 4, Color: ind}
	rarities := d.RarityInd.ArtLayers()

	d.ResolveBG.SetColor(ps.Colors["Gray"])
	d.CostBG.SetColor(banner)
	for _, lyr := range d.TypeInd.ArtLayers() {
		lyr.SetColor(ind)
	}
	// Use indices instead of range because the bottom layer is the
	// rarity_background and we want it to stay black.
	for i := 0; i < 3; i++ {
		rarities[i].SetColor(rarity)
	}
	d.LBar.SetColor(banner)
	d.HeroLifeBG.SetColor(ind)
	d.DamageBG.SetColor(ind)
	d.SpeedBG.SetColor(ind)
	d.LifeBG.SetColor(ind)

	for _, lyr := range d.DeckInd.ArtLayers() {
		if lyr.Name() == name {
			lyr.SetVisible(true)
		} else {
			lyr.SetVisible(false)
		}
	}
	d.Cost.SetStroke(counterStroke, ps.Colors["White"])
	d.Resolve.SetStroke(counterStroke, ps.Colors["White"])

	d.Speed.SetStroke(barStroke, ps.Colors["White"])
	d.Damage.SetStroke(barStroke, ps.Colors["White"])
	d.Life.SetStroke(barStroke, ps.Colors["White"])
	d.HeroLife.SetStroke(barStroke, ps.Colors["White"])
}

// Format rearranges, hides, and colors layers as appropriate.
func (d *DeckTemplate) Format() {
	err := d.FormatTitle()
	if err != nil {
		panic(err)
	}
	d.FormatTextbox()
}

// FormatTitle finds the correct length background for the card's title, makes
// it visible, and hides the rest. Returns an error if the title was longer than
// the longest background.
func (d *DeckTemplate) FormatTitle() error {
	tol := sk.Tolerances["title"]
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
		return errors.New("Title too long.")
	}
	return nil
}

// FormatTextbox arranges text and background layers inside the textbox, hiding
// layers as necessary.
func (d *DeckTemplate) FormatTextbox() {
	log.Println("Formatting Textbox")
	bot := d.Doc.Height() - sk.Tolerances["flavor"]

	if d.Speed.Visible() {
		d.Speed.SetColor(ps.Colors["Gray"])
	}
	d.Short.SetVisible(d.Short.Text != nil)
	d.Long.SetVisible(d.Long.Text != nil && *d.Long.Text != "")
	d.Flavor.SetVisible(d.Flavor.Text != nil)

	d.ShortBG.SetPos(d.ShortBG.X1(), d.Short.Y2()+sk.Tolerances["short"], "BL")
	d.Long.SetPos(d.Long.X1(), d.ShortBG.Y2()+sk.Tolerances["long"], "TL")
	d.Flavor.SetPos(d.Flavor.X1(), bot, "BL")

	if d.Long.Visible() {
		if d.Long.Y2() > bot {
			d.Long.SetVisible(false)
		} else {
			if d.Flavor.Visible() && d.Long.Y2() > d.Flavor.Y1() {
				d.Flavor.SetVisible(false)
			}
		}
	}
}

// Save saves a copy the produced card image as a .png in the appropriate
// subfolder of  "SK_OUT".
// If crop is true, the bleed area around the card is cropped out of the image
// before saving.
func (d *DeckTemplate) Save(crop bool) {
	leader := "Heroes" // TODO: Fix skirmish.Leader(lyr.TextItem)
	if !crop {
		err := ps.SaveAs(filepath.Join(os.Getenv("SK_PS"), "Decks", leader,
			*d.ID.Text))
		if err != nil {
			log.Panic(err)
		}
		return
	}
	err := ps.DoAction("DK", "Crop")
	if err != nil {
		panic(err)
	}
	err = ps.SaveAs(filepath.Join(os.Getenv("SK_PS"), "Decks", leader,
		*d.ID.Text))
	if err != nil {
		panic(err)
	}
	err = ps.DoAction("DK", "Undo")
	if err != nil {
		panic(err)
	}
}
