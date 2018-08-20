// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xml is a test package
package xml

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/skir/internal/base"
)

const DeckVersion = "1"

var CmdXML = &base.Command{
	Run:       runXML,
	UsageLine: "xml",
	Short:     "test xml output",
	Long:      ``,
}

type Card struct {
	XMLName xml.Name `xml:"card"`
	Number  int      `xml:"number,attr"`
	Name    string   `xml:"name,attr"`
}

type Zone struct {
	XMLName xml.Name `xml:"zone"`
	Name    string   `xml:"name,attr"`
	Cards   []Card
}
type DeckXML struct {
	XMLName  xml.Name `xml:"cockatrice_deck"`
	Name     string   `xml:"deckName"`
	Version  string   `xml:"version,attr"`
	Comments string   `xml:"comments"`
	Zones    []Zone   `xml:"zone"`
}

func runXML(cmd *base.Command, args []string) {
	// if err := MarshalDeck(args[0], args[1], os.Stdout); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := FactionDecks(args[0]); err != nil {
	// 	fmt.Println(err)
	// }
}

func MarshalDeck(a, b string, f *os.File) error {
	f.WriteString(xml.Header)

	enc := xml.NewEncoder(f)
	enc.Indent("", "\t")
	defer enc.Flush()

	deck := DeckXML{
		Version: DeckVersion,
		Name:    fmt.Sprintf("%s_%s", a, b),
	}

	args := []struct {
		Name string
		Cond string
	}{
		{"main", fmt.Sprintf("cards.leader~'%s|%s'", a, b)},
		{"side", fmt.Sprintf("cards.leader IS NULL AND cards.name~'%s|%s'", a, b)},
	}

	for _, arg := range args {
		condition := arg.Cond
		cards, err := skirmish.LoadMany(condition)
		if err != nil {
			panic(err)
		}
		cardData := make([]Card, len(cards))
		for i, card := range cards {
			cardData[i] = Card{Name: card.Name(), Number: card.Copies()}
			switch card.(type) {
			case *skirmish.NonDeckCard:
				val := card.(*skirmish.NonDeckCard)
				if val.ResolveB() != "" {
					cardData = append(cardData, Card{Name: card.Name() + " (Halo)", Number: 1})
				}
			}
		}
		deck.Zones = append(deck.Zones, Zone{
			Name:  arg.Name,
			Cards: cardData,
		})
	}
	return enc.Encode(deck)
}

/*
func FactionDecks(faction string) error {
	rows, err := skirmish.Query(fmt.Sprintf("SELECT name from %s", faction))
	if err != nil {
		return err
	}
	troika := make([]string, 0)
	for rows.Next() {
		troika = append(troika, "")
		err = rows.Scan(&troika[len(troika)-1])
		if err != nil {
			return err
		}
	}
	light := combin.NewSet(troika...).Combine()
	for _, combo := range light {
		err := MarshalDeck(combo[0].(string), combo[1].(string), os.Stdout)
		if err != nil {
			return err
		}
	}
	return nil
}
*/
