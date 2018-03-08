package sql

import (
	"encoding/json"
	"errors"
	"fmt"
	// "io/ioutil"
	// "log"
	// "os"
	"path/filepath"
	// "regexp"
	"strings"
)

// Card is an interface shared by DeckCards and NonDeckCards.
//
// It is important to note that each Card object can represent more than one card.
// In the case of DeckCards, a Card object can have more than one unique art.
// For NonDeckCards, a Card object can hold their frontside values as well as their
// flipped values.
//
// TODO: Figure out how to handle id.
type Card interface {
	Name() string
	Type() string
	Resolve() string
	// Cost() (string, error)
	// ID() string
	Labels() []string
	String() string
	CSV() [][]string
	Images() ([]string, error)
}

// Card is the base struct for DeckCards and NonDeckCards.
type card struct {
	name    string // The name of the card.
	ctype   string // The card's type.
	resolve string // The resolve this card produces when in play.
	speed   int    // The speed the card has (if it's a character).
	damage  int    // The damage the card deals in combat (if it's a character).
	life    int    // The damage the card can take before being discarded (if it's a hero).
	short   string // The card's basic rules text.
	long    string // The card's reminder text.
	flavor  string // The card's flavor (non-rules) text.
	// BoldWords  string   // A regex matching all the words in short text that need to be bold.
	// labels     []string // Labels to use when converting to csv output.
	// dir        string   // The directory to look for the card image in.
}

func (c *card) Name() string {
	return c.name
}

func (c *card) Cost() (string, error) {
	return "", errors.New(fmt.Sprintf(`card "%s" has no cost.`))
}

func (c *card) Resolve() string {
	return fmt.Sprint(c.resolve)
}

func (c *card) Type() string {
	return c.ctype
}

// ID returns an id unique to the card.
//
// Cards with only one art will have identical ID and Name.
// Cards with more than one art will have an ID containing their name
// and which version of art they use.
func (c *card) ID(ver int) string {
	return c.name
}

// Labels prints the column labels for .csv output.
func (c card) Labels() []string {
	return []string{
		"id", "name", "resolve", "type", "speed", "damage", "life",
		"short", "long", "flavor", "card_image",
		"action", "event", "continuous", "item",
		"show_resolve", "show_speed", "show_tough", "show_life",
	}
}

// Image returns the path to the card's illustration.
// Images for deck cards must be in a directory named after their leader.
// Images for leader and partner heroes should be in a directory names "Heroes".
//
// path = [$SK_SRC]/[folder]]/[c.Name].png
func (c *card) Images() (paths []string, err error) {
	return []string{filepath.Join(ImageDir, "ImageNotFound.png")}, nil
	/*
		path := filepath.Join(ImageDir, c.Name())
		dir, err := ioutil.ReadDir(path)
		if err != nil {
			log.SetPrefix("[ERROR]")
			log.Print(path, " does not exist!")
			return []string{}, err
		}
		if _, err = os.Stat(path); os.IsNotExist(err) {
			err = errors.New(fmt.Sprint(path, " does not exist!"))
		}
		for i, file := range dir {
			paths[i] = file.Name()
		}
		return paths, err
	*/
}

// TODO: Default image / image handling.
func (c *card) CSV() [][]string {
	str := make([]string, len(c.Labels()))
	for i, label := range c.Labels() {
		switch label {
		// case "ID":
		// fallthrough
		// str[i] += fmt.Sprintf("\"%s\"", c.Name)
		case "name":
			str[i] += c.Name() //fmt.Sprintf("\"%s\"", c.Name())
		case "resolve":
			str[i] += fmt.Sprint(c.resolve)
		case "type":
			str[i] += c.ctype
		case "speed":
			str[i] += fmt.Sprint(c.speed)
		case "damage":
			str[i] += fmt.Sprint(c.damage)
		case "life":
			str[i] += fmt.Sprint(c.life)
		case "short":
			var s string
			if s = c.short; len(s) == 0 {
				s = " "
			}
			str[i] += s
		case "long":
			var s string
			if s = c.long; len(s) == 0 {
				s = " "
			}
			str[i] += s
		case "flavor":
			var s string
			if s = c.flavor; len(s) == 0 {
				s = " "
			}
			str[i] += s
		case "card_image":
			img, err := c.Images()
			if err != nil {
				panic(err)
			}
			str[i] += img[0]
			// img, err := c.Image( /*i*/ ) //Borked, need leader name
			/*if err != nil {
				pre := log.Prefix()
				log.SetPrefix("[ERROR] ")
				log.Println(err)
				log.SetPrefix(pre)
				str[i] += ""
			} else {
				str[i] += fmt.Sprintf("\"%s\"", img)
			}*/
		case "action":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Action"))
		case "event":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Event"))
		case "continuous":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Continuous"))
		case "item":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Item"))
		case "show_resolve":
			str[i] += fmt.Sprintf("%v", c.Resolve() != "0")
		case "show_speed":
			// TODO: Clumsy
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Follower") ||
				strings.Contains(c.Type(), "Hero"))
		case "show_tough":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Follower"))
		case "show_life":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Hero"))

		}
		str[i] += Delim
	}
	for i := range str {
		str[i] = strings.TrimSuffix(str[i], Delim)
	}
	return [][]string{c.Labels(), str}
}

func (c *card) JSON() ([]byte, error) {
	bytes, err := json.Marshal(c)
	return bytes, err
}

func (c *card) String() string {
	str := c.Name()
	if c.resolve != "0" {
		str += fmt.Sprintf(" %+d", c.resolve)
	}
	str += fmt.Sprintf(" %d/%d", c.damage, c.life)
	str += fmt.Sprintf(" %s \"%s", c.Type(),
		strings.Replace(c.short, "\r\n", " ", -1))
	str += "\""
	return str
}

// TODO: Test
/*func (c *Card) BoldText() ([][]int, error) {
	reg, err := regexp.Compile(c.BoldWords)
	if err != nil {
		return [][]int{}, err
	}
	return reg.FindAllStringIndex(c.ShortText, -1), nil
}
*/

type DeckCard struct {
	card
	cost   int
	rarity int
	leader string
}

func (d *DeckCard) Cost() (string, error) {
	return fmt.Sprint(d.cost), nil
}

func (d *DeckCard) String() string {
	str := d.card.String()
	return strings.Replace(str, d.name+" ", fmt.Sprintf("%s (%d)",
		d.name, d.cost), 1)
	/*	str := d.Name() //fmt.Sprintf(`"%s"`, d.Name())
		if cost, err := d.Cost(); err == nil {
			str += fmt.Sprintf(" (%s)", cost)
		}
		if d.resolve != 0 {
			str += fmt.Sprintf("%+d", d.resolve)
		}
		str += fmt.Sprintf(" %d/%d", d.damage, d.life)
		str += fmt.Sprintf(" %s \"%s", d.Type(),
			strings.Replace(d.short, "\r\n", "\\r", -1))
		str += "\""
		return str
	*/
}

func (d *DeckCard) Rarity() string {
	switch d.rarity {
	case 1:
		return "rare"
	case 2:
		return "uncommon"
	case 3:
		return "common"
	}
	return ""
}

func (d *DeckCard) Labels() []string {
	labels := append(d.card.Labels(), "cost", "border_normal",
		"common", "uncommon", "rare")
	return append(labels, Leaders...)
}

func (d *DeckCard) CSV() [][]string {
	out := d.card.CSV()
	out[0] = d.Labels()
	l := d.Labels()[len(d.card.Labels()):]
	for _, label := range l {
		fmt.Println(label)
		switch label {
		case "cost":
			cost, err := d.Cost()
			if err == nil {
				out[1] = append(out[1], cost)
			} else {
				panic(err)
			}
		case "border_normal":
			out[1] = append(out[1], "true")
		}
		// TODO: Force skip
		if strings.Contains(strings.Join(Leaders, ","), label) {
			out[1] = append(out[1], fmt.Sprint(d.leader == label))
		}
		if strings.Contains("common,uncommon,rare", label) {
			out[1] = append(out[1], fmt.Sprint(d.Rarity() == label))
		}
		// out[1] += Delim
	}
	out[1][0] = d.Name()
	// for i, elem := range out {
	// 	out[i] = strings.TrimSuffix(elem, ",")
	// }
	// Add Leaders, Rarity, cost, image, id?, toughness,
	return out
}

type NonDeckCard struct {
	card
	resolveB int
	/*
		Faction
		HaloResolve    int
		HaloSpeed      int
		HaloDamage     int
		HaloLife       int
		HaloShortText  string
		HaloLongText   string
		HaloFlavorText string
	*/
}

func (n *NonDeckCard) String() string {
	return n.card.String()
	// str := n.card.String()
	// if n.resolveB > 0 {
	// 	str += fmt.Sprintf("/%+d", n.resolveB)
	// }
	// str += fmt.Sprintf(" %d/%d", n.damage, n.life)
	// str += fmt.Sprintf(" %s \"%s", n.Type(),
	// 	strings.Replace(n.short, "\r\n", "\\r", -1))
	// str += "\""
	// return str
}
