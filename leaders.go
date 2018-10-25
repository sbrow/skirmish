package skirmish

import (
	"log"
	"path/filepath"
	"reflect"
)

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
	Leaders = []Leader{}
	if err := Leaders.load(); err != nil {
		log.Println(err)
	}
	if Cfg == nil || reflect.DeepEqual(*Cfg, Config{}) {
		if err := Cfg.Load(filepath.Join(CfgDir, "config.yml")); err != nil {
			log.Println(err)
		}
	}
	ImageDir = filepath.Join(Cfg.PS.Dir, "Images")
}

// Leader represents a deck leader
type Leader struct {
	Name      string  // The Leader's name.
	Banner    []uint8 // The leader's Banner color (in hexadecimal format).
	Indicator []uint8 // The leader's Indicator color (in hexadecimal format).
}
type leaders []Leader

func (l *leaders) names() []string {
	s := make([]string, len(*l))
	for i, ldr := range *l {
		s[i] = ldr.Name
	}
	return s
}

func (l *leaders) load() error {
	rows, err := Query(
		`SELECT leaders.name, leaders.banner, leaders.indicator FROM leaders ORDER BY leaders.name ASC`)
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
		next := Leader{name, banner, indicator}
		if i >= len(*l) {
			*l = append(*l, next)
		} else {
			(*l)[i] = next
		}
		i++
	}
	return nil
}
