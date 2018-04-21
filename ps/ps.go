package ps

import (
	"errors"
	"fmt"
	"github.com/sbrow/ps"
	sk "github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/sql"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Template struct {
	Doc         *ps.Document
	ResolveSymb *ps.LayerSet
	Card        sk.Card
	Dataset     string
	ID          *ps.ArtLayer
	Name        *ps.ArtLayer
	Resolve     *ps.ArtLayer
	Speed       *ps.ArtLayer
	Life        *ps.ArtLayer
	Damage      *ps.ArtLayer
	Short       *ps.ArtLayer
	Long        *ps.ArtLayer
	Flavor      *ps.ArtLayer
	ShortBG     *ps.ArtLayer
	ResolveBG   *ps.ArtLayer
	DeckInd     *ps.LayerSet
	SpeedBG     *ps.ArtLayer
	LifeBG      *ps.ArtLayer
}

// TODO: Recover - run in safe mode.
func New(mode ps.ModeEnum, file string) *Template {
	t := &Template{}
	ps.Open(file)
	ps.Mode = mode
	log.Printf("Creating new template with mode %d", mode)
	doc, err := ps.ActiveDocument()
	if err != nil {
		log.Fatal(err)
	}
	t.Doc = doc
	t.ResolveSymb = doc.LayerSet("ResolveGem")
	if t.ResolveSymb == nil {
		log.Panic("LayerSet \"ResolveGem\" was not found!")
	}
	txt := doc.LayerSet("Text")
	if txt == nil {
		log.Panic("LayerSet \"Text\" was not found!")
	}
	t.Name = txt.ArtLayer("name")
	if t.Name == nil {
		log.Panic("ArtLayer \"name\" was not found!")
	}
	t.ID = txt.ArtLayer("id")
	if t.ID == nil {
		log.Panic("ArtLayer \"id\" was not found!")
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
	t.SpeedBG = ind.ArtLayer("speed_background")
	if t.SpeedBG == nil {
		log.Panic("ArtLayer \"speed_background\" was not found!")
	}
	return t
}

func (t *Template) ApplyDataset(id, name string) {
	if ps.Mode == ps.Fast && t.Dataset == id {
		return
	}
	if ps.Mode == ps.Normal {
		defer t.Doc.Dump()
	}
	log.Printf("Applying dataset %s\n", id)
	log.SetPrefix(fmt.Sprintf("[%s] ", id))
	if t.Card == nil {
		card, err := sql.Load(name)
		if err != nil {
			log.Println(card)
			log.Panic(err)
		}
		t.Card = card
	}
	ps.ApplyDataset(id)
	t.Name.Refresh()
	t.ID.Refresh()
	t.Resolve.Refresh()
	t.Speed.Refresh()
	t.Life.Refresh()
	t.Damage.Refresh()
	t.Short.Refresh()
	t.Long.Refresh()
	t.Flavor.Refresh()
	t.DeckInd.Refresh()
	t.ResolveBG.Refresh()
	t.SpeedBG.Refresh()
	for _, lyr := range t.Doc.ArtLayers() {
		if lyr.Name() == "card_image" {
			lyr.Refresh()
		}
	}
}

func (t *Template) SetLeader(name string) (banner, ind ps.Hex) {
	if ps.Mode == ps.Normal {
		defer t.Doc.Dump()
	}
	for _, ldr := range sk.Leaders {
		if ldr.Name == name {
			banner = ldr.Banner
			ind = ldr.Indicator
		}
	}
	if banner == nil || ind == nil {
		log.Panicf("Leader \"%s\" not found!", name)
	}
	barStroke := ps.Stroke{Size: 4, Color: banner}
	t.ResolveBG.SetColor(banner)
	t.SpeedBG.SetColor(ind)

	t.Speed.SetStroke(barStroke, nil)
	t.Damage.SetStroke(barStroke, ps.Colors["White"])
	return banner, ind
}

// FormatTextbox arranges text and background layers inside the textbox, hiding
// layers as necessary.
func (t *Template) FormatTextbox() {
	if ps.Mode == ps.Normal {
		defer t.Doc.Dump()
	}
	log.Println("Formatting Textbox")
	bot := t.Doc.Height() - sk.Tolerances["bottom"]

	if t.Speed.Visible() {
		t.Speed.SetColor(ps.Colors["Gray"])
	}
	t.Short.SetVisible(t.Short.Text != nil)
	t.Long.SetVisible(t.Long.Text != nil && *t.Long.Text != "")
	t.Flavor.SetVisible(t.Flavor.Text != nil)

	t.AddSymbols()
	bold, err := t.Card.Bold()
	if err != nil {
		log.Println(t.Card.Name(), err)
	}
	t.Short.SetActive()
	t.Short.Format(0, len(*t.Short.Text), "Arial", "Regular")
	for _, rng := range bold {
		t.Short.Format(rng[0], rng[1], "Arial", "Bold")
	}
	// t.Short.Refresh()

	t.ShortBG.SetPos(t.ShortBG.X1(), t.Short.Y2()+sk.Tolerances["short"], "BL")
	t.Long.SetPos(t.Long.X1(), t.ShortBG.Y2()+sk.Tolerances["long"], "TL")
	t.Flavor.SetPos(t.Flavor.X1(), t.Doc.Height()-sk.Tolerances["flavor"], "BL")

	if t.Long.Visible() {
		if t.Long.Y2() > bot {
			t.Long.SetVisible(false)
		} else {
			if t.Flavor.Visible() && t.Long.Y2() > t.Flavor.Y1() {
				t.Flavor.SetVisible(false)
			}
		}
	}
}

func (t *Template) AddSymbols() {
	if ps.Mode == ps.Normal {
		defer t.Doc.Dump()
	}
	// Confirm that there is a resolve symbol in the text.
	reg, err := regexp.Compile("{[1-9]}")
	if err != nil {
		log.Panic(err)
	}
	temp := reg.FindStringIndex(*t.Short.Text)
	if temp == nil {
		t.ResolveSymb.SetVisible(false)
		return
	}
	t.ResolveSymb.SetVisible(true)

	// Reverse engineer the line breaks in the text.
	lineHeight := 30
	var bnd [2][2]int
	words := strings.Split(strings.Replace(*t.Short.Text, "\r", "\\r ", -1), " ")
	out := words[0]
	for _, word := range words[1:] {
		tmp := out
		if !strings.HasSuffix(out, "\\r") {
			tmp += " "
		}
		tmp += word
		bnd = t.Short.Bounds()
		t.Short.SetText(tmp)
		switch {
		case t.Short.Y2()-bnd[1][1] >= lineHeight:
			if !strings.HasSuffix(out, "\\r") {
				out += "\\r"
			}
		default:
			out += " "
		}
		out += word
	}
	out = strings.Replace(out, "\\r ", "\\r", -1)
	t.Short.SetText(out)

	// Find the resolve symbol
	rows := strings.Split(out, "\\r")
	for i, r := range rows {
		temp = reg.FindStringIndex(r)
		if temp != nil {
			// Get the BR y value.
			if i+1 != len(rows) {
				t.Short.SetText(strings.Join(rows[:i+1], "\\r"))
			}
			y := t.Short.Y2()
			// Get the BR x val
			t.Short.SetText(rows[i][temp[0]:temp[1]])
			x := t.Short.X2()
			// Move it.
			t.ResolveSymb.SetVisible(true)
			t.ResolveSymb.SetPos(x+3, y+7, "BR")
		}
	}
	t.Short.SetText(out)
}

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
	d := &DeckTemplate{Template: *New(mode, sk.Template)}
	txt := d.Doc.LayerSet("Text")
	if txt == nil {
		log.Panic("LayerSet \"Text\" was not found!")
	}
	d.Banners = d.Doc.LayerSet("Areas").LayerSet("TitleBackground")
	if d.Banners == nil {
		log.Panic("LayerSet \"TitleBackground\" was not found!")
	}
	if ps.Mode == 2 {
		d.Dataset = *d.ID.Text
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
	bottom := areas.LayerSet("Bottom")
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
	// Skip if dataset already applied.
	if d.Dataset == id {
		return
	}

	if ps.Mode == ps.Normal {
		defer d.Doc.Dump()
	}
	name := strings.TrimRight(id, `_123`)
	card, err := sql.Load(name)
	if err != nil {
		log.Println(card)
		log.Panic(fmt.Sprintf("Card '%s' not found. Check your spelling.", name))
	}
	d.Card = card
	d.SetLeader(d.Card.Leader())
	d.Template.ApplyDataset(id, name)

	// Update layer data
	d.Cost.Refresh()
	d.Type.Refresh()
	d.HeroLife.Refresh()
	d.RarityInd.Refresh()
	d.HeroLifeBG.Refresh()
	d.DamageBG.Refresh()
	d.LifeBG.Refresh()
	d.TypeInd.Refresh()

	// TODO: Fix Border.Refresh()
	// doc.LayerSet("Border").Refresh()
	d.FormatTitle()
	d.FormatTextbox()
}

func (d *DeckTemplate) SetLeader(name string) {
	// TODO: Fix
	// if ps.Mode == 2 && name == d.Card.Leader() {
	// 	return
	// }
	if ps.Mode == ps.Normal {
		defer d.Doc.Dump()
	}
	banner, ind := d.Template.SetLeader(name)

	rarity := ps.Compare(banner, ind)
	barStroke := ps.Stroke{Size: 4, Color: banner}
	counterStroke := ps.Stroke{Size: 4, Color: ind}
	rarities := d.RarityInd.ArtLayers()

	if strings.Contains(strings.Join(d.Card.STypes(), ","), "Channeled") {
		d.CostBG.SetColor(rarity)
	} else {
		d.CostBG.SetColor(ps.Colors["Gray"])
	}
	for _, lyr := range d.TypeInd.ArtLayers() {
		lyr.SetColor(ind)
	}
	// Use indices instead of range because the bottom layer is the
	// rarity_background and we want it to stay black.
	for i := 0; i < 3; i++ {
		rarities[i].SetColor(rarity)
	}
	d.Resolve.SetStroke(counterStroke, ps.Colors["White"])
	d.Life.SetStroke(barStroke, ps.Colors["White"])

	d.LBar.SetColor(banner)
	d.HeroLifeBG.SetColor(ind)
	d.DamageBG.SetColor(ind)
	d.LifeBG.SetColor(ind)

	d.Cost.SetStroke(ps.Stroke{4, rarity}, ps.Colors["White"])
	d.HeroLife.SetStroke(barStroke, ps.Colors["White"])
}

// FormatTitle finds the correct length background for the card's title, makes
// it visible, and hides the rest. Returns an error if the title was longer than
// the longest background.
func (d *DeckTemplate) FormatTitle() error {
	if ps.Mode == ps.Normal {
		defer d.Doc.Dump()
	}
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

// PNG saves a copy the produced card image as a .png file in the appropriate
// subfolder of  "SK_OUT".
// If crop is true, the bleed area around the card is cropped out of the image
// before saving.
func (d *DeckTemplate) PNG(crop bool) {
	log.Println("Saving copy as PNG")
	if !crop {
		err := ps.SaveAs(filepath.Join(os.Getenv("SK_PS"), "Decks", d.Card.Leader(),
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
	err = ps.SaveAs(filepath.Join(os.Getenv("SK_PS"), "Decks", d.Card.Leader(),
		*d.ID.Text))
	if err != nil {
		panic(err)
	}
	err = ps.DoAction("DK", "Undo")
	if err != nil {
		panic(err)
	}
}

type NonDeckTemplate struct {
	Template
	Plus     *ps.ArtLayer
	HaloInd  *ps.LayerSet
	HeroInd  *ps.ArtLayer
	Factions *ps.LayerSet
	LBar     *ps.LayerSet
	BtmBG    *ps.ArtLayer
}

func NewNonDeck(mode ps.ModeEnum) *NonDeckTemplate {
	log.SetPrefix("[ps.NewNonDeck] ")
	n := &NonDeckTemplate{Template: *New(mode, sk.HeroTemplate)}
	areas := n.Doc.LayerSet("Areas")
	if areas == nil {
		log.Panic("LayerSet \"Areas\" was not found!")
	}
	txt := n.Doc.LayerSet("Text")
	if txt == nil {
		log.Panic("LayerSet \"Text\" was not found!")
	}
	n.Plus = txt.ArtLayer("+")
	if n.Plus == nil {
		log.Panic("ArtLayer \"Text\" was not found!")
	}
	n.BtmBG = areas.ArtLayer("bottom_color")
	if n.BtmBG == nil {
		log.Panic("ArtLayer \"bottom_color\" was not found!")
	}
	n.LBar = areas.LayerSet("LeaderBar")
	if n.LBar == nil {
		log.Panic("LayerSet \"LeaderBar\" was not found!")
	}
	ind := n.Doc.LayerSet("Indicators")
	if ind == nil {
		log.Panic("LayerSet \"Indicators\" was not found!")
	}
	n.HeroInd = ind.ArtLayer("HeroIcon")
	if n.HeroInd == nil {
		log.Panic("ArtLayer \"HeroIcon\" was not found!")
	}
	n.HaloInd = ind.LayerSet("Halo")
	if n.HaloInd == nil {
		log.Panic("LayerSet \"Halo\" was not found!")
	}
	n.Factions = ind.LayerSet("Faction")
	if n.Factions == nil {
		log.Panic("LayerSet \"Faction\" was not found!")
	}
	return n
}

func (n *NonDeckTemplate) ApplyDataset(name string) {
	if n.Dataset == name {
		return
	}
	if ps.Mode == ps.Normal {
		defer n.Doc.Dump()
	}
	id := name
	card, err := sql.Load(name)
	if err != nil {
		log.Println(card)
		log.Panic(err)
	}
	n.Card = card
	log.SetPrefix(fmt.Sprintf("[ps.%s] ", name))
	n.SetLeader(n.Card.Leader())
	if strings.Contains(name, "(Halo)") {
		tmp := strings.Split(name, " ")
		name = tmp[0]
	}
	n.SetLeader(name)
	n.Template.ApplyDataset(id, name)
}

func (n *NonDeckTemplate) SetLeader(name string) {
	if ps.Mode == ps.Normal {
		defer n.Doc.Dump()
	}
	banner, ind := n.Template.SetLeader(name)
	barStroke := ps.Stroke{Size: 4, Color: banner}
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
	n.Plus.SetStroke(barStroke, ps.Colors["White"])
	n.Resolve.SetStroke(barStroke, ps.Colors["White"])
	n.Life.SetStroke(ps.Stroke{Size: 0, Color: ind}, ps.Colors["Black"])
}
