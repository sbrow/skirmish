package deck

import (
	"fmt"
	"log"
	"os"
	"testing"
)

/*func TestRarity(t *testing.T) {
	d := NewCard()
	d.Name = "Bushwack Squad"
	d.rarity = 3
	// s := strings.Split(d.Rarity(), ",")
	fmt.Println(strings.Split(d.Rarity(), Delim)[Rarities["Common"]%3])
	// if s[0] != Rarities[Common] {
	// t.Fatal("Rarities returned false.")
	// }
}
*/

func TestSetArts(t *testing.T) {
	c := NewCard()
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

func TestImage_One(t *testing.T) {
	c := NewCard()
	c.Name = "Blaze"
	c.Rarity = 3
	c.Arts = 1
	_, err := c.Image("Bast")
	if err != nil {
		t.Fatal(err)
	}
}

func TestImage_Many(t *testing.T) {
	c := NewCard()
	c.Name = "Loyal Trooper"
	c.Rarity = 3
	c.Arts = 3
	path, err := c.Image("Igrath")
	if err != nil {
		log.Println(path)
		t.Fatal(err)
	}
}

// TODO: Add fail conditions.
func TestString(t *testing.T) {
	t.Run("Action", func(t *testing.T) {
		d := NewCard()
		d.Name = "Smoke Trap"
		d.Cost = 2
		d.Type = "Action"
		fmt.Println(d)
	})

	t.Run("Follower", func(t *testing.T) {
		d := NewCard()
		d.Leader = &Card{Name: "Scinter"}
		d.Name = "Big Ninja"
		d.Rarity = Uncommon
		d.Cost = 3
		d.Type = "Follower"
		d.Damage = 3
		d.Toughness = 3
		fmt.Println(d)
	})

	t.Run("Hero", func(t *testing.T) {
		d := NewCard()
		d.Leader = &Card{Name: "Vi"}
		d.Name = "Rumour"
		d.Rarity = Uncommon
		d.Speed = 1
		d.Cost = 3
		d.Type = "Deck Hero"
		d.Damage = 2
		d.Life = 8
		fmt.Println(d)
	})

	t.Run("EventContinuous", func(t *testing.T) {
		d := NewCard()
		d.Leader = &Card{Name: "Igrath"}
		d.Name = "High Ground"
		d.Cost = 2
		d.Type = "Event- Continuous"
		fmt.Println(d)
	})
}

func TestDeck_String(t *testing.T) {
	d := New("F:\\GitLab\\dreamkeepers-psd\\card_jsons\\Bast.json")
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintln(f, d.Labels(), d.String())
}
func BenchmarkDeckString(b *testing.B) {
	d := New("F:\\GitLab\\dreamkeepers-psd\\card_jsons\\Bast.json")
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b.ResetTimer()
	b.Run("Wait", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fmt.Fprint(f, d.String())
		}
	})

	/*	b.ResetTimer()
		b.Run("MutMany", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fmt.Fprint(f, d.StringMut())
			}
		})
	*/
}
