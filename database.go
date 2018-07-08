package skirmish

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	// PSQL Driver.
	_ "github.com/lib/pq"
)

// The database to retrieve card info from.
var db *sql.DB

// Connect connects to a postgreSQL database with the given options:
// host is the ip of the server,
// port is the server port,
// dbname is the name of the database,
// user is the username, and
// sslmode declares which ssl mode to use.
//
// See github.com/lib/pq for more information on sslmode.
func Connect(host string, port int, dbname, user, sslmode string) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, os.Getenv("PSQL_PWD"), dbname, sslmode)
	var err error
	db, err = sql.Open("postgres", connStr)
	return err
}

// Dump runs pg_dump, saving the contents of the standard database
// to a SQL file (skirmish_db.sql) in the given directory.
//
// TODO(sbrow): change Dump() to support path instead of dir
func Dump(dir string) {
	var out, errs bytes.Buffer

	cmd := exec.Command("pg_dump", "-U", "postgres", "-n", "skirmish", "-n", "public",
		"-c", "--if-exists", "--column-inserts", "-f", filepath.Join(dir, "skirmish_db.sql"))
	cmd.Stdout = &out
	cmd.Stderr = &errs
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
	if len(out.String()) > 0 {
		fmt.Println(out.String())
	}
	if len(errs.String()) > 0 {
		fmt.Println(errs.String())
	}
}

// Query returns the results of a query to the standard database.
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args...)
}

// Recover runs pg_recover on the database, loading data from the SQL file in the given directory.
//
// TODO(sbrow): change Recover() to support path instead of dir
func Recover(dir string) {
	var out, errs bytes.Buffer

	cmd := exec.Command("psql", "-U", "postgres", "-f", filepath.Join(dir, "skirmish_db.sql"))
	cmd.Stdout = &out
	cmd.Stderr = &errs
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String(), "\n", errs.String())
}
