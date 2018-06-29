package skirmish

import (
	"encoding/json"
	"fmt"
	"log"
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
// TODO(sbrow): Figure out how to handle id.
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
	UEJSON(ident bool) ([]byte, error)
	CSV(bool) [][]string
	XML() ([]byte, error)
}

// Card is the base struct for DeckCards and NonDeckCards.
type card struct {
	name    string   // The name of the card.
	leader  string   // The name of the card's leader.
	ctype   string   // The card's type.
	stype   []string // The card's supertype(s).
	resolve string   // The resolve this card produces when in play.
	stats   stats    // The card's speed, life, and damage, if applicable.
	short   string   // The card's basic rules text.
	long    string   // The card's reminder text.
	flavor  string   // The card's flavor (non-rules) text.
	regexp  string   // A regular expression for what characters should be bold in short.
}

type stats struct {
	speed  int
	damage int
	life   int
}

func (s stats) String() string {
	reg := regexp.MustCompile(`^[0-1]\/(0\/)*0?`)
	str := fmt.Sprintf("%d/%d/%d", s.speed, s.damage, s.life)
	return reg.ReplaceAllString(str, "")
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
	return c.stats.speed
}

func (c *card) SetSpeed(s int) {
	c.stats.speed = s
}

func (c *card) Damage() int {
	return c.stats.damage
}

func (c *card) SetDamage(d int) {
	c.stats.damage = d
}

func (c *card) Life() int {
	return c.stats.life
}

func (c *card) SetLife(l int) {
	c.stats.life = l
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
	labels = append(labels, Leaders.names()...)
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

func (c *card) CSV(labels bool) [][]string {
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
			str[i] += fmt.Sprint(c.Speed())
		case "damage":
			str[i] += fmt.Sprint(c.Damage())
		case "life":
			str[i] += fmt.Sprint(c.Life())
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
		if strings.Contains(strings.Join(Leaders.names(), ","), label) {
			str[i] += fmt.Sprint(c.Leader() == label)
		}
		str[i] += Delim
	}
	for i := range str {
		str[i] = strings.TrimSuffix(str[i], Delim)
	}
	if labels {
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
	return noSpaces(fmt.Sprintf("%s %s %s %s \"%s\"", c.Name(), c.Resolve(),
		c.stats.String(), c.FullType(), pruneNewLines(c.short)))
}

func noSpaces(s string) string {
	reg := regexp.MustCompile(`\s\s+`)
	return reg.ReplaceAllString(s, " ")
}
func pruneNewLines(s string) string {
	reg := regexp.MustCompile("[\r\n]+")
	return reg.ReplaceAllString(s, " ")
}
func (c *card) Bold() ([][]int, error) {
	reg, err := regexp.Compile(c.regexp)
	if err != nil {
		fmt.Println(c.regexp)
		return [][]int{}, err
	}
	return reg.FindAllStringIndex(c.short, -1), nil
}
