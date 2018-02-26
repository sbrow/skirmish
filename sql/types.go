package sql

/*
import (
	"errors"
	"os"
)

// Delim is the delimiter to use for csv/tsv output.
const Delim = ","

// Home is the base path for program operation.
var Home = os.Getenv("SK_SRC")

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

// Rarity determines how many copies of a card will be in a deck
//
// TODO: Make private and manually unmarshal the jsons.
type Rarity int

var RarityError = errors.New("Rarity out of bounds")

func rarity(copies int) (Rarity, error) {
	switch copies {
	case 3:
		return Common, nil
	case 2:
		return Uncommon, nil
	case 1:
		return Rare, nil
	case 0:
		return None, nil
	default:
		return None, RarityError
	}
}

// Skirmish decks can't have more than 3 identical cards in them.
const (
	Common   Rarity = 3
	Uncommon Rarity = 2
	Rare     Rarity = 1
	None     Rarity = 0
)

// String returns the name of a rarity level: "Common", "Uncommon", or "Rare".
func (r *Rarity) String() string {
	switch *r {
	case 3:
		return "Common"
	case 2:
		return "Uncommon"
	case 1:
		return "Rare"
	}
	panic("Something went wrong, Rarity out of bounds.")
}

func (r *Rarity) Int() int {
	return int(*r)
}

type Faction string

const (
	Neutral   Faction = "Neutral"
	Nightmare Faction = "Nightmare"
	Troika    Faction = "Troika"
)
*/
