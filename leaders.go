package skirmish

import (
	"log"
	"path/filepath"
)

// TODO(sbrow): Move these vars to appropriate places.

// ImageDir is the path to the root directory where card images are located.
var ImageDir string

// DefaultImage is the name of the image file to use when
// a card's image cannot be found.
var DefaultImage = "ImageNotFound.png"

// Leaders is the list of valid deck leaders.
var Leaders leaders

func init() {
	if db == nil {
		if err := Connect(Cfg.DBArgs()); err != nil {
			log.Println(err)
		}
	}
	Leaders = []leader{}
	if err := Leaders.load(); err != nil {
		log.Println(err)
	}
	ImageDir = filepath.Join(Cfg.PS.Dir, "Images")
}

type leader struct {
	Name      string
	Banner    []uint8
	Indicator []uint8
}
type leaders []leader

func (l *leaders) names() []string {
	s := make([]string, len(*l))
	for i, ldr := range *l {
		s[i] = ldr.Name
	}
	return s
}

func (l *leaders) load() error {
	rows, err := Query(
		`SELECT "name", banner, indicator FROM leaders`)
	if err != nil {
		return err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		var name string
		var banner []uint8
		var indicator []uint8
		if err := rows.Scan(&name, &banner, &indicator); err != nil {
			return err
		}
		next := leader{name, banner, indicator}
		if i >= len(*l) {
			*l = append(*l, next)
		} else {
			(*l)[i] = next
		}
		i++
	}
	return nil
}
