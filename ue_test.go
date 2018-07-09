package skirmish

import (
	"reflect"
	"testing"
)

var blankCardUE = `{
	"Name": "",
	"Class": "CTE_",
	"FrontTexture": "Texture2D'/Game/Textures/Card_Decks/Common/01x_.01x_'",
	"BackTexture": "Texture2D'/Game/Textures/Card_Decks/Common/CardBack.CardBack'",
	"FrameTexture": "None",
	"DeckIcon": "None",
	"CardTypeIcon": "None",
	"LevelIcon": "None",
	"StatsIcon": "None",
	"FrontMaterial": "None",
	"BackMaterial": "None",
	"Resolve": "None",
	"Stats": {
		"Health_Toughness": 0,
		"ResolveGain": 0,
		"ResolveCost": 0,
		"Resolve": 0,
		"Strength": 0,
		"Speed": 0,
		"Action": "",
		"Description": "",
		"Quote": ""
	},
	"Abilities": [],
	"SystemData": {
		"PlayConditions": [],
		"InteractionConditions": [],
		"TurnConditions": []
	}
}`

func Test_card_UEJSON(t *testing.T) {
	tests := []struct {
		name    string
		c       card
		ident   bool
		want    []byte
		wantErr bool
	}{
		// {"Empty", *(NewCard().(*card)), true, []byte(blankCardUE), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.UEJSON(tt.ident)
			if (err != nil) != tt.wantErr {
				t.Errorf("card.UEJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("card.UEJSON() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

/*
func TestNonDeckCard_UEJSON(t *testing.T) {
	type args struct {
		ident bool
	}
	tests := []struct {
		name    string
		n       NonDeckCard
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO(sbrow): Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.UEJSON(tt.args.ident)
			if (err != nil) != tt.wantErr {
				t.Errorf("NonDeckCard.UEJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NonDeckCard.UEJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
