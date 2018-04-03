package skirmish

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
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
	Card() card
	SetName(string)
	Type() string
	SetType(string)
	STypes() []string
	SetSTypes([]string)
	Resolve() string
	SetResolve(string)
	Speed() int
	SetSpeed(int)
	Damage() int
	Leader() string
	SetDamage(int)
	Life() int
	SetLife(int)
	Short() string
	SetShort(string)
	Long() string
	SetLong(string)
	Flavor() string
	SetFlavor(string)
	UEJSON(bool) ([]byte, error)
	Labels() []string
	String() string
	CSV(bool) [][]string
	Images() ([]string, error)
}

// Card is the base struct for DeckCards and NonDeckCards.
type card struct {
	name    string   // The name of the card.
	ctype   string   // The card's type.
	stype   []string // The card's supertype(s).
	resolve string   // The resolve this card produces when in play.
	speed   int      // The speed the card has (if it's a character).
	damage  int      // The damage the card deals in combat (if it's a character).
	life    int      // The damage the card can take before being discarded (if it's a hero).
	short   string   // The card's basic rules text.
	long    string   // The card's reminder text.
	flavor  string   // The card's flavor (non-rules) text.
}

func NewCard() Card {
	return &card{}
}

func (c *card) Name() string {
	return c.name
}

func (c *card) SetName(name string) {
	c.name = name
}

func (c *card) Card() card {
	return *c
}

// func (c *card) Cost() (string, error) {
// 	return "", errors.New(fmt.Sprintf(`card "%s" has no cost.`))
// }

func (c *card) Resolve() string {
	return fmt.Sprint(c.resolve)
}

func (c *card) SetResolve(r string) {
	m, err := regexp.Match(`[+\-][1-9]`, []byte(r))
	if err != nil {
		log.Panic(err)
	}
	if m {
		c.resolve = r
	}
}

func (c *card) Speed() int {
	return c.speed
}

func (c *card) SetSpeed(s int) {
	c.speed = s
}

func (c *card) Damage() int {
	return c.damage
}

func (c *card) SetDamage(d int) {
	c.damage = d
}

func (c *card) Life() int {
	return c.life
}

func (c *card) SetLife(d int) {
	c.life = d
}

func (c *card) Short() string {
	return c.short
}

func (c *card) SetShort(s string) {
	c.short = s
}

func (c *card) Long() string {
	return c.long
}

func (c *card) SetLong(s string) {
	c.long = s
}

func (c *card) Flavor() string {
	return c.flavor
}

func (c *card) SetFlavor(s string) {
	c.flavor = s
}

func (c *card) Type() string {
	return c.ctype
}

func (c *card) SetType(t string) {
	c.ctype = t
}

func (c *card) STypes() []string {
	return c.stype
}

func (c *card) SetSTypes(t []string) {
	c.stype = t
}

func (c *card) Leader() string {
	return ""
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
}

func (c *card) CSV(lbls bool) [][]string {
	str := make([]string, len(c.Labels()))
	for i, label := range c.Labels() {
		switch label {
		case "name":
			str[i] += c.Name()
		case "resolve":
			if c.Resolve() == "" {
				str[i] += fmt.Sprint("0")
			} else {
				str[i] += fmt.Sprint(c.Resolve())
			}
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
				log.Panic(err)
			}
			str[i] += img[0]
		case "action":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Action"))
		case "event":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Event"))
		case "continuous":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Continuous"))
		case "item":
			str[i] += fmt.Sprintf("%v", strings.Contains(c.Type(), "Item"))
		case "show_resolve":
			str[i] += fmt.Sprintf("%v", c.Resolve() != "0" && c.Resolve() != "")
		case "show_speed":
			str[i] += fmt.Sprintf("%v", c.speed != 0)
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
	if lbls {
		return [][]string{c.Labels(), str}
	}
	return [][]string{str}
}

func (c *card) JSON() ([]byte, error) {
	bytes, err := json.Marshal(c)
	return bytes, err
}

func (c *card) String() string {
	str := c.Name()
	if c.resolve != "0" {
		str += fmt.Sprintf(" %s ", c.resolve)
	}
	if c.speed > 1 {
		str += fmt.Sprintf("%d/", c.speed)
	}
	str += fmt.Sprintf("%d/%d ", c.damage, c.life)
	if len(c.stype) > 0 {
		for _, elem := range c.stype {
			str += fmt.Sprintf("%s ", elem)
		}
	}
	str += fmt.Sprintf("%s", c.ctype)
	str += fmt.Sprintf(" \"%s",
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
	cost   string
	rarity int
	leader string
}

func NewDeckCard() *DeckCard {
	return &DeckCard{}
}

// SetCard makes c the DeckCard's base card.
func (d *DeckCard) SetCard(c Card) {
	d.card = c.Card()
}

func (d *DeckCard) Cost() (string, error) {
	return fmt.Sprint(d.cost), nil
}

func (d *DeckCard) SetCost(c string) {
	d.cost = c
}

func (d *DeckCard) Leader() string {
	return d.leader
}

func (d *DeckCard) SetLeader(l string) {
	d.leader = l
}

func (d *DeckCard) String() string {
	str := fmt.Sprintf("%dx[%s] %s", d.rarity, d.leader, d.card.String())
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

func (d *DeckCard) SetRarity(r int) {
	d.rarity = r
}

func (d *DeckCard) Labels() []string {
	labels := append(d.card.Labels(), "cost", "border_normal",
		"common", "uncommon", "rare")
	if len(Leaders) == 0 {
		log.Fatal("No leaders found when computing labels")
	}
	labels = append(labels, Leaders...)
	return labels
}

func (d *DeckCard) CSV(lbls bool) [][]string {
	out := d.card.CSV(true)
	out[0] = d.Labels()
	l := d.Labels()[len(d.card.Labels()):]
	for _, label := range l {
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
	}
	imgs, err := d.Images()
	if err != nil {
		log.Println(err)
	}
	tmp := make([][]string, len(imgs)+1, len(out[0]))
	tmp[0] = out[0]
	tmp[1] = out[1]
	out = tmp
	out[1][0] = fmt.Sprintf("%s_%d", d.name, 1)
	out[1][10] = imgs[0]
	for i := 2; i <= len(imgs); i++ {
		out[i] = make([]string, len(out[i-1]))
		copy(out[i], out[i-1])
		out[i][0] = fmt.Sprintf("%s_%d", d.name, i)
		out[i][10] = imgs[i-1]
	}
	if lbls {
		return out
	}
	return out[1:]
}

type NonDeckCard struct {
	card
	faction  string
	speedB   *int
	resolveB *string
	damageB  *int
	lifeB    *string
	shortB   *string
	longB    *string
	flavorB  *string
}

func (n *NonDeckCard) String() string {
	return fmt.Sprint(n.card.String(), *n.resolveB)
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

func (d *DeckCard) Images() (paths []string, err error) {
	// Path to a subfolder, assuming the card has multiple images.
	path := filepath.Join(ImageDir, d.leader, d.Name())
	// If the card does not have a subfolder, check in the main folder for
	// an image file.
	if _, err = os.Stat(path); os.IsNotExist(err) {
		// If found, return it, if not, throw an error.
		if _, err = os.Stat(path + ".png"); os.IsNotExist(err) {
			return []string{DefaultImage},
				errors.New(fmt.Sprintf(`No image found for card '%s'`, d.name))
		}
		return []string{path + ".png"}, nil
	}
	dir, err := ioutil.ReadDir(path + "\\")
	if err != nil {
		return nil, err
	}
	paths = make([]string, len(dir))
	for i, file := range dir {
		paths[i] = filepath.Join(path, file.Name())
	}
	return paths, err
}
