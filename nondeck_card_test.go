package skirmish

import (
	"reflect"
	"testing"
)

func TestNonDeckCard_CSV(t *testing.T) {
	card := &NonDeckCard{
		card: card{
			name:     "Bast",
			cardType: "Hero",
			// superTypes: nil,
			resolve: "+2",
			stats: stats{
				speed:  1,
				damage: 2,
				life:   14,
			},
			short: "Short range.\nPay 1 speed: flip Bast.",
			long:  "Characters with Short range can’t flank or reinforce\nSpeed can be spent to attack, intercept, or redeploy.",
			// regexp: "",
		},
		faction: "Troika",
		statsB: &stats{
			speed:  1,
			damage: 1,
			life:   0,
		},
	}
	r := "+2"
	card.resolveB = &r
	sh := "Cards that channel Bast deal +1."
	card.shortB = &sh

	tests := []struct {
		name   string
		n      *NonDeckCard
		labels bool
		want   [][]string
	}{
		{"Bast", card, true, [][]string{
			card.Labels(),
			{"Bast", "Hero- Bast", "2", "1", "2", "14", card.short, card.long, `F:\GitLab\dreamkeepers-psd\Images\Heroes\Bast.png`,
				"false", "false", "false", "false", "false", "false", "false", "false", "false", "false", "false", "true", "false"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.CSV(tt.labels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NonDeckCard.CSV() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
