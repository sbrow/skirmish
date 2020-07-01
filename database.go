package skirmish

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"

	// PSQL Driver.
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// The database to retrieve card info from.
var db *sql.DB

func init() {
	viper.BindEnv("DATABASE_URL")
}

// Connect connects to a postgreSQL database with the given options:
// 		- host is the ip of the server.
// 		- port is the server port.
// 		- dbname is the name of the database.
// 		- user is the username.
// 		- sslmode declares which ssl mode to use.
//
// See package github.com/lib/pq for more information on sslmode.
func Connect(host string, port int, dbname, user, pass, sslmode string) error {
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbname, user, pass, sslmode)

	var err error
	db, err = sql.Open("postgres", connStr)
	return err
}

func AutoConnect() error {
	dbUrl := viper.Get("DATABASE_URL")
	urlString, ok := dbUrl.(string)
	if !ok {
		return errors.New("DATABASE_URL is not a string")
	}
	u, err := url.Parse(urlString)
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(u.Port())
	if err != nil {
		return err
	}
	path := strings.TrimLeft(u.Path, "/")
	password, _ := u.User.Password()
	return Connect(u.Hostname(), port, path, u.User.Username(), password, "require")
}

// Dump runs pg_dump on the connected database, saving the contents
// to a SQL formatted file at the given path.
func Dump(path string) {
	args := []string{
		"-h", Cfg.DB.Host,
		"-p", fmt.Sprint(LocalDB.DB.Port),
		"-U", Cfg.DB.User,
		"-d", Cfg.DB.Name,
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

// Query returns the results of a query to the standard database
// with the expectation that the query will return one result.
// Errors are deferred until Row's Scan method is called.
// The args are for any placeholder parameters in the query.
func QueryRow(query string, args ...interface{}) *sql.Row {
	return db.QueryRow(query, args...)
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
