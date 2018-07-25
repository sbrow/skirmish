package ps

import (
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	ps "github.com/sbrow/ps/v2"
	"github.com/sbrow/skirmish"
)

func Skip(t *testing.T) {
	switch {
	case runtime.GOOS != "windows":
		t.Skip("Photoshop is likely not installed")
	case testing.Short():
		t.Skip("Tests involving Photoshop take a long time")
	}
}

func Test_template_SetLeader(t *testing.T) {
	Skip(t)
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
		{"Invalid Leader", args{
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
					if tt.wantErr {
						t.Fatalf("leader \"%s\".Visible() = %t, want %t", ldr.Name,
							leaderInd.Visible(), ldr.Name == tt.args.leader)
					}
					return
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
			if (err != nil) != tt.wantErr {
				t.Errorf("got.SetLeader() err = %s, wanted %t", err, tt.wantErr)
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

func Test_template_Path(t *testing.T) {
	ChaoticBlast := skirmish.NewDeckCard()
	name := "Chaotic Blast"
	leader := "Bast"
	copies := 3
	ChaoticBlast.SetName(&name)
	ChaoticBlast.SetLeader(&leader)
	ChaoticBlast.SetRarity(&copies)

	Bast := skirmish.NewDeckCard()
	copies2 := 1
	Bast.SetName(&leader)
	Bast.SetRarity(&copies2)
	tests := []struct {
		name string
		t    *template
		want string
	}{
		{"Chaotic Blast", &template{Card: ChaoticBlast, Mode: PrintMode},
			filepath.Join(skirmish.Cfg.PS.Dir, "Decks", "Bast", "Chaotic Blast_1")},
		{"Chaotic Blast UE", &template{Card: ChaoticBlast, Mode: UEMode},
			filepath.Join(skirmish.Cfg.PS.Dir, "Card_Decks", "Bast", "3x_Chaotic_Blast")},
		{"Bast", &template{Card: Bast, Mode: PrintMode},
			filepath.Join(skirmish.Cfg.PS.Dir, "Decks", "Heroes", "Bast_1")},
		{"BastUE", &template{Card: Bast, Mode: UEMode},
			filepath.Join(skirmish.Cfg.PS.Dir, "Card_Decks", "Heroes", "1x_Bast")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Path(); got != tt.want {
				t.Errorf("template.Path() = %v, want %v", got, tt.want)
			}
		})
	}
}
