package sql

import (
	// "encoding/csv"
	"fmt"
	// "log"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	Init(filepath.Join(os.Getenv("SK_PS"), "Images"), os.Getenv("SK_SQL"))
}

// func TestSimple(t *testing.T) {
// 	var text *string
// 	err := Database.QueryRow("SELECT type FROM public.all_cards WHERE name='Anger'").Scan(&text)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(*text)

// }

func TestLoad(t *testing.T) {
	c := Load("Anger")
	// l := Load("Bast")
	fmt.Printf("%#v\n", c)
	// fmt.Printf("%#v\n", l)
	fmt.Println(c.CSV())
	// fmt.Println(l)

}

/*
func TestCSV(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping CSV test")
	}
	c := Load("Bushwack Squad")
	file, err := os.Create("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	w.WriteAll(c.CSV())
}

func TestQuery(t *testing.T) {
	var name, typ *string
	query := `SELECT name, type FROM skirmish.deckcards WHERE "name"=$1`
	Database.QueryRow(query, "Anger").Scan(&name, &typ)
	if name == nil || typ == nil {
		log.Fatal("Noop!")
	}
}

func TestSQL(t *testing.T) {
	Recover(DataDir)
	Dump(DataDir)
}
*/
