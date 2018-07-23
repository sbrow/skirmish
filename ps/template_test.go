package ps

import (
	"log"
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
		name    string
		args    args
		wantErr bool
	}{
		{"Interrogation Chamber", args{
			template: CardTemplate,
			wantStr:  "SetLeader.psd",
			leader:   "Tinsel"},
			false,
		},
		{"Interrogation Chamber", args{
			template: CardTemplate,
			wantStr:  "SetLeader.psd",
			leader:   "Flobble"},
			false,
		},
	}
	for _, tt := range tests {
		mode := ps.Safe
		if testing.Short() {
			mode = ps.Normal
		}
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				Errors.Report()
			}()
			want := new(mode, LoadTest(tt.args.wantStr))
			var leader skirmish.Leader
			for _, ldr := range skirmish.Leaders {
				leaderInd := want.DeckInd.MustExist(ldr.Name).(*ps.ArtLayer)
				if ldr.Name == tt.args.leader {
					leader = ldr
				}
				if leaderInd.Visible() != (ldr.Name == tt.args.leader) {
					t.Fatalf("leader \"%s\".Visible() == %t, want %t", ldr.Name,
						leaderInd.Visible(), ldr.Name == tt.args.leader)
				}
			}
			if reflect.DeepEqual(leader, skirmish.Leader{}) {
				if tt.wantErr {
					return
				}
				t.Fatalf("Leader %s was not found", tt.args.leader)
			}
			banner := ps.Hex(leader.Banner)
			ind := ps.Hex(leader.Indicator)
			barStroke := ps.Stroke{Size: 4, Color: banner}

			want.ResolveBG.SetColor(banner)
			want.Speed.SetStroke(barStroke, ps.ColorGray)
			want.SpeedBG.SetColor(ind)
			want.Damage.SetStroke(barStroke, ps.ColorWhite)

			defer want.Doc.Dump()
			got := new(mode, tt.args.template)
			defer got.Doc.Dump()
			_, _, _, err := got.SetLeader(tt.args.leader)
			if (err != nil) != wantErr {
				t.Errorf("got.SetLeader() err = %s, wanted %t", err, wantErr)
			}

			for _, a := range got.DeckInd.ArtLayers() {
				for _, b := range want.DeckInd.ArtLayers() {
					if a.Name() == b.Name() && a.Visible() != b.Visible() {
						t.Errorf("wanted:\n%+v\ngot:\n%+v", b, a)
					}
				}
			}

			if !reflect.DeepEqual(got.ResolveBG.Color.RGB(), want.ResolveBG.Color.RGB()) {
				t.Error("ResolveBG.Color doesn't match")
			}
			if !reflect.DeepEqual(got.Speed.Color.RGB(), want.Speed.Color.RGB()) {
				t.Error("Speed.Color doesn't match")
			}
			if !reflect.DeepEqual(got.Speed.Stroke, want.Speed.Stroke) {
				t.Error("Speed.Stroke doesn't match")
			}
			if !reflect.DeepEqual(got.SpeedBG.Color.RGB(), want.SpeedBG.Color.RGB()) {
				t.Error("SpeedBG.Color doesn't match")
			}
			if !reflect.DeepEqual(got.Damage.Color.RGB(), want.Damage.Color.RGB()) {
				t.Error("Damage.Color doesn't match")
			}
			if !reflect.DeepEqual(got.Damage.Stroke, want.Damage.Stroke) {
				t.Error("Damage.Stroke doesn't match")
			}
			if !reflect.DeepEqual(got.DeckInd, want.DeckInd) {
				log.Println("Sets don't match.")
			}
		})
	}
}

func LoadTest(filename string) string {
	path := filepath.Join(skirmish.Cfg.PS.Dir, "_test", filename)
	return path
}
