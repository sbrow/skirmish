package skirmish

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"path/filepath"
)

// TODO: config file for non-programmers
var Template = filepath.Join(os.Getenv("SK_PS"), "Template009.1.psd")
var ImageDir = filepath.Join(os.Getenv("SK_PS"), "Images")

// Loaded from SQL
var Database *sql.DB
var Leaders []string

// Assigned
var Delim = ","

func init() {
	// Connect to the database
	var err error
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		"postgres", "toor", "postgres", "disable")
	Database, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Load the leaders
	query, err := Database.Query(
		`SELECT "name" FROM skirmish.nondeckcards WHERE NOT faction='Neutral'`)
	defer query.Close()
	if err != nil {
		panic(err)
	}
	for query.Next() {
		var name string
		if err := query.Scan(&name); err != nil {
			panic(err)
		}
		Leaders = append(Leaders, name)
	}
}
