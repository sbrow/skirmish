package sql

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

// LoadMany selects more than one card at a time from the database.
func LoadMany(cond string) ([]Card, error) {
	out := make([]Card, 0)
	props := []string{"\"name\"", "cards.type", "cards.supertypes",
		"short", "reminder", "flavor", "resolve", "cards.speed", "cards.damage",
		"cards.life", "cards.faction, cards.cost, cards.rarity, cards.leader",
		"cards.resolve_b", "cards.life_b", "cards.speed_b", "cards.damage_b",
		"cards.short_b", "cards.long_b", "cards.flavor_b"}
	str := fmt.Sprintf("select %s from cards where %s",
		strings.Join(props, ", "), cond)
	rows, err := Database.Query(str)
	defer rows.Close()
	if err == sql.ErrNoRows {
		log.Printf("No card was found with condition \"%s\"\n", cond)
		return []Card{}, err
	} else if err != nil {
		log.Println("Error:" + str)
		return nil, err
	}
	for rows.Next() {
		var c *card
		var typ, stype, title, short, long, flavor, resolve, faction, leader,
			resolveB, lifeB, shortB, longB, flavorB *string
		var speed, damage, life, cost, rarity, speedB, damageB *int
		err := rows.Scan(&title, &typ, &stype, &short, &long,
			&flavor, &resolve, &speed, &damage, &life, &faction, &cost, &rarity,
			&leader, &resolveB, &lifeB, &speedB, &damageB, &shortB, &longB, &flavorB)
		c = &card{}
		switch {
		case err == sql.ErrNoRows:
			log.Printf("No card was found with condition \"%s\"\n", cond)
			return nil, err
		case err != nil:
			return out, err
		}
		if typ != nil {
			c.ctype = *typ
		}
		if stype != nil {
			c.stype = strings.Split(*stype, ",")
		}
		if title != nil {
			c.name = *title
		}
		if short != nil {
			c.short = *short
		}
		if long != nil {
			c.long = *long
		}
		if flavor != nil {
			c.flavor = *flavor
		}
		if resolve != nil {
			c.resolve = *resolve
		}
		if speed != nil {
			c.speed = *speed
		}
		if damage != nil {
			c.damage = *damage
		}
		if life != nil {
			c.life = *life
		}
		switch {
		case cost != nil:
			d := &DeckCard{}
			d.card = *c
			d.cost = *cost
			if rarity != nil {
				d.rarity = *rarity
			}
			if leader != nil {
				d.leader = *leader
			}
			out = append(out, d)
		case resolveB != nil:
			n := &NonDeckCard{}
			n.card = *c
			n.resolveB = resolveB
			if lifeB != nil {
				n.lifeB = lifeB
			}
			if speedB != nil {
				n.speedB = speedB
			}
			if damageB != nil {
				n.damageB = damageB
			}
			if shortB != nil {
				n.shortB = shortB
			}
			if longB != nil {
				n.longB = longB
			}
			if flavorB != nil {
				n.flavorB = flavorB
			}
			out = append(out, n)
		default:
			out = append(out, c)
		}
	}
	return out, nil
}

// Load Selects a card from the database given it's name, and returns in
// a struct of the appropriate card type.
func Load(name string) (Card, error) {
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
}
