// Package skirmish contains code for production of the
// Dreamkeepers: Skirmish battle card game.
//
// More specifically, it provides an interface between the SQL database
// that contains card data, Photoshop, and the user (via CLI).
//
// Photoshop
//
// This package selects cards from SQL and creates .csv files to be read into
// Photoshop as datasets.
//
// TODO: Cameo card flavor text.
package skirmish

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	// PSQL implementation.
	_ "github.com/lib/pq"
)

var Template = filepath.Join(os.Getenv("SK_PS"), "Template009.1.psd")
var HeroTemplate = filepath.Join(os.Getenv("SK_PS"), "Template009.1h.psd")
var ImageDir = filepath.Join(os.Getenv("SK_PS"), "Images")
var DefaultImage = filepath.Join(ImageDir, "ImageNotFound.png")
var DataDir = filepath.Join(os.Getenv("SK_SQL"))

var Delim string

var DB *sql.DB

var Leaders leaders
var Tolerances map[string]int

type Leader struct {
	Name      string
	Banner    []uint8
	Indicator []uint8
}
type leaders []Leader

func (l *leaders) Names() []string {
	s := make([]string, len(*l))
	for i, ldr := range *l {
		s[i] = ldr.Name
	}
	return s
}

func init() {
	Delim = ","
	// Connect to the Database
	var err error
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s",
		"postgres", "postgres", "disable")
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	// Load the Leaders
	rows, err := DB.Query(
		`SELECT "name", banner, indicator FROM leaders`)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		var banner []uint8
		var indicator []uint8
		if err := rows.Scan(&name, &banner, &indicator); err != nil {
			log.Panic(err)
		}
		Leaders = append(Leaders, Leader{name, banner, indicator})
	}

	Tolerances = make(map[string]int)
	rows, err = DB.Query("SELECT name, px FROM tolerances;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var px int
		if err := rows.Scan(&name, &px); err != nil {
			log.Fatal(err)
		}
		Tolerances[name] = px
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
