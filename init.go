package skirmish

import (
	"database/sql"
	_ "github.com/lib/pq"
	psql "github.com/sbrow/skirmish/sql"
	"os"
	"path/filepath"
)

var Template = filepath.Join(os.Getenv("SK_PS"), "Template009.1.psd")
var ImageDir = filepath.Join(os.Getenv("SK_PS"), "Images")
var DataDir = filepath.Join(os.Getenv("SK_SQL"))

// Loaded from SQL

var Database *sql.DB
var Leaders []string

func init() {
	// Connect to the database
	psql.Init(ImageDir, DataDir)
}
