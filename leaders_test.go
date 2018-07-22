// TODO(sbrow): Test Leaders.load [Issue](https://github.com/sbrow/skirmish/issues/37)
package skirmish

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestImageDir(t *testing.T) {
	Cfg.Load("config.yml")
	want := filepath.Join(Cfg.PS.Dir, "Images")
	if got := ImageDir; got != want {
		t.Fatalf("Wanted: \"%s\"\nGot: \"%s\"", want, got)
	}
}

func TestLeaders(t *testing.T) {
	if len(Leaders) == 0 {
		t.Error("slice 'Leaders' has not been initialized")
	}
}

func Test_Leaders_Names(t *testing.T) {
	tests := []struct {
		name string
		l    leaders
		want []string
	}{
		{"1", []leader{{Name: "Guy"}}, []string{"Guy"}},
		{"2", []leader{
			{Name: "Guy"},
			{Name: "Fawkes", Banner: []uint8{0}, Indicator: []uint8{255}},
		},
			[]string{"Guy", "Fawkes"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leaders.names() = %v, want %v", got, tt.want)
			}
		})
	}
}
