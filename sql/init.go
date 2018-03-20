package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var ImageDir string
var DataDir string //= filepath.Join(os.Getenv("SK_SQL"))
var Delim string

var Database *sql.DB
var Leaders []string

func Init(imageDir, dataDir string) {
	Delim = ","
	ImageDir = imageDir
	DataDir = dataDir

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
		`SELECT "name" FROM public.leaders`)
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
