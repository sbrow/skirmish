package sql

/*
// DeckCard represents a unique card in the deck. Contains most information
// required for updating the Photoshop document.
type DeckCard struct {
	Card               // Basic card values.
	Rarity             // How many copies of the card are in the deck.
	Arts   int         // How many different illustrations this card has.
	Leader NonDeckCard // The leader of the deck.
	Cost   string      // The resolve cost of the card.
}

// NewDeckCard constructs a new card with default values.
func NewDeckCard() *DeckCard {
	c := NewCard()
	lbl := []string{"Cost", "Common", "Uncommon", "Rare", "border_normal"}
	lbl = append(lbl, Leaders...)
	c.labels = append(c.labels, lbl...)
	// c.dir = leader.Name // TODO:

	return &DeckCard{Card: *c,
		Rarity: Common,
		Cost:   "1",
	}
}

func (d *DeckCard) String() string {
	return d.Card.String()
}

// DefaultBorder returns the visibility of the default border layer.
//
// All cards use the default border except actions, continuous events and heroes.
func (d *DeckCard) DefaultBorder() bool {
	switch {
	case d.Rarity == Rare:
		fallthrough
	case d.Type == Action:
		fallthrough
	case d.Type == Continuous:
		fallthrough
	case d.Type == Hero:
		return false
	default:
		return true
	}
}
*/
