// TODO: Add test file.
//
package ps

import (
	"errors"
	"github.com/sbrow/ps"
	"github.com/sbrow/skirmish/sql"
	"log"
)

var tolerances map[string]int
var doc *ps.Document

var txt *ps.LayerSet
var name *ps.ArtLayer
var cost *ps.ArtLayer
var resolve *ps.ArtLayer
var typ *ps.ArtLayer
var speed *ps.ArtLayer
var life *ps.ArtLayer
var damage *ps.ArtLayer
var short *ps.ArtLayer
var long *ps.ArtLayer
var flav *ps.ArtLayer
var heroLife *ps.ArtLayer

var areas *ps.LayerSet
var bottom *ps.LayerSet
var shortbg *ps.ArtLayer
var lBar *ps.ArtLayer
var resolve_bg *ps.ArtLayer
var cost_bg *ps.ArtLayer
var indicators []*ps.ArtLayer
var rarities []*ps.ArtLayer
var types []*ps.ArtLayer
var deck *ps.LayerSet

func init() {
	ps.Mode = ps.Normal
	// ps.Mode = ps.Safe
	// ps.Mode = ps.Fast
	log.Printf("Testing with mode %d", ps.Mode)
	tolerances = make(map[string]int)
	rows, err := sql.Database.Query("SELECT name, px FROM tolerances;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var px int
		if err := rows.Scan(&name, &px); err != nil {
			log.Fatal(err)
		}
		tolerances[name] = px
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	doc, err = ps.ActiveDocument()
	if err != nil {
		log.Fatal(err)
	}
	txt = doc.LayerSet("Text")
	if txt == nil {
		log.Panic("LayerSet \"Text\" was not found!")
	}
	name = txt.ArtLayer("name")
	if name == nil {
		log.Panic("ArtLayer \"name\" was not found!")
	}
	cost = txt.ArtLayer("cost")
	if cost == nil {
		log.Panic("ArtLayer \"cost\" was not found!")
	}
	resolve = txt.ArtLayer("resolve")
	if resolve == nil {
		log.Panic("ArtLayer \"resolve\" was not found!")
	}
	typ = txt.ArtLayer("type")
	if typ == nil {
		log.Panic("ArtLayer \"type\" was not found!")
	}
	speed = txt.ArtLayer("speed")
	if speed == nil {
		log.Panic("ArtLayer \"speed\" was not found!")
	}
	life = txt.ArtLayer("life")
	if life == nil {
		log.Panic("ArtLayer \"life\" was not found!")
	}
	damage = txt.ArtLayer("damage")
	if damage == nil {
		log.Panic("ArtLayer \"damage\" was not found!")
	}
	short = txt.ArtLayer("short")
	if short == nil {
		log.Panic("ArtLayer \"short\" was not found!")
	}
	long = txt.ArtLayer("long")
	if long == nil {
		log.Panic("ArtLayer \"long\" was not found!")
	}
	flav = txt.ArtLayer("flavor")
	if flav == nil {
		log.Panic("ArtLayer \"flav\" was not found!")
	}
	heroLife = txt.ArtLayer("hero life")
	if heroLife == nil {
		log.Panic("ArtLayer \"heroLife\" was not found!")
	}

	areas = doc.LayerSet("Areas")
	if areas == nil {
		log.Panic("LayerSet \"Areas\" was not found!")
	}
	bottom = areas.LayerSet("Bottom")
	if bottom == nil {
		log.Panic("LayerSet \"Bottom\" was not found!")
	}
	shortbg = bottom.ArtLayer("short_text_box")
	if shortbg == nil {
		log.Panic("ArtLayer \"shortbg\" was not found!")
	}
	lBar = bottom.ArtLayer("L Bar")
	if lBar == nil {
		log.Panic("ArtLayer \"L Bar\" not found!")
	}

	resolve_bg = areas.LayerSet("ResolveBackground").
		ArtLayer("resolve_color")
	if resolve_bg == nil {
		log.Panic("resolve_bg layer not found!")
	}
	cost_bg = areas.LayerSet("CostBackground").
		ArtLayer("cost_color")
	if cost_bg == nil {
		log.Panic("cost_bg layer not found!")
	}
	ind := doc.LayerSet("Indicators")
	if ind == nil {
		log.Panic("LayerSet \"Indicators\" was not found!")
	}
	types = ind.LayerSet("Type").ArtLayers()
	if types == nil {
		log.Panic("Type layers not found!")
	}
	rarities = ind.LayerSet("Rarity").ArtLayers()
	if rarities == nil {
		log.Panic("Rarity layers not found!")
	}
	indicators = ind.ArtLayers()
	if indicators == nil {
		log.Panic("[]ArtLayers \"Indicators\" were not found!")
	}
	deck = ind.LayerSet("Deck")
	if deck == nil {
		log.Panic("LayerSet \"Deck\" was not found!")
	}
}

// ApplyDataset applies the given dataset,
// selects card data from the sql server,
// checks all its values against the active document,
// updates any fields that were changed,
// and then calls any necessary formatting functions.
func ApplyDataset(name string) {

}

// Save saves a copy the produced card image as a .png in the appropriate
// subfolder of  "SK_OUT".
// If crop is true, the bleed area around the card is cropped out of the image
// before saving.
// TODO: text layers
/*
func Save(crop bool, args ...string) {
	lyr := doc.LayerSet("Text").ArtLayer("id")
	if lyr == nil {
		log.Panicf("Layer \"%s/%s\" not found", "Text", "id")
	}
	leader := "Heroes" // TODO: Fix skirmish.Leader(lyr.TextItem)
	if !crop {
		err := ps.SaveAs(filepath.Join(os.Getenv("SK_OUT"), leader, lyr.TextItem))
		if err != nil {
			panic(err)
		}
		return
	}
	err := ps.DoAction("DK", "Crop")
	if err != nil {
		panic(err)
	}
	err = ps.SaveAs(filepath.Join(os.Getenv("SK_OUT"), leader, lyr.TextItem))
	if err != nil {
		panic(err)
	}
	err = ps.DoAction("DK", "Undo")
	if err != nil {
		panic(err)
	}
}
*/

func SetLeader(name string) {
	var banner ps.Hex
	var indicator ps.Hex
	err := sql.Database.QueryRow(
		"SELECT banner, indicator FROM public.leaders WHERE name=$1", name).
		Scan(&banner, &indicator)
	if err != nil {
		log.Panic(err)
	}

	rarity := ps.Compare(banner, indicator)
	barStroke := ps.Stroke{Size: 4, Color: banner}
	counterStroke := ps.Stroke{Size: 4, Color: indicator}

	resolve_bg.SetColor(ps.Colors["Gray"])
	cost_bg.SetColor(banner)
	for _, lyr := range types {
		lyr.SetColor(indicator)
	}
	for i := 0; i < 3; i++ {
		rarities[i].SetColor(rarity)
	}
	lBar.SetColor(banner)
	for _, lyr := range indicators {
		lyr.SetColor(indicator)
	}

	for _, lyr := range deck.ArtLayers() {
		if lyr.Name() == name {
			lyr.SetVisible(true)
		} else {
			lyr.SetVisible(false)
		}
	}
	cost.SetStroke(counterStroke, ps.Colors["White"])
	resolve.SetStroke(counterStroke, ps.Colors["White"])

	speed.SetStroke(barStroke, ps.Colors["White"])
	damage.SetStroke(barStroke, ps.Colors["White"])
	life.SetStroke(barStroke, ps.Colors["White"])
	heroLife.SetStroke(barStroke, ps.Colors["White"])
}

// Format rearranges, hides, and colors layers as appropriate.
func Format() {
	err := FormatTitle()
	if err != nil {
		panic(err)
	}
	FormatTextbox()
	// err = FormatTextbox()
	// if err != nil {
	// 	panic(err)
	// }
	SetLeader("Igrath")
	/*	err = SetLeader("Igrath")
		if err != nil {
			panic(err)
		}
	*/
}

// FormatTitle finds the correct length background for the card's title, makes
// it visible, and hides the rest. Returns an error if the title was longer than
// the longest background.
func FormatTitle() error {
	banners := doc.LayerSet("Areas").LayerSet("TitleBackground")
	tol := tolerances["title"]
	found := false

	// Search the TitleBackground layers;
	// show the shortest one that fits,
	// if it's not visible already;
	// hide all other layers,
	// if they aren't already hidden;
	for _, lyr := range banners.ArtLayers() {
		if !found && name.Bounds()[1][0]+tol <= lyr.Bounds()[1][0] {
			found = true
			if !lyr.Visible() {
				lyr.SetVisible(true)
			}
		} else {
			if lyr.Visible() {
				lyr.SetVisible(false)
			}
		}
	}
	if !found {
		return errors.New("Title too long.")
	} else {
		return nil
	}
}

// FormatTextbox arranges text and background layers inside the textbox, hiding
// layers as necessary.
// TODO: Logic incomplete in terms of layer hiding.
func FormatTextbox() {
	bot := doc.Height() - tolerances["flavor"]

	if speed.Visible() { // && speed.Text() == 1 {
		speed.SetColor(ps.Colors["Gray"])
	}
	/*
		short.SetVisible(short.Text() != "“")
		long.SetVisible(long.Text() != "“")
		flav.SetVisible(flav.Text() != "“")
	*/

	shortbg.SetPos(shortbg.X1(), short.Y2()+tolerances["short"], "BL")
	long.SetPos(long.X1(), shortbg.Y2()+tolerances["long"], "TL")
	flav.SetPos(flav.X1(), bot, "BL")

	if long.Visible() {
		if long.Y2() > bot {
			long.SetVisible(false)
		} else {
			if flav.Visible() && long.Y2() > flav.Y1() {
				flav.SetVisible(false)
			}
		}
	}
}
