package deck

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Card struct {
	Name       string // The name of the card.
	Arts       int    // The number of unique arts the card has.
	Type       Type   // The card's type.
	Resolve    int    // The resolve this card produces when in play.
	Speed      int    // The speed the card has (if it's a character).
	Damage     int    // The damage the card deals in combat (if it's a character).
	Toughness  int    // The damage the card can take before being discarded (if it's a follower).
	Life       int    // The damage the card can take before being discarded (if it's a hero).
	ShortText  string // The card's basic rules text.
	LongText   string // The card's reminder text.
	FlavorText string // The card's flavor (non-rules) text.
}

// ID returns an id unique to the card.
//
// Cards with only one art will have identical ID and Name.
// Cards with more than one art will have an ID containing their name
// and which version of art they use.
func (c *Card) ID(ver int) string {
	if c.Arts > 1 {
		return fmt.Sprintf("\"%s_%d\"", c.Name, ver)
	} else {
		return fmt.Sprintf("\"%s\"", c.Name)
	}
}

// Image builds and returns a path to the card's illustration.
// dir should point to the deck's leader for DeckCards and to "Heroes"
// for NonDeckCards.
//
// path = [$SK_SRC]/[folder]]/[c.Name].png
func (c *Card) Image(folder string, ver ...int) (path string, err error) {
	path = fmt.Sprintf(filepath.Join(os.Getenv("SK_IMG"), folder))
	if c.Arts == 1 {
		path = filepath.Join(path, c.Name+".png")
	} else {
		path = filepath.Join(path, c.Name)
		dir, err := ioutil.ReadDir(path)
		if err != nil {
			log.SetPrefix("[ERROR]")
			log.Print(path, " does not exist!")
			return "", err
		}
		if len(ver) == 0 {
			ver = append(ver, 1)
		}
		path = filepath.Join(path, dir[ver[0]-1].Name())
	}
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = errors.New(fmt.Sprint(path, " does not exist!"))
	}
	return path, err
}

type NonDeckCard struct {
	Faction Faction
	Card
}

// DeckCard represents a unique card in the deck. Contains most information required for
// updating the Photoshop document.
type DeckCard struct {
	Card
	Deck   *Deck
	Rarity        // How many copies of the card are in the deck.
	Cost   string // The resolve cost of the card.
}

// NewDeckCard constructs a new card with default values.
func NewDeckCard() *DeckCard {
	return &DeckCard{
		Card: Card{
			Name:       "Card",
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
		},
		Rarity: Common,
		Cost:   "1",
	}
}

// DefaultBorder returns the visibility of the default border layer.
//
// All cards use the default border except actions, continuous events and heroes.
func (c *DeckCard) DefaultBorder() bool {
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

func (c *DeckCard) setArts(n int) {
	if n >= 1 && n <= c.Rarity.Int() {
		c.Arts = n
	}
}
