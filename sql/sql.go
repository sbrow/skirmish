package sql

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

// Load Selects a card from the database given it's name, and returns in
// a struct of the appropriate card type.
func Load(name string) (Card, error) {
	c := &card{}
	var typ, stype, title, short, long, flavor, resolve, faction *string
	var speed, damage, life *int
	props := []string{"\"name\"", "all_cards.type", "all_cards.supertypes",
		"short", "reminder", "flavor", "resolve", "speed", "damage", "life",
		"all_cards.faction"}
	str := fmt.Sprintf("select %[1]s from all_cards where "+
		"\"name\"='%[2]s'", strings.Join(props, ", "), name)
	err := Database.QueryRow(str).Scan(&title, &typ, &stype, &short, &long, &flavor,
		&resolve, &speed, &damage, &life, &faction)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No card was found with name \"%s\"\n", name)
		return nil, err
	case err != nil:
		panic(err)
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
	// TODO: Change to pull from database
	case *typ == "Leader" || *typ == "Partner":
		c.ctype = *typ
		n := &NonDeckCard{}
		n.card = *c
		var speedB, damageB *int
		var resolveB, lifeB, shortB, longB, flavorB *string
		props = []string{"resolve_b", "speed_b", "damage_b", "life_b",
			"short_b", "reminder_b", "flavor_b"}
		err = Database.QueryRow(
			fmt.Sprintf("SELECT %s FROM leaders WHERE name='%s'",
				strings.Join(props, ", "), name)).Scan(&resolveB, &speedB,
			&damageB, &lifeB, &shortB, &longB, &flavorB)
		if err != nil {
			return n, err
		}
		n.resolveB = resolveB
		n.shortB = shortB
		n.longB = longB
		n.flavorB = flavorB
		n.speedB = speedB
		n.damageB = damageB
		n.lifeB = lifeB
		if faction != nil {
			fmt.Println(*faction)
			n.faction = *faction
		}
		return n, nil
	case typ != nil:
		c.ctype = *typ
		d := &DeckCard{}
		d.card = *c
		var cost, rarity *int
		var leader *string
		props = []string{"resolve_cost", "copies", "deck_cards.leader"}
		err = Database.QueryRow(
			fmt.Sprintf("SELECT %s FROM deck_cards WHERE name='%s'",
				strings.Join(props, ", "), name)).Scan(&cost, &rarity,
			&leader)
		if err != nil {
			return d, err
		}
		if cost != nil {
			d.cost = *cost
		}
		if rarity != nil {
			d.rarity = *rarity
		}
		if leader != nil {
			d.leader = *leader
		}
		return d, nil
	}
	return nil, errors.New("Something went wrong.")
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
