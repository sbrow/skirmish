package skirmish

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	c := Load("Anger")
	fmt.Printf("%#v\n", c)
	l := Load("Bast")
	fmt.Printf("%#v\n", l)
}
