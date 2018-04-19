// Command skir is the primary skirmish command.
//
// Build
//
// Generates photoshop dataset (csv) files with data pulled from the PostgresSQL
// database.
//
//		skir build
//		skir build regex / skirmish build -r
//		skir build bold  / skirmish build -b
//		skir build data  / skirmish build -d
//
// PS
//
// Fills out a document (psd) template with data from a dataset (csv) file and
//
// 		skir ps
//		skir ps -card=$CARDNAME // Builds the psd for the given card
//		skir ps -deck=$DECKNAME
//		skir ps action $ACTIONNAME
//
// Card
//
// loads card information from the database and outputs it to STDOUT.
//
//		skir card $CARDNAME
//
// TODO: config file for non-programmers
// TODO: Separate out all subprocesses.
package main

import (
	"flag"
	"fmt"
	// app "github.com/sbrow/ps"
	"github.com/sbrow/skirmish"
	// "github.com/sbrow/skirmish/build"
	// "bytes"
	// "github.com/sbrow/skirmish/ps"
	"github.com/sbrow/skirmish/sql"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func init() {
	cmd := exec.Command("go", "install",
		"github.com/sbrow/skirmish/ps/cmd")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if len(output) > 0 {
		fmt.Println(string(output))
	}
	if err != nil {
		log.Panic(err)
	}

}
func main() {
	log.SetPrefix("[main] ")
	card := flag.String("card", "", "card get info on a card.")
	flag.Parse()
	var args []string
	var cmd string
	switch {
	case len(os.Args) > 2:
		args = os.Args[2:]
		fallthrough
	case len(os.Args) > 1:
		cmd = os.Args[1]
	}
	switch {
	case cmd == "ps":
		comm := exec.Command(filepath.Join(os.Getenv("GOBIN"), "cmd.exe"), args...)
		comm.Stdout = os.Stdout
		comm.Stdin = os.Stdin
		comm.Stderr = os.Stderr
		comm.Run()
		return
	case cmd == "card" || *card != "":
		name := *card
		if name == "" {
			name = strings.Join(args, " ")
		}
		card, err := sql.Load(name)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(card)
	case cmd == "build":
		log.SetPrefix("[main] ")
		log.Println("Generating cards")
		sql.GenData()
		log.Println("Cards successfully generated!")
	case cmd == "db":
		opt := args[0]
		if opt == "save" {
			sql.Dump(skirmish.DataDir)
		} else if opt == "load" {
			sql.Recover(skirmish.DataDir)
		}
	case cmd == "select":
		query := fmt.Sprintf("SELECT %s %s", strings.Join(args[:len(args)-2], ", "),
			strings.Join(args[len(args)-2:], " "))
		rows, err := skirmish.DB.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		columns, _ := rows.Columns()
		fmt.Println(strings.Join(columns, "\t\t"))
		count := len(columns)
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)

		for rows.Next() {

			for i, _ := range columns {
				valuePtrs[i] = &values[i]
			}

			rows.Scan(valuePtrs...)

			for i := range columns {

				var v interface{}

				val := values[i]

				b, ok := val.([]byte)

				if ok {
					v = string(b)
				} else {
					v = val
				}

				fmt.Printf("%v\t\t", v)
			}
			fmt.Println()
		}
	}
}
