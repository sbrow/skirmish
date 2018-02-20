package skirmish

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	resolve int    // The resolve this card produces when in play.
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

/*func NewCard() *card {
	return &card{
		name:    "Card",
		Type:    "card_type",
		resolve: 0,
		speed:   1,
		damage:  0,
		life:    0,
		short:   "",
		long:    "",
		flavor:  "",
		// dir:        "",
	}
}*/

// Labels prints the column labels for .csv output.
func (c card) Labels() []string {
	return []string{
		"ID", "Name", "Resolve", "Speed", "Damage", "Life",
		"ShortText", "LongText", "FlavorText", "card_image",
		"Action", "Event", "Continuous", "Item",
		"show_resolve", "show_speed", "show_tough", "show_life",
	}
}

// Image returns the path to the card's illustration.
// Images for deck cards must be in a directory named after their leader.
// Images for leader and partner heroes should be in a directory names "Heroes".
//
// path = [$SK_SRC]/[folder]]/[c.Name].png
func (c *card) Images() (paths []string, err error) {
	// path[0] = fmt.Sprintf(filepath.Join(ImageDir, dir))
	// if c.Arts == 1 {
	// path = filepath.Join(ImageDir, c.Name+".png")
	// } else {
	path := filepath.Join(ImageDir, c.Name())
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.SetPrefix("[ERROR]")
		log.Print(path, " does not exist!")
		return []string{}, err
	}
	// path = filepath.Join(path, dir[ver-1].Name())
	// }

	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = errors.New(fmt.Sprint(path, " does not exist!"))
	}
	return []string{dir}, err
}

// TODO: Default image / image handling.
func (c *card) CSV() [][]string {
	str := make([]string, len(c.Labels()))
	for i, label := range c.Labels() {
		switch label {
		// case "ID":
		// str[i] += fmt.Sprintf("\"%s\"", c.Name)
		case "Name":
			str[i] += fmt.Sprintf("\"%s\"", c.Name())
		case "Resolve":
			str[i] += fmt.Sprintf("\"%+d\"", c.resolve)
		case "Speed":
			str[i] += fmt.Sprint(c.speed)
		case "Damage":
			str[i] += fmt.Sprint(c.damage)
		case "Life":
			str[i] += fmt.Sprint(c.life)
		case "ShortText":
			str[i] += fmt.Sprintf("\"%s\"", c.short)
		case "LongText":
			str[i] += fmt.Sprintf("\"%s\"", c.long)
		case "FlavorText":
			str[i] += fmt.Sprintf("\"%s\"", c.flavor)
		case "card_image":
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
		case "Action":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Action"))
		case "Event":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Event"))
		case "Continuous":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Continuous"))
		case "Item":
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
	return fmt.Sprintf(`"%s" %+d`, c.Name(), c.resolve)
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
	toughness int // The damage the card can take before being discarded (if it's a follower).
	cost      int
}

func (d *DeckCard) Cost() (string, error) {
	return fmt.Sprint(d.cost), nil
}

func (d *DeckCard) String() string {
	str := fmt.Sprintf(`"%s"`, d.Name())
	if cost, err := d.Cost(); err == nil {
		str += fmt.Sprintf(" (%s)", cost)
	}
	if d.resolve != 0 {
		str += fmt.Sprintf("%+d", d.resolve)
	}
	// str += fmt.Sprintf(" %d/%d", d.Damage, d.Toughness)
	str += fmt.Sprintf(" %d/%d", d.damage, d.life)
	return str
}

func (d *DeckCard) CSV() [][]string {
	// Add Leaders, Rarity, cost, id?, toughness,
	return [][]string{}
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
	str := n.card.String()
	if n.resolveB > 0 {
		str += fmt.Sprintf("/%+d", n.resolveB)
	}
	return str
}
