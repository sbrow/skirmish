package deck

import (
	"fmt"
	"github.com/sbrow/skirmish"
	"os"
	"path/filepath"
	"testing"
)

func TestSetArts(t *testing.T) {
	c := NewDeckCard()
	c.Rarity = 3
	if c.Arts != 1 {
		t.Fail()
	}
	c.setArts(-1)
	if c.Arts != 1 {
		t.Fail()
	}
	c.setArts(4)
	if c.Arts != 1 {
		t.Fail()
	}
	c.setArts(2)
	if c.Arts != 2 {
		t.Fail()
	}
}

func TestRarity(t *testing.T) {
	c := NewDeckCard()
	rarity, err := rarity(4)
	if err == nil {
		t.Fatal("error \"%s\" was not thrown", RarityError)
	}
	c.Rarity = rarity
}

func TestImage_One(t *testing.T) {
	c := NewDeckCard()
	c.Name = "Blaze"
	c.dir = "Bast"
	c.Rarity = 3
	c.Arts = 1
	_, err := c.Image(1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestImage_Many(t *testing.T) {
	c := NewDeckCard()
	c.Name = "Loyal Trooper"
	c.dir = "Igrath"
	c.Rarity = 3
	c.Arts = 3
	_, err := c.Image(1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCard_String(t *testing.T) {
	c := NewCard()
	c.Arts = 2
	_ = c.String()
}

func TestCard_JSON(t *testing.T) {
	c := NewCard()
	json, err := c.JSON()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(json))
}

// TODO: Add fail conditions.
func TestString(t *testing.T) {
	t.Run("Action", func(t *testing.T) {
		d := NewDeckCard()
		d.Name = "Smoke Trap"
		d.Cost = "2"
		d.Type = "Action"
		_ = fmt.Sprint(d)
	})

	t.Run("Follower", func(t *testing.T) {
		d := NewDeckCard()
		d.Name = "Big Ninja"
		d.Rarity = Uncommon
		d.Cost = "3"
		d.Type = "Follower"
		d.Damage = 3
		d.Toughness = 3
		_ = fmt.Sprint(d)
	})

	t.Run("Hero", func(t *testing.T) {
		d := NewDeckCard()
		d.Name = "Rumour"
		d.Rarity = Uncommon
		d.Speed = 1
		d.Cost = "3"
		d.Type = "Deck Hero"
		d.Damage = 2
		d.Life = 8
		_ = fmt.Sprint(d)
	})

	t.Run("EventContinuous", func(t *testing.T) {
		d := NewDeckCard()
		d.Name = "High Ground"
		d.Cost = "2"
		d.Type = "Event- Continuous"
		fmt.Println(d)
	})
}

func TestDeck_String(t *testing.T) {
	name := "Bast"
	// leader := &NonDeckCard{Card: Card{Name: "Bast"}}
	d, err := New(filepath.Join(skirmish.ImageDir, name+".json")) //, leader)
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Create("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintln(f, d.Labels(), d.String())
}

func BenchmarkDeckString(b *testing.B) {
	// leader := &NonDeckCard{Card: Card{Name: "Bast"}}
	d, err := New("F:\\GitLab\\dreamkeepers-psd\\card_jsons\\Bast.json") //, leader)
	if err != nil {
		b.Fatal(err)
	}
	f, err := os.Create("test.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	b.ResetTimer()
	b.Run("Wait", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fmt.Fprint(f, d.String())
		}
	})
}

func BenchmarkCardString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := NewCard()
		_ = c.String()
	}
}
