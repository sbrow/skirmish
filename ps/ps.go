// TODO: Add test file.
//
package ps

import (
	"errors"
	"github.com/sbrow/ps"
	"os"
	"path/filepath"
)

// Save saves a copy the produced card image as a .png in the appropriate
// subfolder of  "SK_OUT".
// If crop is true, the bleed area around the card is cropped out of the image
// before saving.
func Save(crop bool, args ...string) {
	lyr, err := ps.Layer("Text/id")
	if err != nil {
		panic(err)
	}
	leader := "Heroes" // TODO: Fix skirmish.Leader(lyr.TextItem)
	/*	if err != nil {
			panic(err)
		}
	*/if !crop {
		err = ps.SaveAs(filepath.Join(os.Getenv("SK_OUT"), leader, lyr.TextItem))
		if err != nil {
			panic(err)
		}
		return
	}
	err = ps.DoAction("DK", "Crop")
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
//
// Scriptcalls: 4
//
// TODO: Combine script calls.
func FormatTitle() error {
	banners, err := ps.Layers("Areas/TitleBackground")
	if err != nil {
		panic(err)
	}
	txt, err := ps.Layer("Text/name")
	if err != nil {
		panic(err)
	}
	tol := 55 // TODO: pull from database

	// Start by hiding all the layers
	_, err = ps.DoJs("setLayerVisibility.jsx", "Areas/TitleBackground", "false")
	if err != nil {
		panic(err)
	}
	// Show the appropriate layer.
	for _, lyr := range banners {
		if txt.Bounds[1][0]+tol <= lyr.Bounds[1][0] {
			lyr.SetVisible()
			return nil
		}
	}
	return errors.New("Title too long.")
}
