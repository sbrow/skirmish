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
package skirmish

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"path/filepath"
)

var Template = filepath.Join(os.Getenv("SK_PS"), "Template009.1.psd")
var ImageDir = filepath.Join(os.Getenv("SK_PS"), "Images")
var DefaultImage = filepath.Join(ImageDir, "ImageNotFound.png")
var DataDir = filepath.Join(os.Getenv("SK_SQL"))

var Delim string

var DB *sql.DB
var Leaders []string

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
	query, err := DB.Query(
		`SELECT "name" FROM leaders`)
	defer query.Close()
	if err != nil {
		log.Fatal(err)
	}
	for query.Next() {
		var name string
		if err := query.Scan(&name); err != nil {
			log.Panic(err)
		}
		Leaders = append(Leaders, name)
	}
}
