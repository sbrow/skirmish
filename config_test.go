package skirmish

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/go-yaml/yaml"
)

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
	// Config for "Current" test.
	curr := Config{
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
	currFile := tmpCfg("currentConfig", curr)
	def := *DefaultCfg()
	defFile := tmpCfg("defaultConfig", def)

	tests := []struct {
		name    string
		path    string
		want    Config
		wantErr bool
	}{
		{"None", defFile.Name(), def, false},
		{"Current", currFile.Name(), curr, false},
		// TODO(sbrow): Re-enable TestConf_Load tests.
		// {"Default", ".default_config.yml", def, false},
		// {"DefaultNoConfig", "config.yml", *DefaultCfg(), false},
		// {"Default_NoArgs", "", curr, false},

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

func TestConf_SetEnvs(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Fatal("Could not get current user.")
	}
	tests := []struct {
		name    string
		cfg     Config
		ps      string
		db      string
		wantErr bool
	}{
		{"Default", *DefaultCfg(), filepath.Join(user.HomeDir, "dreamkeepers-psd"), filepath.Join(user.HomeDir, "dreamkeepers-dat"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.setEnvVars(); (err != nil) != tt.wantErr {
				t.Errorf("Conf.setEnv() error = %v", err)
			}
			gotPS := os.Getenv("SK_PS")
			gotDB := os.Getenv("SK_SQL")
			if gotPS != tt.ps {
				t.Errorf("loadEnvs() = %v, want %v", gotPS, tt.ps)
			}
			if gotDB != tt.db {
				t.Errorf("loadEnvs() = %v, want %v", gotPS, tt.db)
			}
		})
	}
}
