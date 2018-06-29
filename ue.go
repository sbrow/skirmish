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
	resolve, _ := strconv.Atoi(c.resolve)
	obj.Stats = Stats{Life: c.Life(), Damage: c.Damage(), Speed: c.Speed(),
		Resolve: resolve, Short: c.Short(), Long: c.Long(), Flavor: c.Flavor()}
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
	Resolve       string
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
		Resolve:       "None",
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
	obj.Copies = d.copies
	obj.Visual = *NewVisual(d.name, d.leader, d.copies)
	//TODO(sbrow): UE COST BROKEN

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

// TODO(sbrow): Fix UECardJSON
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
	if n.ResolveB != nil {
		resolve, err := strconv.Atoi(*n.ResolveB)
		if err != nil {
			return []byte{}, err
		}
		if n.SpeedB != nil {
			obj.ActiveStats.Speed = *n.SpeedB
		}
		if n.DamageB != nil {
			obj.ActiveStats.Damage = *n.DamageB
		}
		if n.LifeB != nil {
			life, err := strconv.Atoi(*n.LifeB)
			if err != nil {
				return []byte{}, err
			}
			obj.ActiveStats.Life = life
		}
		if n.ShortB != nil {
			obj.ActiveStats.Short = *n.ShortB
		}
		if n.LongB != nil {
			obj.ActiveStats.Long = *n.LongB
		}
		if n.FlavorB != nil {
			obj.ActiveStats.Flavor = *n.FlavorB
		}
		obj.ActiveStats.Resolve = resolve
	}
	obj.Visual.BackTexture = strings.Replace(obj.Visual.BackTexture,
		"CardBack", fmt.Sprintf("01x_%s_Halo", n.name), -1)
	// mat := "MaterialInstanceConstant'/Game/Materials"
	// obj.Visual.FrontMaterial = fmt.Sprintf("%s/Card%s_Inst.Card%[2]s_Inst'", mat, "Front")
	// obj.Visual.BackMaterial = fmt.Sprintf("%s/Card%s_Inst.Card%[2]s_Inst'", mat, "Back")
	obj.DeckCards = fmt.Sprintf("DataTable'/Game/Data/%sDeck.%[1]sDeck'", n.name)
	if ident {
		return json.MarshalIndent(obj, "", "\t")
	}
	return json.Marshal(obj)
}
