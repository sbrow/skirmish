// TODO: config file for non-programmers
package skirmish

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// const LeaderSQL = "SELECT * FROM Skirmish.NonDeckCards WHERE NOT Faction='Neutral'"
// const LeaderNameSQL = "SELECT Name FROM Skirmish.NonDeckCards WHERE NOT Faction='Neutral'"

var Template = filepath.Join(os.Getenv("SK_PS"), "Template009.1.psd")
var DataDir = filepath.Join(os.Getenv("SK_SQL"))
var ImageDir = filepath.Join(os.Getenv("SK_PS"), "Images")

var Leaders []string //Load from SQL
var Database *sql.DB

func init() {
	var err error
	// Initialize database.
	Database, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", "root",
		os.Getenv("SQLPASS"), "Mysql"))
	if err != nil {
		panic(err)
	}

	// Load data into database.
	// TODO: Only load if sha doesn't match.
	err = loadSQL("sql/db.sql") // TODO: Temp
	if err != nil {
		panic(err)
	}

	loadLeaders()
}

// loadSQL executes a given SQL file on the database.
// It is intended to update all data to its most current iteration.
func loadSQL(path string) error {
	fmt.Println(path)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		if request != "" {
			_, err := Database.Exec(request)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// loadLeaders loads the list of valid deck leaders from the database.
func loadLeaders() {
	leaders, err := Database.Query(
		"SELECT Name FROM Skirmish.NonDeckCards WHERE NOT Faction='Neutral'")
	defer leaders.Close()
	if err != nil {
		panic(err)
	}
	for leaders.Next() {
		var name string
		if err := leaders.Scan(&name); err != nil {
			panic(err)
		}
		Leaders = append(Leaders, name)
	}
}

// verifyDatabase checks to make sure all calculated fields are up to date.
func verifyDatabase() {

}
