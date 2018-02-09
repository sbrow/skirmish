package photoshop

import (
	// "fmt"
	"github.com/sbrow/ps"
	"github.com/sbrow/skirmish"
	// "io"
	"os"
	"path/filepath"
	// "strings"
)

func Save(crop bool, args ...string) {
	if crop {
		err := ps.DoAction("DK", "Crop")
		if err != nil {
			panic(err)
		}
	}

	lyr, err := ps.Layer("Text/id")
	if err != nil {
		panic(err)
	}
	leader, err := skirmish.Leader(lyr.TextItem)
	if err != nil {
		panic(err)
	}
	err = ps.SaveAs(filepath.Join(os.Getenv("SK_OUT"), leader, lyr.TextItem))
	if err != nil {
		panic(err)
	}

	if crop {
		err = ps.DoAction("DK", "Undo")
		if err != nil {
			panic(err)
		}
	}
}
