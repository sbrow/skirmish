package skirmish

import (
	"reflect"
	"testing"
)

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

	db = nil
	t.Run("One", func(t *testing.T) {
		tests := []struct {
			name    string
			want    Card
			errWant bool
		}{
			{Ignite.Name(), Ignite, false},

			{"Unknown_Card", nil, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := Load(tt.name)
				if (err != nil) != tt.errWant {
					t.Error(err)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("wanted: %s\ngot: %s\n", tt.want, got)
				}
			})
		}
	})
	db = nil
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.want {
				t.Errorf("NonDeckCard.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
