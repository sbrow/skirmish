package skirmish

import "testing"

func TestImageDir(t *testing.T) {
	want := `F:\GitLab\dreamkeepers-psd\Images`
	got := ImageDir
	if got != want {
		t.Fatalf("Wanted: \"%s\"\nGot: \"%s\"", want, got)
	}
}

func TestLeaders(t *testing.T) {
	if len(Leaders) == 0 {
		t.Error("slice 'Leaders' has not been initialized")
	}
}
