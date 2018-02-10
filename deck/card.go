package deck

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sbrow/skirmish"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type CardInterface interface {
	ID(int) string
	Labels() string
	String() string
}

// Card hold the values shared by Deck cards and NonDeckCards.
type Card struct {
	Name       string   // The name of the card.
	Arts       int      // The number of unique arts the card has.
	Type       Type     // The card's type.
	Resolve    int      // The resolve this card produces when in play.
	Speed      int      // The speed the card has (if it's a character).
	Damage     int      // The damage the card deals in combat (if it's a character).
	Toughness  int      // The damage the card can take before being discarded (if it's a follower).
	Life       int      // The damage the card can take before being discarded (if it's a hero).
	ShortText  string   // The card's basic rules text.
	LongText   string   // The card's reminder text.
	FlavorText string   // The card's flavor (non-rules) text.
	labels     []string // Labels to use when converting to csv output.
	dir        string   // The directory to look for the card image in.
}

func NewCard() *Card {
	return &Card{
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
		dir:        "",
		labels: []string{
			"ID", "Name", "Resolve", "Speed", "Damage", "Toughness", "Life",
			"ShortText", "LongText", "FlavorText",
			"card_image",
			"Action", "Event", "Continuous", "Item",
			"show_resolve", "show_speed", "show_tough", "show_life",
		},
	}
}

// ID returns an id unique to the card.
//
// Cards with only one art will have identical ID and Name.
// Cards with more than one art will have an ID containing their name
// and which version of art they use.
func (c *Card) ID(ver int) string {
	if c.Arts > 1 {
		return fmt.Sprintf("%s_%d", c.Name, ver)
	}
	return c.Name
}

// Labels prints the column labels for .csv output.
func (c *Card) Labels() string {
	str := ""
	for _, label := range c.labels {
		str += fmt.Sprintf("%s%s", label, Delim)
	}
	return strings.ToLower(str[:len(str)-1]) + "\n"
}

// Image returns the path to the card's illustration.
// Images for deck cards must be in a directory named after their leader.
// Images for leader and partner heroes should be in a directory names "Heroes".
//
// path = [$SK_SRC]/[folder]]/[c.Name].png
func (c *Card) Image(ver int) (path string, err error) {
	path = fmt.Sprintf(filepath.Join(skirmish.ImageDir, c.dir))
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
		path = filepath.Join(path, dir[ver-1].Name())
	}
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = errors.New(fmt.Sprint(path, " does not exist!"))
	}
	return path, err
}

func (c *Card) String() string {
	ret := ""
	for i := 1; i <= c.Arts; i++ {
		str := ""
		for _, label := range c.labels {
			switch label {
			case "ID":
				str += fmt.Sprintf("\"%s\"", c.ID(i))
			case "Name":
				str += fmt.Sprintf("\"%s\"", c.Name)
			case "Resolve":
				str += fmt.Sprintf("\"%+d\"", c.Resolve)
			case "Speed":
				str += fmt.Sprint(c.Speed)
			case "Damage":
				str += fmt.Sprint(c.Damage)
			case "Toughness":
				str += fmt.Sprint(c.Toughness)
			case "Life":
				str += fmt.Sprint(c.Life)
			case "ShortText":
				str += fmt.Sprintf("\"%s\"", c.ShortText)
			case "LongText":
				str += fmt.Sprintf("\"%s\"", c.LongText)
			case "FlavorText":
				str += fmt.Sprintf("\"%s\"", c.FlavorText)
			case "card_image":
				img, err := c.Image(i) //Borked, need leader name
				if err != nil {
					pre := log.Prefix()
					log.SetPrefix("[ERROR] ")
					log.Println(err)
					log.SetPrefix(pre)
					str += ""
				} else {
					str += fmt.Sprintf("\"%s\"", img)
				}
			case "Action":
				str += fmt.Sprintf("%v", strings.Contains(string(c.Type), string(Action)))
			case "Event":
				str += fmt.Sprintf("%v", strings.Contains(string(c.Type), string(Event)))
			case "Continuous":
				str += fmt.Sprintf("%v", strings.Contains(string(c.Type), string(Continuous)))
			case "Item":
				str += fmt.Sprintf("%v", strings.Contains(string(c.Type), string(Item)))
			case "show_resolve":
				str += fmt.Sprintf("%v", c.Resolve != 0)
			case "show_speed":
				// TODO: Clumsy
				str += fmt.Sprintf("%v",
					strings.Contains(string(c.Type), "Follower") ||
						strings.Contains(string(c.Type), string(Hero)))
			case "show_tough":
				str += fmt.Sprintf("%v", strings.Contains(string(c.Type), "Follower"))
			case "show_life":
				str += fmt.Sprintf("%v", strings.Contains(string(c.Type), string(Hero)))
			}
			str += Delim
		}
		ret += strings.TrimSuffix(str, ",") + "\n"
	}
	return strings.TrimSuffix(ret, "\n")
}

func (c *Card) JSON() ([]byte, error) {
	bytes, err := json.Marshal(c)
	return bytes, err
}

type NonDeckCard struct {
	Card
	Faction
	HaloResolve    int
	HaloSpeed      int
	HaloDamage     int
	HaloLife       int
	HaloShortText  string
	HaloLongText   string
	HaloFlavorText string
}
