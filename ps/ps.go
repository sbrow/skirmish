// TODO: Add test file.
//
package ps

import (
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
