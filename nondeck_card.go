package skirmish

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// TODO(sbrow): Make getters for NonDeckCard
type NonDeckCard struct {
	card
	faction  string
	resolveB *string
	statsB   *stats
	shortB   *string
	longB    *string
	flavorB  *string
}

func (n *NonDeckCard) DamageB() int {
	if n.statsB == nil {
		return 0
	}
	return (*n.statsB).damage
}

func (n *NonDeckCard) LifeB() string {
	if n.statsB == nil {
		return ""
	}
	life := (*n.statsB).life
	if life >= 0 {
		return fmt.Sprintf("+%d", life)
	}
	return fmt.Sprintf("-%d", life)
}

func (n *NonDeckCard) Faction() string {
	return n.faction
}

func (n *NonDeckCard) SpeedB() int {
	if n.statsB == nil {
		return 0
	}
	return (*n.statsB).speed
}

func (n *NonDeckCard) SetDamageB(d *int) {
	if n.statsB == nil {
		n.statsB = &stats{}
	}
	if d != nil {
		n.statsB.damage = *d
	}
}

func (n *NonDeckCard) SetFaction(faction *string) {
	if faction != nil {
		n.faction = *faction
	}
}

func (n *NonDeckCard) SetFlavorB(f *string) {
	n.flavorB = f
}
func (n *NonDeckCard) SetLifeB(l *string) {
	if n.statsB == nil {
		n.statsB = &stats{}
	}
	if l != nil {
		life, err := strconv.Atoi(*l)
		if err != nil {
			log.Println(err)
		} else {
			n.statsB.life = life
		}
	}
}

func (n *NonDeckCard) SetLongB(l *string) {
	n.longB = l
}

func (n *NonDeckCard) SetResolveB(r *string) {
	n.resolveB = r
}

func (n *NonDeckCard) SetShortB(s *string) {
	n.shortB = s
}

func (n *NonDeckCard) SetSpeedB(s *int) {
	if n.statsB == nil {
		n.statsB = &stats{}
	}
	if s != nil {
		n.statsB.speed = *s
	}
}

func (n *NonDeckCard) StatsB() string {
	reg := regexp.MustCompile(`(\/)([^\/-])*$`)
	return reg.ReplaceAllString(n.statsB.String(), "/+$2")
}
func (n *NonDeckCard) String() string {
	return fmt.Sprintf("%s //\n%s (Halo) %s %s %s \"%s\"", n.card.String(), n.card.Name(), *n.resolveB,
		n.card.Type(), n.StatsB(), pruneNewLines(*n.shortB))
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
