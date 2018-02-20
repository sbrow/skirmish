package skirmish

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	c := Load("Anger")
	fmt.Printf("%#v\n", c)
	fmt.Println(c.CSV())
	l := Load("Bast")
	fmt.Printf("%#v\n", l)
}

func TestCSV(t *testing.T) {
	c := Load("Bushwack Squad")
	// dat := []byte(c.CSV())
	file, err := os.Create("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	// defer w.Flush()
	w.WriteAll(c.CSV())

}
