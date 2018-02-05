// Package deck contains code for creating Photoshop data sets for Skirmish cards
// from json files.
//
// TODO: Implement ID in string.
// TODO: Implement ID in Photoshop.
// TODO: Add support for heroes.
package deck

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// pathpkg "path"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

// Home is the base path for program operation.
const Home = "F:\\GitLab\\dreamkeepers-psd\\Images"

// Size is the number of cards in a deck
const Size = 20

// Delim is the delimiter to use for csv/tsv output.
const Delim = ","

const (
	//Leader names
	BAST     = "Bast"
	IGRATH   = "Igrath"
	LILITH   = "Lilith"
	VI       = "Vi"
	RAVAT    = "Ravat"
	SCUTTLER = "Scuttler"
	TENDRIL  = "Tendril"
	WISP     = "Wisp"
	SCINTER  = "Scinter"
	TINSEL   = "Tinsel"
)

// Rarity determines how many copies of the card will be in the deck
type Rarity int

var Rarities = map[string]int{
	"Common":   3,
	"Uncommon": 2,
	"Rare":     1,
}

// Skirmish decks can't have more then 3 cards with the same name.
const (
	Common   = 3
	Uncommon = 2
	Rare     = 1
)

// Type is the variety of types a card can have.
type Type string

// This is not an expansive list, only the types that actually
// affect layer visibility.
const (
	Action     Type = "Action"
	Continuous Type = "Event- Continuous"
	Event      Type = "Event"
	Hero       Type = "Deck Hero"
	Item       Type = "Item"
)

// Card represents a Unique card in the deck. Contains most information required for
// updating the Photoshop document.
type Card struct {
	Name       string // The name of the card.
	Leader     *Card  // The Leader of the deck.
	Rarity     int    // How many copies of the card are in the deck.
	Cost       int    // The resolve cost of the card.
	Arts       int    // The number of unique arts the card has.
	Type       Type   // The card's type.
	Resolve    int    // The resolve this card produces when in play.
	Speed      int    // The speed the card has (if it's a character).
	Damage     int    // The damage the card deals in combat (if it's a character).
	Toughness  int    // The damage the card can take before being discarded (if it's a follower)
	Life       int    // The damage the card can take before being discarded (if it's a hero)
	ShortText  string // The card's basic rules text.
	LongText   string // The card's reminder text.
	FlavorText string // The card's flavor (non-rules) text.
}

// NewCard constructs a new card with default values.
func NewCard() *Card {
	return &Card{
		Name:       "Card",
		Rarity:     Common,
		Cost:       1,
		Arts:       1,
		Type:       "card_type",
		Resolve:    0,
		Speed:      1,
		Damage:     0,
		Toughness:  0,
		Life:       0,
		ShortText:  "",
		LongText:   "",
		FlavorText: "",
	}
}

func (c *Card) setArts(n int) {
	if n >= 1 && n <= c.Rarity {
		c.Arts = n
	}
}

func (c *Card) ID(ver int) string {
	if c.Arts > 1 {
		return fmt.Sprint(c.Name, "_", ver)
	} else {
		return c.Name
	}
}

// Image builds and returns a path to the card's illustration.
//
// path = [$SK_SRC]/[c.Leader]/[c.Name].png
func (c *Card) Image(leader string, ver ...int) (path string, err error) {
	path = fmt.Sprintf(filepath.Join(os.Getenv("SK_IMG"), leader))
	if c.Arts == 1 {
		path = filepath.Join(path, c.Name+".png")
	} else {
		path = filepath.Join(path, c.Name)
		folder, err := ioutil.ReadDir(path)
		if err != nil {
			log.SetPrefix("[ERROR]")
			log.Print(path, " does not exist!")
			return "", err
		}
		if len(ver) == 0 {
			ver = append(ver, 0)
		}
		path = filepath.Join(path, folder[ver[0]-1].Name())
	}
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = errors.New(fmt.Sprint(path, " does not exist!"))
	}
	return path, err
}

// DefaultBorder returns the visibility of the default border layer.
//
// All cards use the default border except actions, continuous events and heroes.
func (c *Card) DefaultBorder() bool {
	switch {
	case c.Rarity == Rare:
		fallthrough
	case c.Type == Action:
		fallthrough
	case c.Type == Continuous:
		fallthrough
	case c.Type == Hero:
		return false
	default:
		return true
	}
}

// TODO: Change this.
func (c *Card) RarityString() string {
	return fmt.Sprintf("%v%s%v%s%v",
		c.Rarity == Common, Delim,
		c.Rarity == Uncommon, Delim,
		c.Rarity == Rare,
	)
}

// Deck represents a skirmish deck, with leader and deck cards.
type Deck struct {
	Leader *Card
	Cards  [Size]Card
	labels []string
}

// New takes an input file and creates a Deck from the data.
// Input must be in JSON format and have a ".json" extension.
func New(path string) (d *Deck) {
	d = &Deck{}
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(contents, &d.Cards)
	if err != nil {
		panic(err)
	}
	reg, _ := regexp.Compile(".json") // TODO: Fix this.
	d.Leader = &Card{Name: reg.ReplaceAllString(filepath.Base(path), "")}
	d.labels = []string{
		"ID", "Name", "Cost", "Type", "Resolve",
		"Speed", "Damage", "Toughness", "Life",
		"ShortText", "LongText", "FlavorText",
		"card_image",
		"Common", "Uncommon", "Rare",
		"Action", "Event", "Continuous", "Item",
		"Bast", "Igrath", "Lilith", "Vi",
		"Ravat", "Scuttler", "Tendril", "Wisp",
		"Scinter", "Tinsel",
		"show_resolve", "show_speed", "show_tough", "show_life",
		"border_normal",
	}
	return d
}

func (d *Deck) String() string {
	var wg sync.WaitGroup
	wg.Add(len(d.Cards))
	out := make([]string, Size)
	for i := range d.Cards {
		go func(i int, out []string) {
			defer wg.Done()
			card := d.Cards[i]
			c := reflect.ValueOf(card)
			for v := 1; v <= int(card.Arts); v++ {
				str := ""
				for _, label := range d.labels {
					if (c.FieldByName(label) != reflect.Value{}) {
						switch c.FieldByName(label).Interface().(type) {
						case Type:
							str += fmt.Sprintf("\"%v\"", c.FieldByName(label))
						case string:
							str += fmt.Sprintf("\"%v\"", c.FieldByName(label))
						case int:
							if label == "Resolve" {
								str += fmt.Sprintf("\"%+d\"", c.FieldByName(label))
							} else {
								str += fmt.Sprintf("%d", c.FieldByName(label))
							}
						}
					} else {
						switch label {
						case "ID":
							str += card.ID(v)
						case "card_image":
							img, err := card.Image(d.Leader.Name, v)
							if err != nil {
								log.SetPrefix("[ERROR]")
								log.Println(err)
								log.SetPrefix("")

							} else {
								str += fmt.Sprintf("\"%s\"", img)
							}
						case "Common":
							fallthrough
						case "Uncommon":
							fallthrough
						case "Rare":
							str += fmt.Sprintf("%s", strings.Split(card.RarityString(), Delim)[Rarities[label]%3])
						case "Action":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(Action)))
						case "Event":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(Event)))
						case "Continuous":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(Continuous)))
						case "Item":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(Item)))
						case "Bast":
							fallthrough
						case "Igrath":
							fallthrough
						case "Lilith":
							fallthrough
						case "Vi":
							fallthrough
						case "Scinter":
							fallthrough
						case "Ravat":
							fallthrough
						case "Scuttler":
							fallthrough
						case "Tendril":
							fallthrough
						case "Wisp":
							fallthrough
						case "Tinsel":
							str += fmt.Sprintf("%v", d.Leader.Name == label)
						case "show_resolve":
							str += fmt.Sprintf("%v", card.Resolve != 0)
						case "show_speed":
							// TODO: Clumsy
							str += fmt.Sprintf("%v",
								strings.Contains(string(card.Type), "Follower") ||
									strings.Contains(string(card.Type), string(Hero)))
						case "show_tough":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), "Follower"))
						case "show_life":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(Hero)))
						case "border_normal":
							str += fmt.Sprintf("%v", card.DefaultBorder())
						}
					}
					str += Delim
				}
				out[i] += str[:len(str)-1] + "\n"
			}
		}(i, out)
	}
	wg.Wait()

	ret := ""
	for _, line := range out {
		ret += fmt.Sprintf("%s", line)
	}
	return ret[:len(ret)-2]
}

func (c *Card) checkRarityString(r Rarity) bool {
	return fmt.Sprintf("%d") == fmt.Sprintf("%d", r)
}

// wrapString wraps a string in double quotes.
func wrapString(s string) string {
	return fmt.Sprintf("\"%s\",", s)
}

// Labels prints the column labels for .csv output.
func (d *Deck) Labels() string {
	str := ""
	for _, label := range d.labels {
		str += fmt.Sprintf("%s%s", label, Delim)
	}
	return strings.ToLower(str[:len(str)-1]) + "\n"
}
