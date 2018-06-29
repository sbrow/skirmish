package skirmish

import (
	"fmt"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestConnect(t *testing.T) {
	type args struct {
		user    string
		dbname  string
		sslmode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"standard", args{"postgres", "postgres", "disable"}, false},
		{"ssl_Enabled", args{"postgres", "postgres", "enable"}, false},
		// TODO(sbrow): find out why connecting to database with non-existant user doesn't throw an error.
		{"wrong user", args{"", "postgres", "disable"}, false},

		{"wrong database", args{"postgres", "postgres2", "disable"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Connect(tt.args.user, tt.args.dbname, tt.args.sslmode); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	fmt.Println(DB)
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
	t.Run("One", func(t *testing.T) {
		tests := []struct {
			name    string
			want    Card
			errWant error
		}{
			{Ignite.Name(), Ignite, nil},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Load(tt.name)
				if err != tt.errWant {
					t.Error(err)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("wanted: %s\ngot: %s\n", tt.want, got)
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
