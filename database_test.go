package skirmish

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"

	_ "github.com/lib/pq"
)

var DBMutex sync.Mutex

const PSQLVersion = "psql (PostgreSQL) 10.2\r\n"

func TestConnect(t *testing.T) {
	type args struct {
		Host string
		Port int
		User string
		Name string
		SSL  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"standard", args{"localhost", 5432, "postgres", "skirmish", "disable"}, false},
		{"ssl_Enabled", args{"localhost", 5432, "postgres", "skirmish", "require"}, false},

		// TODO(sbrow): find out why connecting to database with non-existent user/database doesn't throw an error. [Issue](https://github.com/sbrow/skirmish/issues/35)
		{"wrong_user", args{"localhost", 5432, "butts", "skirmish", "disable"}, false},
		{"wrong_database", args{"localhost", 5432, "postgres", "skirmish23", "disable"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DBMutex.Lock()
			defer DBMutex.Unlock()
			err := Connect(tt.args.Host, tt.args.Port, tt.args.Name, tt.args.User, tt.args.SSL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRecover(t *testing.T) {
	tests := []struct {
		name    string
		dir     string
		want    []string
		wantErr bool
	}{
		{"", filepath.Join(".", "_test", "test.sql"), []string{"ex_1", "ex_2"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DBMutex.Lock()
			defer DBMutex.Unlock()
			err := Connect(LocalDB.DBArgs())
			if err != nil {
				t.Fatal(err)
			}
			_, err = Recover(tt.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("Recover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			query := "SELECT text_field FROM example"
			rows, err := Query(query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query(\"%s\") error = %v, wantErr %v", query, err, tt.wantErr)
				return
			}

			got := make([]string, 0)
			for rows.Next() {
				var result *string
				rows.Scan(&result)
				if result != nil {
					got = append(got, *result)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wanted: %v\ngot: %v", tt.want, got)
			}
		})
	}
}

// TODO(sbrow): TestDump accurate testing. [Issue](https://github.com/sbrow/skirmish/issues/36)
func TestDump(t *testing.T) {
	var out bytes.Buffer
	cmd := exec.Command("psql", "-V")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		t.Error(err)
	}
	if out.String() != PSQLVersion {
		t.Skipf("Wrong version of PSQL installed- want \"%s\", got \"%s\"",
			PSQLVersion, out.String())
	}
	type args struct {
		in  string
		out string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{in: filepath.Join("_test", "test.sql"),
			out: filepath.Join("_test", "test_dump.sql")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DBMutex.Lock()
			defer DBMutex.Unlock()
			if err := Connect(LocalDB.DBArgs()); err != nil {
				t.Fatal(err)
			}
			if _, err := Recover(tt.args.in); err != nil {
				t.Fatal(err)
			}
			path, _ := filepath.Abs(tt.args.out)
			Dump(path)
			g, err := os.Open(path)
			if err != nil {
				t.Fatal(err)
			}
			defer g.Close()
			base := filepath.Base(path)
			w, err := os.Open(strings.Replace(path, base, "."+base, 1))
			log.Println(w.Name())
			if err != nil {
				t.Fatal(err)
			}
			defer w.Close()
			/*
				got, err := ioutil.ReadAll(g)
				if err != nil {
					t.Fatal(err)
				}
				want, err := ioutil.ReadAll(w)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(string(got), string(want)) {
					t.Fatalf("wanted: %s\ngot: %s", string(want), string(got))
				}
			*/
			os.Remove(path)
		})
	}
}
