// Package deck contains code for creating Photoshop data sets for Skirmish cards
// from json files.
//
// TODO: Sync the order between Label and Card.String()
// TODO: Implement images for cards with cameos.
// TODO: Add support for heroes.
// TODO: Enumerate Bordertype.
package deck

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	ACTION     Type = "Action"
	CONTINUOUS Type = "Event- Continuous"
	EVENT      Type = "Event"
	HERO       Type = "Deck Hero"
	ITEM       Type = "Item"
)

// Card represents a Unique card in the deck. Contains most information required for
// updating the Photoshop document.
type Card struct {
	Name       string // The name of the card.
	Leader     *Card  // The Leader of the deck.
	Rarity     int    // How many copies of the card are in the deck.
	Cost       int    // The resolve cost of the card.
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

// CardImage builds and returns a path to the card's illustration.
//
// path = [Home]/[c.Leader]/[c.Name].png
func (c *Card) CardImage(leader string) (path string) {
	path = fmt.Sprintf("\"%s\\%s\\%s.png\",", Home, leader, c.Name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.SetPrefix("[ERROR]")
		log.Print(path, " does not exist!")
	}
	return path
}

// DefaultBorder returns the visibility of the default border layer.
//
// All cards use the default border except actions, continuous events and heroes.
func (c *Card) DefaultBorder() bool {
	switch {
	case c.Rarity == Rare:
		fallthrough
	case c.Type == ACTION:
		fallthrough
	case c.Type == CONTINUOUS:
		fallthrough
	case c.Type == HERO:
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

/*
Deck represents a skirmish deck, with leader and deck cards.
*/
type Deck struct {
	Leader *Card
	Cards  [Size]Card
	labels []string
}

/*
New takes an input file and creates a Deck from the data.
Input must be in JSON format and have a ".json" extension.
	default:
		return true
*/
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
		"Name", "Cost", "Type", "Resolve",
		"Speed", "Damage", "Toughness", "Life",
		"ShortText", "LongText", "FlavorText",
		// "card_image,"
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
			v := reflect.ValueOf(card)
			for j := 1; j <= int(card.Rarity); j++ {
				str := ""
				for _, label := range d.labels {
					if (v.FieldByName(label) != reflect.Value{}) {
						switch v.FieldByName(label).Interface().(type) {
						case Type:
							str += fmt.Sprintf("\"%v\"", v.FieldByName(label))
						case string:
							str += fmt.Sprintf("\"%v\"", v.FieldByName(label))
						case int:
							if label == "Resolve" {
								str += fmt.Sprintf("\"%+d\"", v.FieldByName(label))
							} else {
								str += fmt.Sprintf("%d", v.FieldByName(label))
							}
						}
					} else {
						switch label {
						case "Common":
							fallthrough
						case "Uncommon":
							fallthrough
						case "Rare":
							str += fmt.Sprintf("%s", strings.Split(card.RarityString(), Delim)[Rarities[label]%3])
						case "Action":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(ACTION)))
						case "Event":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(EVENT)))
						case "Continuous":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(CONTINUOUS)))
						case "Item":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(ITEM)))
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
									strings.Contains(string(card.Type), string(HERO)))
						case "show_tough":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), "Follower"))
						case "show_life":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(HERO)))
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
