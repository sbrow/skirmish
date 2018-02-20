package skirmish

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Load(name string) Card {
	c := &card{}
	var typ, title, short, long, flavor *string
	var resolve, speed, damage, life *int
	props := []string{"\"name\"", "\"type\"", "short", "reminder", "flavor",
		"resolve", "speed", "damage", "life"}
	str := fmt.Sprintf("select %[1]s from ( select %[1]s FROM "+
		"skirmish.deckcards as t1 UNION select %[1]s FROM skirmish.leaders "+
		"as t2 UNION select %[1]s from skirmish.guests as t3 ) as t where "+
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
	if &resolve != nil {
		c.resolve = *resolve
	}
	if &speed != nil {
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
		var tough, cost *int
		props = []string{"toughness", "cost"}
		err = Database.QueryRow(
			fmt.Sprintf("SELECT %s FROM skirmish.deckcards WHERE name='%s'",
				strings.Join(props, ", "), name)).Scan(&tough, &cost)
		if err != nil {
			panic(err)
		}
		if tough != nil {
			d.toughness = *tough
		}
		if cost != nil {
			d.cost = *cost
		}
		return d
	}
	return nil
}
