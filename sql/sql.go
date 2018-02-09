package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", "root", os.Getenv("SQLPASS"),
		"Mysql"))
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(`SELECT * FROM skirmish.variables`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var key string
		var value string
		if err := rows.Scan(&key, &value); err != nil {
			panic(err)
		}
		fmt.Println(value)
	}
}
