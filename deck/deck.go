/*
Package deck contains code for creating Photoshop data sets for Skirmish cards
from json files.

Optimal Workflow:
	Photoshop plug-in (sheet -> .csv -> dataset) -> .pngs

Current Workflow:
	Google Sheet -> GAS -> json -> gocode -> dataset -> photoshop -> .pngs

TODO: Sync the order between Label and Card.String()
TODO: Implement images for cards with cameos.
TODO: One Speed Color
TODO: Add support for heroes.
TODO: Enumerate Bordertype.
*/
package deck

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	// "os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

//Home is the base path for program operation.
const Home = "F:\\GitLab\\dreamkeepers-psd\\Images"

//Size is the number of cards in a deck
const Size = 20

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

// Skirmish decks can't have more then 3 cards with the same name.
const (
	Common   Rarity = 3
	Uncommon Rarity = 2
	Rare     Rarity = 1
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
Card represents a Unique card in the deck. Contains most information required for
updating the Photoshop document.
>>>>>>> Stashed changes
*/
type Card struct {
	Name       string // The name of the card.
	Leader     *Card  // The Leader of the deck.
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
		Leader:     nil,
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
CardImage builds and returns a path to the card's illustration.

path = [Home]/[c.Leader]/[c.Name].png
*/
/*func (c *Card) CardImage(leader string) (path string) {
	path = fmt.Sprintf("\"%s\\%s\\%s.png\",", Home, leader, c.Name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.SetPrefix("[ERROR]")
		log.Print(path, " does not exist!")
	}
<<<<<<< Updated upstream
	return path
}
=======
	return
}*/

/*
DefaultBorder returns the visibility of the default border layer.

All cards use the default border except actions, continuous events and heroes.
*/
func (d *Card) DefaultBorder() bool {
	switch {
	case d.Rarity == Rare:
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
Deck represents a skirmish deck, with leader and deck cards.
*/
type Deck struct {
	Leader Card
	Cards  [Size]Card
}

/*
New takes an input file and creates a Deck from the data.
Input must be in JSON format and have a ".json" extension.
	default:
		return true
*/
func New(path string) (d *Deck) {
	d = &Deck{}
	contents, _ := ioutil.ReadFile(path)
	_ = json.Unmarshal(contents, &d.Cards)
	reg, _ := regexp.Compile(".json") // TODO: Fix this.
	d.Leader = Card{Name: reg.ReplaceAllString(filepath.Base(path), "")}
	for _, card := range d.Cards {
		card.Leader = &d.Leader
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
			for j := 1; j <= int(card.Rarity); j++ {
				if i == 1 {
					fmt.Println(j, int(card.Rarity))
				}
				str := wrapString(fmt.Sprint(card.Name, "_", j))
				str += fmt.Sprintf("%v,", card.Cost)
				str += wrapString(string(card.Type))
				str += fmt.Sprintf("%+d,", card.Resolve)
				str += fmt.Sprintf("%v,", card.Speed)
				str += fmt.Sprintf("%v,", card.Damage)
				str += fmt.Sprintf("%v,", card.Toughness)
				str += fmt.Sprintf("%v,", card.Life)
				str += wrapString(card.ShortText)
				str += wrapString(card.LongText)
				str += wrapString(card.FlavorText)
				// str += card.CardImage(d.Leader.Name)
				str += fmt.Sprintf("%v,", strings.Contains(string(card.Type), string(ACTION)))
				str += fmt.Sprintf("%v,", strings.Contains(string(card.Type), string(EVENT)))
				str += fmt.Sprintf("%v,", strings.Contains(string(card.Type), string(CONTINUOUS)))
				str += fmt.Sprintf("%v,", strings.Contains(string(card.Type), string(ITEM)))
				str += fmt.Sprintf("%v,", card.Rarity == Common)
				str += fmt.Sprintf("%v,", card.Rarity == Uncommon)
				str += fmt.Sprintf("%v,", card.Rarity == Rare)
				str += fmt.Sprintf("%v,", d.Leader.Name == BAST)
				str += fmt.Sprintf("%v,", d.Leader.Name == IGRATH)
				str += fmt.Sprintf("%v,", d.Leader.Name == LILITH)
				str += fmt.Sprintf("%v,", d.Leader.Name == VI)
				str += fmt.Sprintf("%v,", d.Leader.Name == RAVAT)
				str += fmt.Sprintf("%v,", d.Leader.Name == SCUTTLER)
				str += fmt.Sprintf("%v,", d.Leader.Name == TENDRIL)
				str += fmt.Sprintf("%v,", d.Leader.Name == WISP)
				str += fmt.Sprintf("%v,", d.Leader.Name == SCINTER)
				str += fmt.Sprintf("%v,", d.Leader.Name == TINSEL)
				str += fmt.Sprintf("%v,", card.Resolve != 0)
				// TODO: Clumsy
				str += fmt.Sprintf("%v,",
					strings.Contains(string(card.Type), "Follower") ||
						strings.Contains(string(card.Type), string(HERO)))
				str += fmt.Sprintf("%v,", strings.Contains(string(card.Type), "Follower"))
				str += fmt.Sprintf("%v,", strings.Contains(string(card.Type), string(HERO)))
				str += fmt.Sprintf("%v", card.DefaultBorder())
				out[i] += str + "\n"
			}
			if i == i {
				fmt.Println("================")
			}
		}(i, out)
	}
	wg.Wait()

	ret := ""
	for _, line := range out {
		ret += line
	}
	return ret
}

/*
wrapString wraps a string in double quotes.
*/
func wrapString(s string) string {
	return fmt.Sprintf("\"%s\",", s)
}

/*
Labels prints the column labels for .csv output.

Labels is order sensitive: changing the order of labels will break the output
unless a corresponding change is made in Deck.String.
TODO: Fix this issue.
*/
func Labels() string {
	str := "name,cost,type,resolve,"
	str += "speed,damage,toughness,life,"
	str += "short_text,long_text,flavor_text,"
	// str += "card_image,"
	str += "show_action,show_event,show_continuous,show_item,"
	str += "show_common,show_uncommon,show_rare,"
	str += "show_bast,show_igrath,show_lilith,show_vi,"
	str += "show_ravat,show_scuttler,show_tendril,show_wisp,"
	str += "show_scinter,show_tinsel,"
	str += "show_resolve,show_speed,show_tough,show_life,"
	str += "border_normal"
	str += "\n"
	return str
}
