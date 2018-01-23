/*
Package deck contains code for creating Photoshop data sets for Skirmish cards
from json files.

Optimal Workflow:
	Spreadsheet -> gocode -> dataset -> photoshop -> .pngs

Current Workflow:
	Spreadsheet -> GAS -> json -> gocode -> dataset -> photoshop -> .pngs

TODO: Sync the order between Label and Card.String()
*/
package deck

import (
	"encoding/json"
	"fmt"
	"github.com/sbrow/debug"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

//HOME is the base path for program operation.
const HOME = "F:\\GitLab\\dreamkeepers-psd\\Images"

const (
	//Leader names
	BAST     = "Bast"
	IGRATH   = "Igrath"
	LILITH   = "Lilth"
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

// Skirmish decks can't have more then 3 cards with the same name.
const (
	COMMON   Rarity = 3
	UNCOMMON Rarity = 2
	RARE     Rarity = 1
)

// Type is the variety of types a card can have.
type Type string

/*
This is not an expansive list, only the types that actually
affect layer visibility.
*/
const (
	ACTION     Type = "Action"
	CONTINUOUS Type = "Event- Continuous"
	EVENT      Type = "Event"
	HERO       Type = "Deck Hero"
	ITEM       Type = "Item"
)

/*
Card represents a unique card in the deck. Contains most information required for
updating the Photoshop document.
*/
type Card struct {
	Name       string // The name of the card.
	Leader     string // Deprecated, TODO: convert this to a Card pointer.
	Rarity     Rarity // How many copies of the card are in the deck.
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

/*
NewCard constructs a new card with default values.
*/
func NewCard() *Card {
	return &Card{
		Leader:     BAST,
		Name:       "Card",
		Rarity:     3,
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

/*
CardImage returns the path to the card's illustration.
*/
func (c *Card) CardImage(leader string) (path string) {
	return fmt.Sprintf("%s\\%s\\%s.png,", HOME, leader, c.Name)
}

func (d *Card) normalBorder() bool {
	switch {
	case d.Rarity == RARE:
		fallthrough
	case d.Type == ACTION:
		fallthrough
	case d.Type == CONTINUOUS:
		fallthrough
	case d.Type == HERO:
		return false
	default:
		return true
	}
}

/*
Deck represents a deck, with leader and deck cards.
*/
type Deck struct {
	Leader Card
	Cards  []Card
}

/*
New takes an input file and creates a Deck from the data.
Input must be in JSON format and have a ".json" extension.
*/
func New(path string) (d *Deck) {
	d = &Deck{}
	contents, _ := ioutil.ReadFile(path)
	err := json.Unmarshal(contents, &d.Cards)
	debug.Check(err)
	reg, err := regexp.Compile("Go.json")
	debug.Check(err)

	d.Leader = Card{Name: reg.ReplaceAllString(filepath.Base(path), "")}
	return d
}

func (d *Deck) String() string {
	var wg sync.WaitGroup
	wg.Add(len(d.Cards))
	out := make([]string, len(d.Cards)+1)
	out[0] = Labels()
	for i, card := range d.Cards {
		go func(i int, out []string) {
			defer wg.Done()
			card := d.Cards[i]
			str := fmt.Sprintf("\"%v\",", card.Name)
			str += fmt.Sprintf("%v,", card.Cost)
			str += fmt.Sprintf("\"%v\",", card.Type)
			str += fmt.Sprintf("%v,", card.Resolve)
			str += fmt.Sprintf("%v,", card.Speed)
			str += fmt.Sprintf("%v,", card.Damage)
			str += fmt.Sprintf("%v,", card.Toughness)
			str += fmt.Sprintf("%v,", card.Life)
			str += fmt.Sprintf("\"%v\",", card.ShortText)
			str += fmt.Sprintf("\"%v\",", card.LongText)
			str += fmt.Sprintf("\"%v\",", card.FlavorText)
			str += card.CardImage(d.Leader.Name)
			str += fmt.Sprintf("%v,", strings.Contains(string(card.Type), string(ACTION)))
			str += fmt.Sprintf("%v,", strings.Contains(string(card.Type), string(EVENT)))
			str += fmt.Sprintf("%v,", strings.Contains(string(card.Type), string(CONTINUOUS)))
			str += fmt.Sprintf("%v,", strings.Contains(string(card.Type), string(ITEM)))
			str += fmt.Sprintf("%v,", card.Rarity == COMMON)
			str += fmt.Sprintf("%v,", card.Rarity == UNCOMMON)
			str += fmt.Sprintf("%v,", card.Rarity == RARE)
			str += fmt.Sprintf("%v,", d.Leader.Name == BAST)
			str += fmt.Sprintf("%v,", d.Leader.Name == IGRATH)
			str += fmt.Sprintf("%v,", d.Leader.Name == LILITH)
			str += fmt.Sprintf("%v,", d.Leader.Name == VI)
			str += fmt.Sprintf("%v,", d.Leader.Name == RAVAT)
			str += fmt.Sprintf("%v,", d.Leader.Name == SCUTTLER)
			// str += fmt.Sprintf("%v,", d.Leader.Name == SCINTER))
			// str += fmt.Sprintf("%v,", d.Leader.Name == TINSEL))
			str += fmt.Sprintf("%v,", d.Leader.Name == TENDRIL)
			str += fmt.Sprintf("%v,", d.Leader.Name == WISP)
			str += fmt.Sprintf("%v,", card.Resolve != 0)
			str += fmt.Sprintf("%v,",
				card.Type == "Follower" || card.Type == HERO)
			str += fmt.Sprintf("%v,", card.Type == "Follower")
			str += fmt.Sprintf("%v,", card.Type == HERO)
			str += fmt.Sprintf("%v", card.normalBorder())
			out[i+1] = str
		}(i, out)
	}

	wg.Wait()
	ret := ""
	for _, line := range out {
		ret += line + "\n"
	}
	return ret
}

/*
Labels prints the column labels for .csv output.
*/
func Labels() string {
	str := "name,cost,type,resolve,"
	str += "speed,damage,toughness,life,"
	str += "short_text,long_text,flavor_text,"
	str += "card_image,"
	str += "show_action,show_event,show_continuous,show_item,"
	str += "show_common,show_uncommon,show_rare,"
	str += "show_bast,show_igrath,show_lilith,show_vi,"
	str += "show_ravat,show_scuttler,show_tendril,show_wisp,"
	// str += "show_scinter,show_tinsel"
	str += "show_resolve,show_speed,show_tough,show_life,"
	str += "border_normal"
	return str
}
