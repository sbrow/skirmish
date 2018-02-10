package deck

import (
	"github.com/sbrow/skirmish"
)

// DeckCard represents a unique card in the deck. Contains most information
// required for updating the Photoshop document.
type DeckCard struct {
	Card               // Basic card values.
	Rarity             // How many copies of the card are in the deck.
	Leader NonDeckCard // The leader of the deck.
	Cost   string      // The resolve cost of the card.
}

// NewDeckCard constructs a new card with default values.
func NewDeckCard() *DeckCard {
	c := NewCard()
	lbl := []string{"Cost", "Common", "Uncommon", "Rare", "border_normal"}
	lbl = append(lbl, skirmish.Leaders...)
	c.labels = append(c.labels, lbl...)
	// c.dir = leader.Name // TODO:

	return &DeckCard{Card: *c,
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
