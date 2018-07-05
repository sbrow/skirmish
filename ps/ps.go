// go:generate sh -c "godoc2md -template ../.doc.template github.com/sbrow/skirmish/ps > README.md"

package ps

import (
	"fmt"
	"log"
	"runtime"

	"github.com/sbrow/skirmish"
)

// TODO(sbrow): Fix err.time
type err struct {
	err error
	// time string
	file string
	line int
	ok   bool
}

func (e *err) String() string {
	if e.ok {
		return fmt.Sprintf(`%s error at %s:%d %s` /*e.time*/, "", e.file, e.line, e.err)
	}
	return fmt.Sprintf(`%s error at {corrupted_data} %s` /*e.time*/, "", e.err)
}

// Error adds an error to the list of runtime errors that have occurred so far.
func Error(e error) {
	err := err{}
	// err.time = fmt.Sprint(time.Now().Format("yyyy/MM/dd hh:mm:ss"))
	_, file, line, ok := runtime.Caller(1)
	err.err = e
	err.ok = ok
	if ok {
		err.file = file
		err.line = line
	}
	log.Println(err)
	Errors = append(Errors, err)
}

// Errors holds runtime errors that occur.
var Errors []err

var Tolerances map[string]int

// TODO(sbrow): Cover GetTolerances
func GetTolerances() {
	Tolerances = make(map[string]int)
	rows, err := skirmish.DB.Query("SELECT name, px FROM tolerances;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var px int
		if err := rows.Scan(&name, &px); err != nil {
			log.Fatal(err)
		}
		Tolerances[name] = px
	}
}
func init() {
	Errors = []err{}
	// GetTolerances()
}
