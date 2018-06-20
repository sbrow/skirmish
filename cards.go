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
	FullType() string
	SetName(string)
	Type() string
	SetType(string)
	STypes() []string
	SetSTypes([]string)
	Resolve() string
	SetResolve(string)
	Speed() int
	Faction() string
	SetSpeed(int)
	Damage() int
	Leader() string
	SetLeader(string)
	SetDamage(int)
	Life() int
	SetLife(int)
	Short() string
	SetShort(string)
	Regexp() string
	SetRegexp(string)
	Long() string
	SetLong(string)
	Flavor() string
	SetFlavor(string)
	Labels() []string
	String() string
	Images() ([]string, error)
	Bold() ([][]int, error)

	UEJSON(bool) ([]byte, error)
	CSV(bool) [][]string
	XML() ([]byte, error)
}

// Card is the base struct for DeckCards and NonDeckCards.
type card struct {
	name    string // The name of the card.
	leader  string
	ctype   string   // The card's type.
	stype   []string // The card's supertype(s).
	resolve string   // The resolve this card produces when in play.
	speed   int      // The speed the card has (if it's a character).
	damage  int      // The damage the card deals in combat (if it's a character).
	life    int      // The damage the card can take before being discarded (if it's a hero).
	short   string   // The card's basic rules text.
	long    string   // The card's reminder text.
	flavor  string   // The card's flavor (non-rules) text.
	regexp  string   // A regular expression for what characters should be bold in short.
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

func (c *card) Resolve() string {
	return fmt.Sprint(c.resolve)
}

func (c *card) Faction() string {
	return ""
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
	return c.leader
}

func (c *card) SetLeader(l string) {
	c.leader = l
}

func (c *card) SetRegexp(reg string) {
	c.regexp = reg
}

func (c *card) Regexp() string {
	return c.regexp
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
	labels := []string{
		"id", "name", "resolve", "speed", "damage", "life",
		"short", "long", "flavor", "card_image",
	}
	if len(Leaders) == 0 {
		log.Fatal("No leaders found when computing labels")
	}
	labels = append(labels, Leaders.Names()...)
	return labels
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

		}
		// fmt.Println(c.name, label,
		// strings.Contains(strings.Join(Leaders.Names(), ","), label),
		// fmt.Sprintf("\"%s\"", c.Leader()), c.Leader() == label)
		if strings.Contains(strings.Join(Leaders.Names(), ","), label) {
			str[i] += fmt.Sprint(c.Leader() == label)
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

func (c *card) FullType() string {
	if len(c.STypes()) == 0 {
		return c.ctype
	}
	return fmt.Sprintf("%s- %s", c.ctype, strings.Join(c.stype, " "))
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
	str += c.ctype
	str += fmt.Sprintf(" \"%s",
		strings.Replace(c.short, "\r\n", " ", -1))
	str += "\""
	return str
}

func (c *card) Bold() ([][]int, error) {
	reg, err := regexp.Compile(c.regexp)
	if err != nil {
		fmt.Println(c.regexp)
		return [][]int{}, err
	}
	return reg.FindAllStringIndex(c.short, -1), nil
}

type DeckCard struct {
	card
	cost   string
	rarity int
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

func (d *DeckCard) String() string {
	str := fmt.Sprintf("%dx[%s] %s", d.rarity, d.leader, d.card.String())
	return strings.Replace(str, d.name+" ", fmt.Sprintf("%s (%s)",
		d.name, d.cost), 1)
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
	labels := append(d.card.Labels(), "cost", "type", "border_normal", "action",
		"event", "continuous", "item", "show_resolve", "show_speed",
		"show_tough", "show_life", "common", "uncommon", "rare", "rare_border")
	return labels
}

func (d *DeckCard) NormalBorder() bool {
	switch {
	case d.rarity == 1:
		fallthrough
	case d.Type() == "Action":
		fallthrough
	case d.Type() == "Deck Hero":
		fallthrough
	case d.Type() == "Item":
		fallthrough
	case strings.Contains(strings.Join(d.STypes(), ","), "Continuous"):
		return false
	default:
		return true
	}
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
				log.Panic(err)
			}
		case "type":
			if len(d.stype) > 0 {
				out[1] = append(out[1], fmt.Sprintf("%s %s",
					strings.Join(d.stype, " "), d.Type()))
			} else {
				out[1] = append(out[1], d.Type())
			}
		case "action":
			fallthrough
		case "event":
			fallthrough
		case "item":
			out[1] = append(out[1], fmt.Sprint(strings.Contains(d.Type(),
				strings.Title(label))))
		case "continuous":
			out[1] = append(out[1], fmt.Sprint(strings.Contains(
				strings.Join(d.STypes(), ","), "Continuous")))
		case "show_resolve":
			out[1] = append(out[1], fmt.Sprint(d.Resolve() != "0" &&
				d.Resolve() != ""))
		case "show_speed":
			out[1] = append(out[1], fmt.Sprint(d.speed != 0))
		case "show_tough":
			out[1] = append(out[1], fmt.Sprint(strings.Contains(d.Type(), "Follower")))
		case "show_life":
			out[1] = append(out[1], fmt.Sprint(strings.Contains(d.Type(), "Hero")))
		case "border_normal":
			out[1] = append(out[1], fmt.Sprint(d.NormalBorder()))
		case "rare_border":
			out[1] = append(out[1], fmt.Sprint(d.Rarity() == "rare" &&
				!strings.Contains(strings.Join(d.STypes(), ","), "Continuous")))
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
	img_i, type_i := -1, -1
	for j, col := range out[0] {
		switch col {
		case "card_image":
			img_i = j
		case "type":
			type_i = j
		}
		if img_i != -1 && type_i != -1 {
			break
		}
	}
	if img_i == -1 {
		log.Panic("card_image not found!")
	}
	if type_i == -1 {
		log.Panic("type not found!")
	}
	tmp[0] = out[0]
	tmp[1] = out[1]
	out = tmp
	out[1][0] = fmt.Sprintf("%s_%d", d.name, 1)
	out[1][img_i] = imgs[0]
	if len(imgs) > 1 {
		out[1][type_i] = fmt.Sprintf("%s- %s", out[1][type_i],
			strings.TrimSuffix(filepath.Base(imgs[0]), ".png"))
	}
	for j := 2; j <= len(imgs); j++ {
		out[j] = make([]string, len(out[j-1]))
		copy(out[j], out[j-1])
		out[j][0] = fmt.Sprintf("%s_%d", d.name, j)
		out[j][img_i] = imgs[j-1]
		out[j][type_i] = fmt.Sprintf("%s- %s",
			strings.Split(out[j-1][type_i], "-")[0],
			strings.TrimSuffix(filepath.Base(imgs[j-1]), ".png"))
	}
	if lbls {
		return out
	}
	return out[1:]
}

func (d *DeckCard) Type() string {
	if d.ctype == "Hero" {
		return "Deck Hero"
	}
	return d.card.Type()
}

// TODO: Return error when not found.
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

// TODO: Make getters/setters for NonDeckCard
type NonDeckCard struct {
	card
	faction  string
	SpeedB   *int
	ResolveB *string
	DamageB  *int
	LifeB    *string
	ShortB   *string
	LongB    *string
	FlavorB  *string
}

func (n *NonDeckCard) Faction() string {
	return n.faction
}

func (n *NonDeckCard) SetFaction(faction string) {
	n.faction = faction
}
func (n *NonDeckCard) String() string {
	return fmt.Sprint(n.card.String(), *n.ResolveB)
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

func (n *NonDeckCard) Images() (paths []string, err error) {
	basePath := filepath.Join(ImageDir, "Heroes", n.Name())
	imgs := []string{basePath + ".png", basePath + " (Halo).png"}
	for i, img := range imgs {
		if _, err = os.Stat(img); os.IsNotExist(err) {
			imgs[i] = DefaultImage
		}
	}
	if imgs[1] == DefaultImage {
		return []string{imgs[0]}, nil
	}
	return imgs, nil
}

// SetCard makes c the DeckCard's base card.
func (n *NonDeckCard) SetCard(c Card) {
	n.card = c.Card()
}
func (n *NonDeckCard) Labels() []string {
	labels := append(n.card.Labels(), "Halo", "Troika", "Nightmares")
	return labels
}

func (n *NonDeckCard) CSV(lbls bool) [][]string {
	out := n.card.CSV(true)
	out[0] = n.Labels()
	l := n.Labels()[len(n.card.Labels()):]
	for _, label := range l {
		switch label {
		case "Halo":
			out[1] = append(out[1], "false")
		case "Troika":
			fallthrough
		case "Nightmares":
			out[1] = append(out[1], fmt.Sprint(n.Faction() == label))
		}
	}
	imgs, err := n.Images()
	if err != nil {
		log.Println(err)
	}
	tmp := make([][]string, len(imgs)+1, len(out[0]))
	tmp[0] = out[0]
	tmp[1] = out[1]
	i := -1
	for j, col := range out[0] {
		if col == "card_image" {
			i = j
			break
		}
	}
	out = tmp
	out[1][0] = n.name
	out[1][1] = fmt.Sprintf("%s- %s", n.Type(), n.name)
	out[1][2] = string(out[1][2][1])
	out[1][i] = imgs[0]
	// out[1][20] = "true"
	for j := 2; j <= len(imgs); j++ {
		out[j] = make([]string, len(out[j-1]))
		copy(out[j], out[j-1])
		out[j][0] = fmt.Sprintf("%s (Halo)", n.name)
		out[j][1] = fmt.Sprintf("%s- %s", n.Type(), n.name)
		out[j][i] = imgs[j-1]
		out[j][20] = "true"
	}
	if lbls {
		return out
	}
	return out[1:]
}
