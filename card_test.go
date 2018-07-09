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

func execDB(query string, args ...interface{}) {
	_, err := db.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
}

func createDB() {
	err := Connect("localhost", 5432, "postgres", "postgres", "disable")
	if err != nil {
		log.Fatal(err)
	}
	execDB("DROP TABLE IF EXISTS cards")
	cards := `CREATE TABLE cards (`
	for i, c := range cols {
		cards += fmt.Sprintf("\n\t%s TEXT", c)
		if i+1 != len(cols) {
			cards += ","
		}
	}
	cards += "\n)"
	execDB(cards)
}
func insertRecords(records ...Card) {
	query := `INSERT INTO cards VALUES`
	for i, r := range records {
		superTypes := "NULL"
		if len(r.STypes()) > 0 {
			superTypes = fmt.Sprintf("'%s'", strings.Join(r.STypes(), Delim))
		}
		sql := fmt.Sprintf(`('%s', '%s', %s, '%s', '%s', '%s', '%s', '%d', '%d', '%d', '%s'`,
			r.Name(), r.Type(), superTypes, r.Short(), strings.Replace(r.Long(), "'", "''", -1),
			r.Flavor(), r.Resolve(), r.Speed(), r.Damage(), r.Life(), r.Faction())
		switch r.(type) {
		case *DeckCard:
			val := r.(*DeckCard)
			cost, _ := val.Cost()
			sql += fmt.Sprintf(", '%s', '%d', '%s'", cost, val.Copies(), val.Leader())
			sql += fmt.Sprintf(",NULL, NULL, NULL, NULL, NULL, NULL, NULL, '%s'", val.Regexp())
		}
		query += "\n" + sql + ")"
		if i+1 != len(records) {
			query += ","
		}
	}
	log.Println(query)
	execDB(query + ";")
}

func init() {
	cols = make([]string, len(props))
	for i, p := range props {
		cols[i] = strings.TrimPrefix(p, "cards.")
	}
	createDB()
}

func TestLoad(t *testing.T) {
	Ignite := &DeckCard{
		card: card{
			name:    "Ignite",
			leader:  "Bast",
			ctype:   "Action",
			stype:   []string{"Channeled"},
			resolve: "",
			stats:   stats{},
			short:   "Deal 3 to a hero.",
			long: "Ignite can be played on leaders, partners, or deck heroes." +
				"\rChanneled cards must be played with their leader's resolve.",
			flavor: "There are two kinds of things in this world-" +
				" things that can catch fire, and things that are on fire.",
			regexp: `(3)|(Deal)|(hero.)`,
		},
		cost:   "1",
		copies: 3,
	}
	LoyalTrooper := &DeckCard{
		card: card{
			name:    "Loyal Trooper",
			leader:  "Igrath",
			ctype:   "Follower",
			resolve: "",
			stats:   stats{speed: 1, damage: 1, life: 3},
			short:   "Uncontested- +2/+0.",
			long:    "A Lane is uncontested if it is not contested.",
			flavor:  "Sometimes loyalty means not asking questions.",
			regexp:  `(\+2/\+0.)|(Uncontested-)`,
		},
		cost:   "2",
		copies: 3,
	}
	insertRecords(Ignite, LoyalTrooper)
	t.Run("One", func(t *testing.T) {
		tests := []struct {
			name    string
			want    Card
			errWant bool
		}{
			{Ignite.Name(), Ignite, false},
			{LoyalTrooper.Name(), LoyalTrooper, false},

			{"Unknown_Card", nil, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Load(tt.name)
				if (err != nil) != tt.errWant {
					t.Error(err)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("wanted: \"%+v\"\ngot: \"%+v\"\n", tt.want, got)
					t.Error(tt.want.Long() == got.Long())
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
			{"Many", "name~'Ignite|Loyal Trooper'", []Card{Ignite, LoyalTrooper}, false},
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

func Test_card_String(t *testing.T) {
	tests := []struct {
		name string
		c    *card
		want string
	}{
		// TODO(sbrow): Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("card.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeckCard_String(t *testing.T) {
	tests := []struct {
		name string
		d    *DeckCard
		want string
	}{
		// TODO(sbrow): Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("DeckCard.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNonDeckCard_String(t *testing.T) {
	tests := []struct {
		name string
		n    *NonDeckCard
		want string
	}{
		// TODO(sbrow): Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.want {
				t.Errorf("NonDeckCard.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
