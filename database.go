package skirmish

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	// PSQL Driver.
	_ "github.com/lib/pq"
)

// The database to retrieve card info from.
var db *sql.DB

// Connect connects to a postgreSQL database with the given options:
// host is the ip of the server,
// port is the server port,
// DBName is the name of the database,
// user is the username, and
// sslmode declares which ssl mode to use.
//
// See github.com/lib/pq for more information on sslmode.
func Connect(host string, port int, DBName, user, sslmode string) error {
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s sslmode=%s",
		host, port, DBName, user, sslmode)
	pwd, ok := os.LookupEnv("PSQL_PWD")
	if ok && !(user == "guest") {
		connStr += fmt.Sprintf(" password=%s", pwd)
	}
	var err error

	db, err = sql.Open("postgres", connStr)
	return err
}

// Dump runs pg_dump, saving the contents of the standard database
// to a SQL file (skirmish_db.sql) in the given directory.
func Dump(path string) {
	args := []string{
		"-h", LocalDB.DB.Host,
		"-p", fmt.Sprint(LocalDB.DB.Port),
		"-U", LocalDB.DB.User,
		"-d", LocalDB.DB.Name,
		"-n", "public",
		"--if-exists",
		"-c",
		"--column-inserts",
		"-f", path,
	}
	cmd := exec.Command("pg_dump", args...)
	var out, errs bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errs
	log.Println(args)
	if err := cmd.Run(); err != nil {
		log.Println(err)
		if len(out.Bytes()) > 0 {
			log.Println(out.String())
		}
		if len(errs.Bytes()) > 0 {
			log.Println(errs.String())
		}
	}
}

// Exec executes a query on the standard database without returning any rows.
// The args are for any placeholder parameters in the query.
func Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

// Query returns the results of a query to the standard database.
// The args are for any placeholder parameters in the query.
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args...)
}

// Recover runs pg_recover on the database, loading data from the SQL file in the given directory.
func Recover(path string) (sql.Result, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return Exec(string(data))
}
