package skirmish

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	// PSQL Driver.
	_ "github.com/lib/pq"
)

// The database to retrieve card info from.
var DB *sql.DB

// Connect connects to the given database with options.
func Connect(user, dbname, sslmode string) error {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s",
		user, dbname, sslmode)
	var err error
	DB, err = sql.Open(dbname, connStr)
	return err
}

// LoadMany queries the database for all cards that match the given condition
// and returns them as a slice of Card objects.
func LoadMany(cond string) ([]Card, error) {
	if DB == nil {
		// TODO(sbrow): Fix
		if err := Connect("postgres", "postgres", "disable"); err != nil {
			return []Card{}, err
		}
	}
	out := make([]Card, 0)
	props := []string{"\"name\"", "cards.type", "cards.supertypes",
		"cards.short", "cards.long", "flavor", "resolve", "cards.speed", "cards.damage",
		"cards.life", "cards.faction, cards.cost, cards.rarity, cards.leader",
		"cards.resolve_b", "cards.life_b", "cards.speed_b", "cards.damage_b",
		"cards.short_b", "cards.long_b", "cards.flavor_b, cards.regexp"}
	str := fmt.Sprintf("select %s from cards where %s ORDER BY name ASC",
		strings.Join(props, ", "), cond)
	rows, err := DB.Query(str)
	defer rows.Close()
	if err == sql.ErrNoRows {
		log.Printf("No card was found with condition \"%s\"\n", cond)
		return []Card{}, err
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
		var c Card
		c = NewCard()
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
			d := &DeckCard{}
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
			n := &NonDeckCard{}
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

// Load queries the database for a card with the given name, and returns the
// results as a Card object.
func Load(name string) (Card, error) {
	cards, err := LoadMany(fmt.Sprintf("name='%s'", strings.Replace(name, "'", "''", -1)))
	if len(cards) > 0 {
		return cards[0], err
	}
	return nil, errors.New("No card found with name " + name + ", check your spelling.")
}

// Recover runs pg_recover on the database, loading data from the SQL file in the given directory.
func Recover(dir string) {
	var out, errs bytes.Buffer

	cmd := exec.Command("psql", "-U", "postgres", "-f", filepath.Join(dir, "skirmish_db.sql"))
	cmd.Stdout = &out
	cmd.Stderr = &errs
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String(), "\n", errs.String())
}

// Dump runs pg_dump, saving the contents of the database to a SQL file in the given directory.
func Dump(dir string) {
	var out, errs bytes.Buffer

	cmd := exec.Command("pg_dump", "-U", "postgres", "-n", "skirmish", "-n", "public",
		"-c", "--if-exists", "--column-inserts", "-f", filepath.Join(dir, "skirmish_db.sql"))
	cmd.Stdout = &out
	cmd.Stderr = &errs
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
	if len(out.String()) > 0 {
		fmt.Println(out.String())
	}
	if len(errs.String()) > 0 {
		fmt.Println(errs.String())
	}
}
