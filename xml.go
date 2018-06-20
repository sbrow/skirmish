package skirmish

import (
	"encoding/xml"
	"fmt"
)

type CardXML struct {
	XMLName xml.Name `xml:"card"`
	Name    string   `xml:"name"`
	Set     string   `xml:"set"`
	Type    string   `xml:"type"`
	PT      string   `xml:"pt,omitempty"`
	Row     uint8    `xml:"row"`
	Text    string   `xml:"text,omitempty"`
}

func (c *card) MarshalXML() ([]byte, error) {
	obj := CardXML{
		Name: c.name,
		Set:  "DKC",
		Type: c.FullType(),
		Row:  2,
		Text: fmt.Sprintf("%s\n(%s)", c.Short(), c.Long()),
	}
	if c.Damage() != 0 && c.Life() != 0 {
		obj.PT = fmt.Sprintf("%d/%d", c.Damage(), c.Life())
	}
	for _, t := range c.STypes() {
		if t == "Continuous" {
			obj.Row = 1
		}
	}
	if obj.Row == 2 && c.Type() == "Action" || c.Type() == "Event" {
		obj.Row = 3
	}
	if c.Speed() > 0 {
		obj.Text = fmt.Sprintf("%d Speed.\n%s", c.Speed(), obj.Text)
	}
	return xml.MarshalIndent(obj, "\t\t", "\t")
}

func (c *card) XML() ([]byte, error) {
	return c.MarshalXML()
}

func (c *DeckCard) MarshalXML() ([]byte, error) {
	obj := CardXML{
		Name: c.name,
		Set:  "DKC",
		Type: c.FullType(),
		Row:  2,
		Text: fmt.Sprintf("%s\n(%s)", c.Short(), c.Long()),
	}
	if c.Damage() != 0 && c.Life() != 0 {
		obj.PT = fmt.Sprintf("%d/%d", c.Damage(), c.Life())
	}
	for _, t := range c.STypes() {
		if t == "Continuous" {
			obj.Row = 1
		}
	}
	if obj.Row == 2 && c.Type() == "Action" || c.Type() == "Event" {
		obj.Row = 3
	}
	if c.Speed() > 0 {
		obj.Text = fmt.Sprintf("%d Speed.\n%s", c.Speed(), obj.Text)
	}
	return xml.MarshalIndent(obj, "\t\t", "\t")
}
