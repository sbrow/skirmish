// go:generate sh -c "godoc2md -template ../.doc.template github.com/sbrow/skirmish/ps > README.md"

package ps

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sbrow/skirmish"
)

// TODO(sbrow): Fix psError.time [Issue](https://github.com/sbrow/skirmish/issues/45)
type psError struct {
	err error
	// time string
	file string
	line int
	ok   bool
}

type psErrors []psError

func (e psErrors) Report() error {
	log.SetOutput(os.Stdout)
	log.SetPrefix("")
	log.Printf("Process completed with %d error(s)\n", len(e))
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("runtime.Caller(0) returned !ok")
	}
	dir := filepath.Dir(filepath.Dir(file))
	f, err := os.Create(filepath.Join(dir, "errors.log"))
	if err != nil {
		return err
	}
	for _, err := range e {
		fmt.Fprintf(f, "%s\n", err.String())
	}
	return nil
}

func (e *psError) String() string {
	return fmt.Sprintf(`%s error at %s:%d %s` /*e.time*/, "", e.file, e.line, e.err)
}

// Error adds an error to the list of runtime errors that have occurred so far.
func Error(e error) {
	if e == nil {
		return
	}
	err := psError{}
	// err.time = fmt.Sprint(time.Now().Format("yyyy/MM/dd hh:mm:ss"))
	_, file, line, ok := runtime.Caller(1)
	err.err = e
	err.ok = ok
	if ok {
		err.file = filepath.Clean(file)
		err.line = line
	}
	log.Println(err)
	Errors = append(Errors, err)
}

// Errors holds runtime errors that occur.
var Errors psErrors

// Tolerances holds values for offset of template objects.
var Tolerances map[string]int

// GetTolerances selects the tolerance values from the database into Tolerances.
//
// See 'Tolerances'.
func GetTolerances() error {
	Tolerances = make(map[string]int)
	rows, err := skirmish.Query("SELECT name, px FROM tolerances;")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var px int
		rows.Scan(&name, &px)
		Tolerances[name] = px
	}
	return nil
}

func init() {
	Errors = []psError{}
	if err := GetTolerances(); err != nil {
		Error(err)
	}
}
