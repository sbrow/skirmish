package sql

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	sk "github.com/sbrow/skirmish"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func init() {
	if sk.DB == nil {
		log.Fatal("No database")
	}
	if len(sk.Leaders) == 0 {
		// log.Fatal("No Leaders")
	}
}

// LoadMany selects more than one card at a time from the database.
func LoadMany(cond string) ([]sk.Card, error) {
	out := make([]sk.Card, 0)
	props := []string{"\"name\"", "cards.type", "cards.supertypes",
		"cards.short", "cards.long", "flavor", "resolve", "cards.speed", "cards.damage",
		"cards.life", "cards.faction, cards.cost, cards.rarity, cards.leader",
		"cards.resolve_b", "cards.life_b", "cards.speed_b", "cards.damage_b",
		"cards.short_b", "cards.long_b", "cards.flavor_b, cards.regexp"}
	str := fmt.Sprintf("select %s from cards where %s",
		strings.Join(props, ", "), cond)
	rows, err := sk.DB.Query(str)
	defer rows.Close()
	if err == sql.ErrNoRows {
		log.Printf("No card was found with condition \"%s\"\n", cond)
		return []sk.Card{}, err
	} else if err != nil {
		log.Println("Error:" + str)
		return nil, err
	}
	for rows.Next() {
		var typ, stype, title, short, long, flavor, resolve, faction, leader,
			resolveB, lifeB, shortB, longB, flavorB, cost, regexp *string
		var speed, damage, life, rarity, speedB, damageB *int
		err := rows.Scan(&title, &typ, &stype, &short, &long,
			&flavor, &resolve, &speed, &damage, &life, &faction, &cost, &rarity,
			&leader, &resolveB, &lifeB, &speedB, &damageB, &shortB, &longB, &flavorB,
			&regexp)
		var c sk.Card
		c = sk.NewCard()
		switch {
		case err == sql.ErrNoRows:
			log.Printf("No card was found with condition \"%s\"\n", cond)
			return nil, err
		case err != nil:
			return out, err
		}
		if typ != nil {
			c.SetType(*typ)
		}
		if stype != nil {
			c.SetSTypes(strings.Split(*stype, ","))
		}
		if title != nil {
			c.SetName(*title)
		}
		if short != nil {
			c.SetShort(*short)
		}
		if long != nil {
			c.SetLong(*long)
		}
		if flavor != nil {
			c.SetFlavor(*flavor)
		}
		if resolve != nil {
			c.SetResolve(*resolve)
		}
		if speed != nil {
			c.SetSpeed(*speed)
		}
		if damage != nil {
			c.SetDamage(*damage)
		}
		if life != nil {
			c.SetLife(*life)
		}
		if regexp != nil {
			c.SetRegexp(*regexp)
		}
		switch {
		case cost != nil:
			d := &sk.DeckCard{}
			d.SetCard(c)
			d.SetCost(*cost)
			if rarity != nil {
				d.SetRarity(*rarity)
			}
			if leader != nil {
				d.SetLeader(*leader)
			}
			out = append(out, d)
		case *typ == "Leader":
			n := &sk.NonDeckCard{}
			c.SetLeader(*title)
			n.SetCard(c)
			n.ResolveB = resolveB
			if lifeB != nil {
				n.LifeB = lifeB
			}
			if speedB != nil {
				n.SpeedB = speedB
			}
			if damageB != nil {
				n.DamageB = damageB
			}
			if shortB != nil {
				n.ShortB = shortB
			}
			if longB != nil {
				n.LongB = longB
			}
			if flavorB != nil {
				n.FlavorB = flavorB
			}
			n.SetFaction(*faction)
			out = append(out, n)
		default:
			out = append(out, c)
		}
	}
	return out, nil
}

// Load Selects a card from the database given it's name, and returns in
// a struct of the appropriate card type.
func Load(name string) (sk.Card, error) {
	cards, err := LoadMany(fmt.Sprintf("name='%s'", name))
	return cards[0], err
}

// Recover runs pg_recover, loading database data from a SQL file.
func Recover(dir string) {
	var out bytes.Buffer
	var errs bytes.Buffer

	cmd := exec.Command("psql", "-U", "postgres", "-f", filepath.Join(dir, "skirmish_db.sql"))
	cmd.Stdout = &out
	cmd.Stderr = &errs
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out.Bytes()))
	fmt.Println(string(errs.Bytes()))
}

// Dump runs pg_dump, saving the contents of the database to a SQL file.
func Dump(dir string) {
	var out bytes.Buffer
	var errs bytes.Buffer

	cmd := exec.Command("pg_dump", "-U", "postgres", "-n", "skirmish", "-n", "public",
		"-c", "--if-exists", "--column-inserts", "-f", filepath.Join(dir, "skirmish_db.sql"))
	cmd.Stdout = &out
	cmd.Stderr = &errs
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	if len(out.Bytes()) > 0 {
		fmt.Println(string(out.Bytes()))
	}
	if len(errs.Bytes()) > 0 {
		fmt.Println(string(errs.Bytes()))
	}
}

// GenData creates a dataset file for Photoshop to load from.
func GenData() {
	genDataSet("deckcards", "cards.Leader IS NOT NULL")
	genDataSet("nondeckcards", "cards.Leader IS NULL")
}

func genDataSet(name, query string) {
	log.SetPrefix(fmt.Sprintf("[%s] ", name))
	log.Println(`Generating Dataset`)
	cards, err := LoadMany(query)
	if err != nil {
		log.Panic(err)
	}
	dat := [][]string{cards[0].Labels()}
	for _, card := range cards {
		dat = append(dat, card.CSV(false)...)
	}
	path := filepath.Join(sk.DataDir, fmt.Sprintf("%s.csv", name))
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(dat)
	log.Println(path, "generated!")
}
