package skirmish

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

var cols []string

func createDB() {
	err := Connect(LocalDB.DBArgs())
	if err != nil {
		log.Fatal(err)
	}
	_, err = Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
	if err != nil {
		log.Fatal(err)
	}
	_, err = Exec("DROP TABLE IF EXISTS cards CASCADE")
	if err != nil {
		log.Fatal(err)
	}
	cards := `CREATE TABLE cards (`
	for i, c := range cols {
		cards += fmt.Sprintf("\n\t%s TEXT", c)
		if i+1 != len(cols) {
			cards += ","
		}
	}
	cards += "\n)"
	Exec(cards)
}
func insertRecords(records ...Card) {
	query := `INSERT INTO cards VALUES`
	for i, r := range records {
		superTypes := "NULL"
		if len(r.STypes()) > 0 {
			superTypes = fmt.Sprintf("'%s'", strings.Join(r.STypes(), Delim))
		}
		faction := "NULL"
		if r.Faction() != "" {
			faction = fmt.Sprintf("'%s'", r.Faction())
		}
		sql := fmt.Sprintf(`('%s', '%s', %s, '%s', '%s', '%s', '%s', '%d', '%d', '%d', %s`,
			r.Name(), r.Type(), superTypes, r.Short(), strings.Replace(r.Long(), "'", "''", -1),
			r.Flavor(), r.Resolve(), r.Speed(), r.Damage(), r.Life(), faction)
		switch r.(type) {
		case *DeckCard:
			val := r.(*DeckCard)
			cost, _ := val.Cost()
			sql += fmt.Sprintf(", '%s', '%d', '%s'", cost, val.Copies(), val.Leader())
			sql += fmt.Sprintf(",NULL, NULL, NULL, NULL, NULL, NULL, NULL")
		case *NonDeckCard:
			val := r.(*NonDeckCard)
			sql += fmt.Sprintf(", NULL, NULL, NULL, '%s', '%s', '%d', '%d', '%s', %s, %s",
				*val.resolveB, val.LifeB(), val.SpeedB(), val.DamageB(), *val.shortB, "NULL", "NULL")
		case *card:
			val := r.(*card)
			sql += fmt.Sprintf(", NULL, NULL, '%s'", val.Leader())
			sql += fmt.Sprintf(",NULL, NULL, NULL, NULL, NULL, NULL, NULL")
		}
		query += fmt.Sprintf("\n%s, '%s')", sql, r.Regexp())
		if i+1 != len(records) {
			query += ","
		}
	}
	Exec(query)
}

func TestLoad(t *testing.T) {
	Ignite := &DeckCard{
		card: card{
			name:       "Ignite",
			leader:     "Bast",
			cardType:   "Action",
			superTypes: []string{"Channeled"},
			resolve:    "",
			stats:      stats{},
			short:      "Deal 3 to a hero.",
			long: "Ignite can be played on leaders, partners, or deck heroes." +
				"\rChanneled cards must be played with their leader's resolve.",
			flavor: "There are two kinds of things in this world-" +
				" things that can catch fire, and things that are on fire.",
			regexp: `(3)|(Deal)|(hero.)`,
		},
		cost:   "1",
		copies: 3,
	}
	Bast := &NonDeckCard{
		card: card{
			name:       "Bast",
			leader:     "",
			cardType:   "Hero",
			superTypes: []string{"Leader"},
			resolve:    "+2",
			stats:      stats{speed: 1, damage: 2, life: 14},
			short:      "Uncontested- +2/+0.",
			long:       "A Lane is uncontested if it is not contested.",
			flavor:     "Sometimes loyalty means not asking questions.",
			regexp:     `(\+2/\+0.)|(Uncontested-)`,
		},
		faction: "Troika",
	}
	Basic := &card{
		name:     "John Doe",
		leader:   "Blarg",
		cardType: "Action",
		short:    "Do a thing.",
		long:     "You can do any kind of thing.",
		flavor:   "Helps to have a map.",
	}
	speed := 1
	Bast.SetSpeedB(&speed)
	resolve := "+2"
	Bast.SetResolveB(&resolve)
	dam := 1
	Bast.SetDamageB(&dam)
	life := "+0"
	Bast.SetLifeB(&life)
	short := "Cards that channel Bast deal +1.\nPay 1 speed: flip Bast."
	Bast.SetShortB(&short)

	DBMutex.Lock()
	defer DBMutex.Unlock()
	cols = make([]string, len(props))
	for i, p := range props {
		cols[i] = strings.TrimPrefix(p, "cards.")
	}
	createDB()
	insertRecords(Ignite, Bast, Basic)
	db = nil
	t.Run("One", func(t *testing.T) {
		tests := []struct {
			name    string
			want    Card
			errWant bool
		}{
			{Ignite.Name(), Ignite, false},
			{Ignite.Name(), Ignite, false},
			{Bast.Name(), Bast, false},
			{Basic.Name(), Basic, false},

			{"Unknown_Card", nil, true},
			{"Igrath", nil, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Load(tt.name)
				if (err != nil) != tt.errWant {
					t.Error(err)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("wanted: \"%s\"\ngot: \"%s\"\n", tt.want, got)
				}
			})
		}
	})
	t.Run("Many", func(t *testing.T) {
		tests := []struct {
			name    string
			cond    string
			want    []Card
			wantErr bool
		}{
			{"Single", "name='Ignite'", []Card{Ignite}, false},
			{"Many", "name~'Ignite|Bast'", []Card{Bast, Ignite}, false},
			{"Unknown_Cards", "name~'Big Cass|La Croix|Solid Snake'", []Card{}, false},

			{"Invalid_Query", "name~", []Card{}, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := LoadMany(tt.cond)
				if (err != nil) != tt.wantErr {
					t.Errorf("LoadMany() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LoadMany() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func TestStatsString(t *testing.T) {
	tests := []struct {
		name string
		s    stats
		want string
	}{
		{"0/0/0", stats{speed: 0, damage: 0, life: 0}, ""},
		{"1/1/1", stats{speed: 1, damage: 1, life: 1}, "1/1"},
		{"1/1/3", stats{speed: 1, damage: 1, life: 3}, "1/3"},
		{"2/1/1", stats{speed: 2, damage: 1, life: 1}, "2/1/1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("stats.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
