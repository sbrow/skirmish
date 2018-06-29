package skirmish

import "testing"

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
