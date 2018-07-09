package skirmish

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type DeckCard struct {
	card
	cost   string
	copies int
}

func NewDeckCard() *DeckCard {
	return &DeckCard{}
}

// Copies returns the number of copies that appear in this card's deck.
func (d *DeckCard) Copies() int {
	return d.copies
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
	cost, _ := d.Cost() // d.Cost can't return an error.
	s := fmt.Sprintf("%s %dx[%s]", d.card.String(), d.Copies(), d.Leader())
	old := d.Name()
	new := fmt.Sprintf("%s (%s)", d.Name(), cost)
	if d.Resolve() != "" {
		old += " "
	}
	return strings.Replace(s, old, new, 1)
}

func (d *DeckCard) Rarity() string {
	switch d.copies {
	case 1:
		return "rare"
	case 2:
		return "uncommon"
	case 3:
		return "common"
	}
	return ""
}

func (d *DeckCard) SetRarity(r *int) {
	if r != nil {
		d.copies = *r
	}
}

func (d *DeckCard) Labels() []string {
	labels := append(d.card.Labels(), "cost", "type", "border_normal", "action",
		"event", "continuous", "item", "show_resolve", "show_speed",
		"show_tough", "show_life", "common", "uncommon", "rare", "rare_border")
	return labels
}

func (d *DeckCard) NormalBorder() bool {
	switch {
	case d.copies == 1:
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
			out[1] = append(out[1], fmt.Sprint(d.Speed() != 0))
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

func (d *DeckCard) Images() (paths []string, err error) {
	// Path to a subfolder, assuming the card has multiple images.
	path := filepath.Join(ImageDir, d.leader, d.Name())
	// If the card does not have a subfolder, check in the main folder for
	// an image file.
	if _, err = os.Stat(path); os.IsNotExist(err) {
		// If found, return it, if not, throw an error.
		if _, err = os.Stat(path + ".png"); os.IsNotExist(err) {
			return []string{DefaultImage}, errors.New(fmt.Sprintf(`No image found for card '%s'`, d.name))
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
