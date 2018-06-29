package skirmish

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// TODO(sbrow): Make getters/setters for NonDeckCard
type NonDeckCard struct {
	card
	faction  string
	SpeedB   *int
	ResolveB *string
	DamageB  *int
	LifeB    *string
	ShortB   *string
	LongB    *string
	FlavorB  *string
}

func (n *NonDeckCard) Faction() string {
	return n.faction
}

func (n *NonDeckCard) SetFaction(faction string) {
	n.faction = faction
}
func (n *NonDeckCard) String() string {
	return fmt.Sprintf("%s //\n%s (Halo) %s %s %d/%d/%s \"%s\"", n.card.String(), n.card.Name(), *n.ResolveB,
		n.card.Type(), *n.SpeedB, *n.DamageB, *n.LifeB, pruneNewLines(*n.ShortB))
}

func (n *NonDeckCard) Images() (paths []string, err error) {
	basePath := filepath.Join(ImageDir, "Heroes", n.Name())
	imgs := []string{basePath + ".png", basePath + " (Halo).png"}
	for i, img := range imgs {
		if _, err = os.Stat(img); os.IsNotExist(err) {
			imgs[i] = DefaultImage
		}
	}
	if imgs[1] == DefaultImage {
		return []string{imgs[0]}, nil
	}
	return imgs, nil
}

// SetCard makes c the DeckCard's base card.
func (n *NonDeckCard) SetCard(c Card) {
	n.card = c.Card()
}
func (n *NonDeckCard) Labels() []string {
	labels := append(n.card.Labels(), "Halo", "Troika", "Nightmares")
	return labels
}

func (n *NonDeckCard) CSV(lbls bool) [][]string {
	out := n.card.CSV(true)
	out[0] = n.Labels()
	l := n.Labels()[len(n.card.Labels()):]
	for _, label := range l {
		switch label {
		case "Halo":
			out[1] = append(out[1], "false")
		case "Troika":
			fallthrough
		case "Nightmares":
			out[1] = append(out[1], fmt.Sprint(n.Faction() == label))
		}
	}
	imgs, err := n.Images()
	if err != nil {
		log.Println(err)
	}
	tmp := make([][]string, len(imgs)+1, len(out[0]))
	tmp[0] = out[0]
	tmp[1] = out[1]
	i := -1
	for j, col := range out[0] {
		if col == "card_image" {
			i = j
			break
		}
	}
	out = tmp
	out[1][0] = n.name
	out[1][1] = fmt.Sprintf("%s- %s", n.Type(), n.name)
	out[1][2] = string(out[1][2][1])
	out[1][i] = imgs[0]
	// out[1][20] = "true"
	for j := 2; j <= len(imgs); j++ {
		out[j] = make([]string, len(out[j-1]))
		copy(out[j], out[j-1])
		out[j][0] = fmt.Sprintf("%s (Halo)", n.name)
		out[j][1] = fmt.Sprintf("%s- %s", n.Type(), n.name)
		out[j][i] = imgs[j-1]
		out[j][20] = "true"
	}
	if lbls {
		return out
	}
	return out[1:]
}
