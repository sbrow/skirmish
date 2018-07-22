package ps

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/sbrow/ps"
	"github.com/sbrow/skirmish"
)

func Test_template_SetLeader(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Photoshop is likely not installed")
	}
	type args struct {
		template string
		wantStr  string
		leader   string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Interrogation Chamber", args{
			template: CardTemplate,
			wantStr:  "SetLeader.psd",
			leader:   "Tinsel"}},
	}
	for _, tt := range tests {
		mode := ps.Safe
		if testing.Short() {
			mode = ps.Normal
		}
		t.Run(tt.name, func(t *testing.T) {
			// got := new(mode, tt.args.template)
			// got.SetLeader(tt.args.leader)
			want := new(mode, LoadTest(tt.args.wantStr))
			var leader *skirmish.Leader
			for _, ldr := range skirmish.Leaders {
				leaderInd := want.DeckInd.MustExist(ldr.Name).(*ps.ArtLayer)
				if ldr.Name == tt.args.leader {
					leader = &ldr
				}
				if leaderInd.Visible() != (ldr.Name == tt.args.leader) {
					t.Fatalf("leader \"%s\".Visible() == %t, want %t", ldr.Name,
						leaderInd.Visible(), ldr.Name == tt.args.leader)
				}
			}
			if leader == nil {
				t.Fatalf("Leader %s was not found", tt.args.leader)
			}
			defer want.Doc.Dump()
			got := new(mode, tt.args.template)
			got.SetLeader(tt.args.leader)
			defer got.Doc.Dump()

			if !reflect.DeepEqual(got.DeckInd, want.DeckInd) {
				t.Errorf("wanted:\n%+v\ngot:\n%+v", want.DeckInd, got.DeckInd)
			}
		})
	}
}

func LoadTest(filename string) string {
	path := filepath.Join(skirmish.Cfg.PS.Dir, "_test", filename)
	return path
}

/*
type ArtLayer struct {
    name    string
    bounds    [2][2]int
    parent    Group
    visible    bool
    current    bool
    Color
    *Stroke
    *TextItem
}

type LayerSet struct {
    name        string
    bounds        [2][2]int
    parent        Group
    current        bool
    visible        bool
    artLayers    []*ArtLayer
    layerSets    []*LayerSet
}
*/
