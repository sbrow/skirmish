// Package deck contains code for creating Photoshop data sets from json files.
//
// TODO: Add support for heroes.
package deck

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "os"
	"reflect"
	"strings"
	"sync"
)

// Deck represents a skirmish deck, with leader and deck cards.
type Deck struct {
	Leader    *NonDeckCard
	DeckCards []DeckCard
	labels    []string
}

// New takes an input file and creates a Deck from the data.
// Input must be in JSON format and have a ".json" extension.
func New(path string, leader *NonDeckCard) (d *Deck) {
	d = &Deck{Leader: leader}
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(contents, &d.DeckCards)
	if err != nil {
		panic(err)
	}
	d.labels = []string{
		"ID", "Name", "Cost", "Type", "Resolve",
		"Speed", "Damage", "Toughness", "Life",
		"ShortText", "LongText", "FlavorText",
		"card_image",
		"Common", "Uncommon", "Rare",
		"Action", "Event", "Continuous", "Item",
		"Bast", "Igrath", "Lilith", "Vi",
		"Ravat", "Scuttler", "Tendril", "Wisp",
		"Scinter", "Tinsel",
		"show_resolve", "show_speed", "show_tough", "show_life",
		"border_normal",
	}
	return d
}

func (d *Deck) String() string {
	var wg sync.WaitGroup
	wg.Add(len(d.DeckCards))
	out := make([]string, 10) // TODO: Bad style
	for i := range d.DeckCards {
		go func(i int, out []string) {
			defer wg.Done()
			card := d.DeckCards[i]
			c := reflect.ValueOf(card)
			for v := 1; v <= int(card.Arts); v++ {
				str := ""
				for _, label := range d.labels {
					if (c.FieldByName(label) != reflect.Value{}) {
						switch c.FieldByName(label).Interface().(type) {
						case Type:
							str += fmt.Sprintf("\"%v\"", c.FieldByName(label))
						case string:
							str += fmt.Sprintf("\"%v\"", c.FieldByName(label))
						case int:
							switch label {
							case "Resolve":
								str += fmt.Sprintf("\"%+d\"", card.Resolve)
							case "Cost":
								if card.Cost != "X" {
									str += fmt.Sprintf("\"%s\"", card.Cost)
								} else {
									str += fmt.Sprintf("%d", card.Cost)
								}
							default:
								str += fmt.Sprintf("%d", c.FieldByName(label))
							}
						}
					} else {
						switch label {
						case "ID":
							str += card.ID(v)
						case "card_image":
							img, err := card.Image(d.Leader.Name, v)
							if err != nil {
								log.SetPrefix("[ERROR]")
								log.Println(err)
								log.SetPrefix("")

							} else {
								str += fmt.Sprintf("\"%s\"", img)
							}
						case "Common":
							fallthrough
						case "Uncommon":
							fallthrough
						case "Rare":
							str += fmt.Sprintf("%v", label == card.Rarity.String())
						case "Action":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(Action)))
						case "Event":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(Event)))
						case "Continuous":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(Continuous)))
						case "Item":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(Item)))
						case "Bast":
							fallthrough
						case "Igrath":
							fallthrough
						case "Lilith":
							fallthrough
						case "Vi":
							fallthrough
						case "Scinter":
							fallthrough
						case "Ravat":
							fallthrough
						case "Scuttler":
							fallthrough
						case "Tendril":
							fallthrough
						case "Wisp":
							fallthrough
						case "Tinsel":
							str += fmt.Sprintf("%v", d.Leader.Name == label)
						case "show_resolve":
							str += fmt.Sprintf("%v", card.Resolve != 0)
						case "show_speed":
							// TODO: Clumsy
							str += fmt.Sprintf("%v",
								strings.Contains(string(card.Type), "Follower") ||
									strings.Contains(string(card.Type), string(Hero)))
						case "show_tough":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), "Follower"))
						case "show_life":
							str += fmt.Sprintf("%v", strings.Contains(string(card.Type), string(Hero)))
						case "border_normal":
							str += fmt.Sprintf("%v", card.DefaultBorder())
						}
					}
					str += Delim
				}
				// out[i] += str[:len(str)-1] + "\n"
				out[i] += strings.TrimSuffix(str, ",") + "\n"
			}
		}(i, out)
	}
	wg.Wait()

	ret := ""
	for _, line := range out {
		ret += fmt.Sprintf("%s", line)
	}
	return strings.TrimSuffix(ret, "\n")
}

func (c *DeckCard) checkRarityString(r Rarity) bool {
	return fmt.Sprintf("%d") == fmt.Sprintf("%d", r)
}

// Labels prints the column labels for .csv output.
func (d *Deck) Labels() string {
	str := ""
	for _, label := range d.labels {
		str += fmt.Sprintf("%s%s", label, Delim)
	}
	return strings.ToLower(str[:len(str)-1]) + "\n"
}
