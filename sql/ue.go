package sql

import (
	"fmt"
)

type CardUEJSON struct {
	Name       string
	Type       string `json:"Class"`
	Visual     Visual
	Stats      Stats
	Abilities  []string
	SystemData SystemData
}

type Visual struct {
	FrontTexture  string
	BackTexture   string
	FrameTexture  string
	DeckIcon      string
	CardTypeIcon  string
	LevelIcon     string
	StatsIcon     string
	FrontMaterial string
	BackMaterial  string
}

func NewVisual(name, leader string, copies int) *Visual {
	img := fmt.Sprintf("0%dx_%s", copies, name)
	return &Visual{
		FrontTexture:  fmt.Sprintf("Texture2D'/Game/Textures/Card_Decks/%[1]s/%[2]s.%[2]s'", leader, img),
		BackTexture:   "Texture2D'/Game/Textures/Card_Decks/Common/CardBack.CardBack'",
		FrameTexture:  "None",
		DeckIcon:      "None",
		CardTypeIcon:  "None",
		LevelIcon:     "None",
		StatsIcon:     "None",
		FrontMaterial: "None",
		BackMaterial:  "None",
	}
}

type Stats struct {
	Life        int `json:"Health_Toughness"`
	Resolve     int `json:"ResolveGain"`
	Cost        int `json:"ResolveCost"`
	CurrResolve int `json:"Resolve"`
	Damage      int `json:"Strength"`
	Speed       int
	Short       string `json:"Action"`
	Long        string `json:"Description"`
	Flavor      string `json:"Quote"`
}

type SystemData struct {
	PlayConditions        []string
	InteractionConditions []string
	TurnConditions        []string
}

// func (c CardUEJSON) MarshalJSON() ([]byte, error) {
// 	obj := map[string]interface{}{
// 		"Name":   c.Name,
// 		"Class":  "CTE_" + c.Type,
// 		"Visual": NewVisual(),
// 		"Stats":  c.Stats,
// 	}
// 	return json.Marshal(obj)
// }

type DeckCardUEJSON struct {
	CardUEJSON
	CardName   string
	Leader     string
	Supertypes string `json:"Subclass"`
	Copies     int    `json:"CardCountInDeck"`
}

type NonDeckCardUEJSON struct {
	CardUEJSON
	Faction     string
	DeckCards   string
	Power       bool `json:"bHaveActivePower"`
	ActiveStats Stats
}
