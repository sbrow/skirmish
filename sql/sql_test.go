package sql

import (
	"encoding/csv"
	"fmt"
	sk "github.com/sbrow/skirmish"
	"log"
	"os"
	"testing"
)

func TestSimple(t *testing.T) {
	var text *string
	err := sk.DB.QueryRow(
		"SELECT type FROM public.all_cards WHERE name='Anger'").Scan(&text)
	if err != nil {
		panic(err)
	}
	fmt.Println(*text)
}

func TestLoad(t *testing.T) {
	c, err := Load("Indigo")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n\n", c)
	fmt.Println(c)
}

func TestUEJSON(t *testing.T) {
	card, err := Load("Bast")
	if err != nil {
		t.Fatal(err)
	}
	j, err := card.UEJSON(true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
	card, err = Load("Tendril")
	if err != nil {
		t.Fatal(err)
	}
	j, err = card.UEJSON(true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
}

func TestCSV(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping CSV test")
	}
	// c, err := Load("Bushwack Squad")
	c, err := Load("Anger")
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.Create("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	w.WriteAll(c.CSV(true))
}

func TestQuery(t *testing.T) {
	var name, typ *string
	query := `SELECT name, type FROM skirmish.deckcards WHERE "name"=$1`
	sk.DB.QueryRow(query, "Anger").Scan(&name, &typ)
	if name == nil || typ == nil {
		log.Fatal("Noop!")
	}
}

func TestSQL(t *testing.T) {
	Recover(sk.DataDir)
	Dump(sk.DataDir)
}

func TestLoadMany(t *testing.T) {
	c, err := LoadMany("cards.leader='Vi'")
	fmt.Println(err)
	for _, card := range c {
		fmt.Printf("%#v\n\n", card)
	}
}

func TestGenData(t *testing.T) {
	GenData()
}
