package sql

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func Load(name string) Card {
	c := &card{}
	var typ, title, short, long, flavor *string
	var resolve, speed, damage, life *int
	props := []string{"\"name\"", "\"type\"", "short", "reminder", "flavor",
		"resolve", "speed", "damage", "life"}
	str := fmt.Sprintf("select %[1]s from skirmish.cards where "+
		"\"name\"='%[2]s'", strings.Join(props, ", "), name)
	err := Database.QueryRow(str).Scan(&title, &typ, &short, &long, &flavor,
		&resolve, &speed, &damage, &life)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No card was found with name \"%s\"\n", name)
		return nil
	case err != nil:
		panic(err)
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
	case *typ == "Leader" || *typ == "Guest":
		c.ctype = *typ
		n := &NonDeckCard{}
		n.card = *c
		var resolveB *string
		props = []string{"resolve_b"}
		err = Database.QueryRow(
			fmt.Sprintf("SELECT %s FROM skirmish.nondeckcards WHERE name='%s'",
				strings.Join(props, ", "), name)).Scan(&resolveB)
		if err != nil {
			panic(err)
		}
		if resolveB != nil {
			if i, err := strconv.Atoi(*resolveB); err == nil {
				n.resolveB = i
			} else {
				panic(err)
			}
		}
		return n
	case typ != nil:
		c.ctype = *typ
		d := &DeckCard{}
		d.card = *c
		var cost, rarity *int
		var leader *string
		props = []string{"cost", "rarity", "leader"}
		err = Database.QueryRow(
			fmt.Sprintf("SELECT %s FROM skirmish.deckcards WHERE name='%s'",
				strings.Join(props, ", "), name)).Scan(&cost, &rarity,
			&leader)
		if err != nil {
			panic(err)
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
		return d
	}
	return nil
}

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
}

func Dump(dir string) {
	var out bytes.Buffer
	var errs bytes.Buffer

	cmd := exec.Command("pg_dump", "-U", "postgres", "-n", "skirmish", "-c", "--if-exists",
		"--column-inserts", "-f", filepath.Join(dir, "skirmish_db.sql"))
	cmd.Stdout = &out
	cmd.Stderr = &errs
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
