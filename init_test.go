package skirmish

import (
	"fmt"
	"testing"
)

func TestVars(t *testing.T) {
	if Template != `F:\GitLab\dreamkeepers-psd\Template009.1.psd` {
		t.Fatal("Template is not correct.\n", Template)
	}
	if ImageDir != `F:\Gitlab\dreamkeepers-psd\Images` {
		t.Fatal("ImageDir is not correct.")
	}
	fmt.Println(Leaders)
}
