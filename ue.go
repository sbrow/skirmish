package skirmish

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type CardUEJSON struct {
	Name       string
	Type       string `json:"Class"`
	Visual     Visual
	Stats      Stats
	Abilities  []string
	SystemData SystemData
}

func (c card) UEJSON(ident bool) ([]byte, error) {
	obj := CardUEJSON{}
	obj.Name = c.name
	obj.Type = "CTE_" + c.ctype
	resolve, err := strconv.Atoi(c.resolve)
	if err != nil {
		return []byte{}, err
	}
	obj.Stats = Stats{Life: c.life, Damage: c.damage, Speed: c.speed,
		Resolve: resolve, Short: c.short, Long: c.long, Flavor: c.flavor}
	obj.Abilities = make([]string, 0)
	obj.Visual = *NewVisual(c.name, "Common", 1)
	obj.SystemData = SystemData{make([]string, 0), make([]string, 0), make([]string, 0)}
	if ident {
		return json.MarshalIndent(obj, "", "\t")
	}
	return json.Marshal(obj)
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

type DeckCardUEJSON struct {
	CardUEJSON
	CardName   string
	Leader     string
	Supertypes string `json:"Subclass"`
	Copies     int    `json:"CardCountInDeck"`
}

func (d DeckCard) UEJSON(ident bool) ([]byte, error) {
	byt, err := d.card.UEJSON(ident)
	if err != nil {
		log.Panic(err)
	}
	obj := DeckCardUEJSON{}
	err = json.Unmarshal(byt, &obj.CardUEJSON)
	if err != nil {
		log.Panic(err)
	}
	obj.CardName = d.name
	obj.Supertypes = "CTE_" + strings.Join(d.stype, "_")
	obj.Name = strings.Replace(d.name, " ", "", -1)
	obj.Leader = d.leader
	obj.Copies = d.rarity
	obj.Visual = *NewVisual(d.name, d.leader, d.rarity)
	//TODO: UE COST BROKEN
	// obj.Stats.Cost = d.cost
	if ident {
		return json.MarshalIndent(obj, "", "\t")
	}
	return json.Marshal(obj)
}

type NonDeckCardUEJSON struct {
	CardUEJSON
	Faction     string
	DeckCards   string
	Power       bool `json:"bHaveActivePower"`
	ActiveStats Stats
}

func (n NonDeckCard) UEJSON(ident bool) ([]byte, error) {
	byt, err := n.card.UEJSON(ident)
	if err != nil {
		log.Panic(err)
	}
	obj := NonDeckCardUEJSON{}
	err = json.Unmarshal(byt, &obj.CardUEJSON)
	if err != nil {
		log.Panic(err)
	}
	obj.Faction = "FE_" + n.faction
	if n.resolveB != nil {
		resolve, err := strconv.Atoi(*n.resolveB)
		if err != nil {
			return []byte{}, err
		}
		if n.speedB != nil {
			obj.ActiveStats.Speed = *n.speedB
		}
		if n.damageB != nil {
			obj.ActiveStats.Damage = *n.damageB
		}
		if n.lifeB != nil {
			life, err := strconv.Atoi(*n.lifeB)
			if err != nil {
				return []byte{}, err
			}
			obj.ActiveStats.Life = life
		}
		if n.shortB != nil {
			obj.ActiveStats.Short = *n.shortB
		}
		if n.longB != nil {
			obj.ActiveStats.Long = *n.longB
		}
		if n.flavorB != nil {
			obj.ActiveStats.Flavor = *n.flavorB
		}
		obj.ActiveStats.Resolve = resolve
	}
	obj.Visual.BackTexture = strings.Replace(obj.Visual.BackTexture,
		"CardBack", fmt.Sprintf("01x_%s_Halo", n.name), -1)
	mat := "MaterialInstanceConstant'/Game/Materials"
	obj.Visual.FrontMaterial = fmt.Sprintf("%s/Card%s_Inst.Card%[2]s_Inst'",
		mat, "Front")
	obj.Visual.BackMaterial = fmt.Sprintf("%s/Card%s_Inst.Card%[2]s_Inst'",
		mat, "Back")
	obj.DeckCards = fmt.Sprintf("DataTable'/Game/Data/%sDeck.%[1]sDeck'", n.name)
	if ident {
		return json.MarshalIndent(obj, "", "\t")
	}
	return json.Marshal(obj)
}
