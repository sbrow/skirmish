package sql

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	_, err := Query(`select * from "PUBLIC"."Bast";`)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFunc(t *testing.T) {
	out, err := Query(`select * from "PUBLIC"."Bast";select * from "PUBLIC"."Igrath";`)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(out))
}
