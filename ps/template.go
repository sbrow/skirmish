package ps

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sbrow/ps"
	"github.com/sbrow/skirmish"
)

type Template interface {
	ApplyDataset(id string)
	GetDoc() *ps.Document
	PNG(crop bool) error
}
type template struct {
	Doc           *ps.Document
	ResolveSymbol *ps.LayerSet
	Card          skirmish.Card
	Dataset       string
	ID            *ps.ArtLayer
	Name          *ps.ArtLayer
	Resolve       *ps.ArtLayer
	Speed         *ps.ArtLayer
	Life          *ps.ArtLayer
	Damage        *ps.ArtLayer
	Short         *ps.ArtLayer
	Long          *ps.ArtLayer
	Flavor        *ps.ArtLayer
	ShortBG       *ps.ArtLayer
	ResolveBG     *ps.ArtLayer
	DeckInd       *ps.LayerSet
	SpeedBG       *ps.ArtLayer
	LifeBG        *ps.ArtLayer
}

// TODO(sbrow): Recover - run in safe mode.
func New(mode ps.ModeEnum, file string) *template {
	t := &template{}
	ps.Open(file)
	ps.Mode = mode
	log.Printf("Creating new template with mode %d", mode)
	doc, err := ps.ActiveDocument()
	if err != nil {
		log.Fatal(err)
	}
	t.Doc = doc

	t.ResolveSymbol = doc.MustExist("ResolveGem").(*ps.LayerSet)

	txt := doc.MustExist("Text").(*ps.LayerSet)
	t.Name = txt.MustExist("name").(*ps.ArtLayer)
	t.ID = txt.MustExist("id").(*ps.ArtLayer)
	t.Resolve = txt.MustExist("resolve").(*ps.ArtLayer)
	t.Speed = txt.MustExist("speed").(*ps.ArtLayer)
	t.Life = txt.MustExist("life").(*ps.ArtLayer)
	t.Damage = txt.MustExist("damage").(*ps.ArtLayer)
	t.Short = txt.MustExist("short").(*ps.ArtLayer)
	t.Long = txt.MustExist("long").(*ps.ArtLayer)
	t.Flavor = txt.MustExist("flavor").(*ps.ArtLayer)

	areas := doc.MustExist("Areas").(*ps.LayerSet)
	bottom := areas.MustExist("Bottom").(*ps.LayerSet)
	resolveBG := areas.MustExist("ResolveBackground").(*ps.LayerSet)
	t.ShortBG = bottom.ArtLayer("short_text_box")
	t.ResolveBG = resolveBG.MustExist("resolve_color").(*ps.ArtLayer)
	ind := doc.MustExist("Indicators").(*ps.LayerSet)
	t.DeckInd = ind.MustExist("Deck").(*ps.LayerSet)
	t.SpeedBG = ind.MustExist("speed_background").(*ps.ArtLayer)
	return t
}

func (t *template) AddSymbols() {
	// Confirm that there is a resolve symbol in the text.
	reg, err := regexp.Compile("({[1-9]})")
	if err != nil {
		t.Doc.Dump()
		log.Panic(err)
	}
	temp := reg.FindStringIndex(t.Short.TextItem.Contents())
	if temp == nil {
		t.ResolveSymbol.SetVisible(false)
		return
	}
	t.ResolveSymbol.SetVisible(true)

	// Reverse engineer the line breaks in the text.
	lineHeight := 30
	var bnd [2][2]int
	words := strings.Split(strings.Replace(t.Short.TextItem.Contents(), "\r", "\\r ", -1), " ")
	out := words[0]
	for _, word := range words[1:] {
		tmp := out
		if !strings.HasSuffix(out, "\\r") {
			tmp += " "
		}
		tmp += word
		bnd = t.Short.Bounds()
		t.Short.TextItem.SetText(tmp)

		// t.Bold()
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
	t.Short.TextItem.SetText(out)

	// Find the resolve symbol
	rows := strings.Split(out, "\\r")
	for i, r := range rows {
		temp = reg.FindStringIndex(r)
		if temp != nil {
			// Get the BR y value.
			if i+1 != len(rows) {
				t.Short.TextItem.SetText(strings.Join(rows[:i+1], "\\r"))
			}
			y := t.Short.Y2()
			// Get the BR x val
			t.Short.TextItem.SetText(rows[i][:temp[1]])
			t.Bold()
			t.Short.Refresh()
			x := t.Short.X2()

			// Move it.
			t.ResolveSymbol.SetVisible(true)
			t.ResolveSymbol.SetPos(x+13, y+7, "BR")
		}
	}
	t.Short.TextItem.SetText(reg.ReplaceAllString(out, " $1"))
}

// ApplyDataset applies the dataset with given id and name
func (t *template) ApplyDataset(id string) {
	if ps.Mode == ps.Fast && t.Dataset == id {
		return
	}
	log.Printf("Applying dataset %s\n", id)
	log.SetPrefix(fmt.Sprintf("[%s] ", id))
	name := strings.TrimRight(id, "_123")
	if t.Card == nil {
		card, err := skirmish.Load(name)
		if err != nil {
			t.Doc.Dump()
			log.Println(card)
			log.Panic(err)
		}
		t.Card = card
	}
	ps.ApplyDataset(id)
	for _, lyr := range t.Doc.ArtLayers() {
		if lyr.Name() == "card_image" {
			lyr.Refresh()
		}
	}
	t.Flavor.Refresh()
	t.ID.Refresh()
	// TODO(sbrow): Skip the rest if id is different but name is not
	if t.Name.TextItem.Contents() == name && t.ID.TextItem.Contents() != id {
		fmt.Println("Skipping")
		return
	}

	t.Name.Refresh()
	t.Resolve.Refresh()
	t.Speed.Refresh()
	t.Life.Refresh()
	t.Damage.Refresh()
	t.Short.Refresh()
	t.Long.Refresh()
	// TODO(sbrow): (5) pprof: Improved, but can still be better.
	for _, lyr := range t.DeckInd.ArtLayers() {
		lyr.Refresh()
	}
	t.ResolveBG.Refresh()
	t.SpeedBG.Refresh()
}

func (t *template) Bold() error {
	reg, err := regexp.Compile(t.Card.Regexp())
	if err != nil {
		fmt.Println(t.Card.Regexp())
		t.Doc.Dump()
		log.Println(t.Card.Name())
		return err
	}
	bold := reg.FindAllStringIndex(t.Short.TextItem.Contents(), -1)
	t.Short.SetActive()
	t.Short.Fmt(0, len(t.Short.TextItem.Contents()), "Arial", "Regular")
	for _, rng := range bold {
		t.Short.TextItem.Fmt(rng[0], rng[1], "Arial", "Bold")
	}
	return nil
}

// FormatTextbox arranges text and background layers inside the textbox, hiding
// layers as necessary.
func (t *template) FormatTextbox() {
	log.Println("Formatting Textbox")
	bot := t.Doc.Height() - Tolerances["bottom"]

	if t.Speed.Visible() {
		t.Speed.SetColor(ps.ColorGray)
	}
	t.Short.SetVisible(t.Short.TextItem != nil)
	t.Long.SetVisible(t.Long.TextItem != nil && t.Long.TextItem.Contents() != "")
	t.Flavor.SetVisible(t.Flavor.TextItem != nil)

	t.AddSymbols()
	t.Bold()

	t.ShortBG.SetPos(t.ShortBG.X1(), t.Short.Y2()+Tolerances["short"], "BL")
	t.Long.SetPos(t.Long.X1(), t.ShortBG.Y2()+Tolerances["long"], "TL")
	t.Flavor.SetPos(t.Flavor.X1(), t.Doc.Height()-Tolerances["flavor"], "BL")

	if t.Long.Visible() {
		if t.Long.Y2() > bot {
			t.Long.SetVisible(false)
		}
		if t.Flavor.Visible() && t.Long.Y2() > t.Flavor.Y1() {
			t.Flavor.SetVisible(false)
		}
	}
}

// PNG saves a copy the produced card image as a .png file in the appropriate
// subfolder of  "SK_OUT".
// If crop is true, the bleed area around the card is cropped out of the image
// before saving.
func (t *template) PNG(crop bool) error {
	log.Println("Saving copy as PNG")

	path := filepath.Join(os.Getenv("SK_PS"), "Decks", t.Card.Leader(), t.ID.TextItem.Contents())
	if t.Card.Leader() == "" {
		path = strings.Replace(path, "//", "/Heroes/", 1)
	}
	if crop {
		err := ps.DoAction("DK", "Crop")
		if err != nil {
			return err
		}
		defer func() {
			err := ps.DoAction("DK", "Undo")
			if err != nil {
				t.Doc.Dump()
				log.Panic(err)
			}
		}()
	}

	return ps.SaveAs(filepath.Join(os.Getenv("SK_PS"), "Decks", t.Card.Leader(),
		t.ID.TextItem.Contents()))
}
func (t *template) SetLeader(name string) (banner, ind ps.Hex, barStroke ps.Stroke, err error) {
	for _, ldr := range skirmish.Leaders {
		if ldr.Name == name {
			banner = ldr.Banner
			ind = ldr.Indicator
			barStroke = ps.Stroke{Size: 4, Color: banner}
		}
		leaderInd := t.DeckInd.MustExist(ldr.Name).(*ps.ArtLayer)
		if ind != nil {
			leaderInd.SetVisible(ldr.Name == name)
		} else {
			err := fmt.Errorf("no Layer found at \"%s%s\"", t.DeckInd.Path(), ldr.Name)
			return banner, ind, barStroke, err
		}
	}
	if banner == nil || ind == nil {
		return banner, ind, barStroke, fmt.Errorf("Leader \"%s\" not found!", name)
	}
	t.ResolveBG.SetColor(banner)
	t.Speed.SetStroke(barStroke, ps.ColorGray)

	t.SpeedBG.SetColor(ind)
	t.Damage.SetStroke(barStroke, ps.ColorWhite)
	return banner, ind, barStroke, nil
}
