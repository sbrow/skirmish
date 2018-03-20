// TODO: Add test file.
//
package ps

import (
	"errors"
	"github.com/sbrow/ps"
	"github.com/sbrow/skirmish/sql"
	"log"
	// "os"
	// "path/filepath"
)

var tolerances map[string]int
var doc *ps.Document

func init() {
	tolerances = make(map[string]int)
	rows, err := sql.Database.Query("SELECT name, px FROM public.tolerances")
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
	// stroke := ps.Stroke{Size: 4, Color: banner}
	// white := &ps.RGB{255, 255, 255}

	resolve := doc.LayerSet("Areas").LayerSet("ResolveBackground").
		ArtLayer("resolve_color")
	if resolve == nil {
		log.Panic("Resolve layer not found!")
	}
	cost := doc.LayerSet("Areas").LayerSet("CostBackground").
		ArtLayer("cost_color")
	if cost == nil {
		log.Panic("Cost layer not found!")
	}
	types := doc.LayerSet("Indicators").LayerSet("Type").ArtLayers()
	if types == nil {
		log.Panic("Type layers not found!")
	}
	rarities := doc.LayerSet("Indicators").LayerSet("Rarity").ArtLayers()
	if rarities == nil {
		log.Panic("Rarity layers not found!")
	}
	lBar := doc.LayerSet("Areas").LayerSet("Bottom").ArtLayer("L Bar")
	if lBar == nil {
		log.Panic("L Bar not found!")
	}
	indicators := doc.LayerSet("Indicators").ArtLayers()
	if indicators == nil {
		log.Panic("Indicators not found!")
	}

	resolve.SetColor(banner)
	cost.SetColor(banner)
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

}

// Format rearranges, hides, and colors layers as appropriate.
func Format() {
	err := FormatTitle()
	if err != nil {
		panic(err)
	}
}

// FormatTitle finds the correct length background for the card's title, makes
// it visible, and hides the rest. Returns an error if the title was longer than
// the longest background.
func FormatTitle() error {
	banners := doc.LayerSet("Areas").LayerSet("TitleBackground")
	txt := doc.LayerSet("Text").ArtLayer("name")
	tol := tolerances["title"]
	found := false

	// Search the TitleBackground layers;
	// show the shortest one that fits,
	// if it's not visible already;
	// hide all other layers,
	// if they aren't already hidden;
	for _, lyr := range banners.ArtLayers() {
		if !found && txt.Bounds()[1][0]+tol <= lyr.Bounds()[1][0] {
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
	txt := doc.LayerSet("Text")
	short := txt.ArtLayer("short_text")
	long := txt.ArtLayer("long_text")
	flav := txt.ArtLayer("flavor_text")
	shortbg := doc.LayerSet("Areas").LayerSet("Bottom").ArtLayer("short_text_box")

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

// FormatSpeed sets color values for speed when speed is 1.
func FormatSpeed() {
	speed := doc.LayerSet("Text").ArtLayer("speed")
	if speed.Visible() { // && speed.Text() == 1 {
		speed.SetColor(&ps.RGB{128, 128, 128})
	}
}
