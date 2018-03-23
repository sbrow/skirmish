package sql

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
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
	Labels() []string
	String() string
	CSV() [][]string
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

func (c card) MarshalJSON() ([]byte, error) {
	obj := CardUEJSON{}
	obj.Name = c.name
	obj.Type = "CTE_" + c.ctype
	resolve, err := strconv.Atoi(c.resolve)
	if err != nil {
		return []byte{}, err
	}
	obj.Stats = Stats{Life: c.life, Damage: c.damage, Speed: c.speed,
		Resolve: resolve, Short: c.short, Long: c.long, Flavor: c.flavor}
	obj.Abilities = make([]string, 0)
	obj.Visual = *NewVisual(c.name, "Common", 1)
	obj.SystemData = SystemData{make([]string, 0), make([]string, 0), make([]string, 0)}
	return json.Marshal(obj)
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

}

func (c *card) CSV() [][]string {
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

func (d DeckCard) MarshalJSON() ([]byte, error) {
	byt, err := d.card.MarshalJSON()
	if err != nil {
		log.Panic(err)
	}
	obj := DeckCardUEJSON{}
	err = json.Unmarshal(byt, &obj.CardUEJSON)
	if err != nil {
		log.Panic(err)
	}
	obj.CardName = d.name
	obj.Supertypes = "CTE_" + strings.Join(d.stype, "_")
	obj.Name = strings.Replace(d.name, " ", "", -1)
	obj.Leader = d.leader
	obj.Copies = d.rarity
	obj.Visual = *NewVisual(d.name, d.leader, d.rarity)
	obj.Stats.Cost = d.cost

	return json.Marshal(obj)
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
		log.Panic(err)
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
	return out
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

func (n NonDeckCard) MarshalJSON() ([]byte, error) {
	byt, err := n.card.MarshalJSON()
	if err != nil {
		log.Panic(err)
	}
	obj := NonDeckCardUEJSON{}
	err = json.Unmarshal(byt, &obj.CardUEJSON)
	if err != nil {
		log.Panic(err)
	}
	obj.Faction = "FE_" + n.faction
	if n.resolveB != nil {
		resolve, err := strconv.Atoi(*n.resolveB)
		if err != nil {
			return []byte{}, err
		}
		if n.speedB != nil {
			obj.ActiveStats.Speed = *n.speedB
		}
		if n.damageB != nil {
			obj.ActiveStats.Damage = *n.damageB
		}
		if n.lifeB != nil {
			life, err := strconv.Atoi(*n.lifeB)
			if err != nil {
				return []byte{}, err
			}
			obj.ActiveStats.Life = life
		}
		if n.shortB != nil {
			obj.ActiveStats.Short = *n.shortB
		}
		if n.longB != nil {
			obj.ActiveStats.Long = *n.longB
		}
		if n.flavorB != nil {
			obj.ActiveStats.Flavor = *n.flavorB
		}
		obj.ActiveStats.Resolve = resolve
	}
	obj.Visual.BackTexture = strings.Replace(obj.Visual.BackTexture,
		"CardBack", fmt.Sprintf("01x_%s_Halo", n.name), -1)
	mat := "MaterialInstanceConstant'/Game/Materials"
	obj.Visual.FrontMaterial = fmt.Sprintf("%s/Card%s_Inst.Card%[2]s_Inst'",
		mat, "Front")
	obj.Visual.BackMaterial = fmt.Sprintf("%s/Card%s_Inst.Card%[2]s_Inst'",
		mat, "Back")
	obj.DeckCards = fmt.Sprintf("DataTable'/Game/Data/%sDeck.%[1]sDeck'", n.name)
	return json.Marshal(obj)
}

func (d *DeckCard) Images() (paths []string, err error) {
	// Path to a subfolder, assuming the card has multiple images.
	path := filepath.Join(ImageDir, d.leader, d.Name())
	// If the card does not have a subfolder, check in the main folder for
	// an image file.
	if _, err = os.Stat(path); os.IsNotExist(err) {
		// If found, return it, if not, throw an error.
		if _, err = os.Stat(path + ".png"); os.IsNotExist(err) {
			ret, _ := d.card.Images()
			return ret, errors.New("No image found")
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
