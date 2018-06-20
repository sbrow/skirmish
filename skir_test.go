package skirmish

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	want := `F:\GitLab\dreamkeepers-psd\Template009.1.psd`
	if Template != want {
		t.Fatalf("Wanted: \"%s\"\nGot: \"%s\"\n", want, Template)
	}
}

func TestImageDir(t *testing.T) {
	want := `F:\GitLab\dreamkeepers-psd\Images`
	if ImageDir != want {
		t.Fatalf("Wanted: \"%s\"\nGot: \"%s\"", want, ImageDir)
	}
}

func TestCard(t *testing.T) {
	card, err := Load("Ignite")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(card)
	data, err := card.XML()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestCards(t *testing.T) {
	data, err := CardsXML("true ORDER BY cards.Leader, name ASC")
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Create(`C:\Users\Spencer\AppData\Local\Cockatrice\Cockatrice\customsets\01.DKSkirmish.xml`)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
	fmt.Fprintln(f, string(data))
}

func CardsXML(query string) ([]byte, error) {
	cards, err := LoadMany(query)
	if err != nil {
		return []byte{}, err
	}
	var buf bytes.Buffer
	buf.WriteString("<cockatrice_carddatabase version=\"3\">\n\t<cards>\n")
	for _, card := range cards {
		data, err := card.XML()
		if err != nil {
			return buf.Bytes(), err
		}
		buf.Write(data)
		buf.WriteRune('\n')
	}
	buf.WriteString("\t</cards>\n</cockatrice_carddatabase>")
	return buf.Bytes(), nil
}
