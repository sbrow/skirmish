package skirmish

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/go-yaml/yaml"
)

// Config for "Current" test.
var current = Config{
	PS: cfgPS{
		Dir:     `F:\GitLab\dreamkeepers-psd`,
		Deck:    `Template009.1.psd`,
		NonDeck: `Template009.1h.psd`,
	},
	DB: cfgDB{
		Dir:  `F:\GitLab\dreamkeepers-dat`,
		Name: "skirmish",
		Host: "localhost",
		Port: 5432,
		User: "sbrow",
		SSL:  false,
	},
}

func TestConf_Load(t *testing.T) {
	tmpCfg := func(name string, cfg Config) *os.File {
		f, err := ioutil.TempFile(os.Getenv("TEMP"), name)
		if err != nil {
			t.Fatal(err)
		}
		f.Close()
		err = cfg.Save(f.Name())
		if err != nil {
			t.Fatal(err)
		}
		return f
	}

	FCurrent := tmpCfg("currentConfig", current)
	def := *DefaultCfg()
	FDef := tmpCfg("defaultConfig", def)

	tests := []struct {
		name    string
		path    string
		want    Config
		wantErr bool
	}{
		{"None", FDef.Name(), def, false},
		{"Current", FCurrent.Name(), current, false},
		// TODO(sbrow): Re-enable TestConf_Load tests. [Issue](https://github.com/sbrow/skirmish/issues/34)
		// {"Default", ".default_config.yml", def, false},
		// {"DefaultNoConfig", "config.yml", *DefaultCfg(), false},
		// {"Default_NoArgs", "", current, false},

		{"FakeConfig", "fake_config.yml", Config{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &Config{}
			err := got.Load(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadCfg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("LoadCfg() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestConf_Save(t *testing.T) {
	type args struct {
		path string
		temp bool
	}
	tests := []struct {
		name string
		args
		wantErr bool
	}{
		{"Temp", args{".test.yml", true}, false},

		{"WriteProtected", args{"/config.yml", true}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Cfg.Save(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Conf.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			f, err := os.Open(tt.args.path)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Fatal(err)
			}
			defer os.Remove(f.Name())
			defer f.Close()
			got, err := ioutil.ReadAll(f)
			if err != nil {
				t.Fatalf(`Could not read from file "%s": "%s"`, tt.path, err)
			}
			want, err := yaml.Marshal(Cfg)
			if err != nil {
				t.Fatalf(`Error when marshalling "%v": "%s`, Cfg, err)
			}
			if string(got) != string(want) {
				t.Errorf("Conf.Save() = %v, want %v", string(got), string(want))
			}
		})
	}
}

func TestConfOverwrite(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	old := *Cfg
	*Cfg = current
	if err := Cfg.Save("config.yml"); err != nil {
		t.Fatal(err)
	}
	want := *Cfg
	cmd := exec.Command("vgo", "run", filepath.Join(dir, "skir", "main.go"), "version")
	if err := cmd.Run(); err != nil {
		t.Error(err)
	}
	if Cfg == nil || *Cfg != want {
		t.Errorf("wanted: %v\ngot: %v", want, *Cfg)
	}
	old.Save("config.yml")
}
