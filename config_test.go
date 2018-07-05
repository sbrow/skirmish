package skirmish

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/go-yaml/yaml"
)

func check(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func TestConf_Load(t *testing.T) {
	curr := Conf{
		PS: PS{
			Dir:     `F:\GitLab\dreamkeepers-psd`,
			Deck:    `Template009.1.psd`,
			NonDeck: `Template009.1h.psd`,
		},
		Database: Database{Dir: `F:\GitLab\dreamkeepers-dat`},
	}
	def := *DefaultCfg()
	tests := []struct {
		name    string
		path    string
		want    Conf
		wantErr bool
	}{
		{"Current", "config.yml", curr, false},
		{"Default", ".default_config.yml", def, false},
		// {"DefaultNoConfig", "config.yml", *DefaultCfg(), false},
		// TODO(sbrow): Fix TestConf_Load.
		{"Default_NoArgs", "", curr, false},

		{"FakeConfig", "fake_config.yml", Conf{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &Conf{}
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
		// {"Current", args{"config.yml", false}, false},
		{"Default", args{".default_config.yml", false}, false},
		{"Temp", args{".test.yml", true}, false},

		{"WriteProtected", args{"/config.yml", true}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.temp {
				os.Remove(tt.args.path)
				defer os.Remove(tt.args.path)
			} else {
				cpy := copyFile(tt.args.path)
				defer os.Remove(cpy)
				defer copyFile(cpy)
			}
			if err := Config.Save(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Conf.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			f, err := os.Open(tt.args.path)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Fatal(err)
			}
			defer f.Close()
			got, err := ioutil.ReadAll(f)
			if err != nil {
				t.Fatalf(`Could not read from file "%s": "%s"`, tt.path, err)
			}
			want, err := yaml.Marshal(Config)
			if err != nil {
				t.Fatalf(`Error when marshalling "%s": "%s`, Config, err)
			}
			if string(got) != string(want) {
				t.Errorf("Conf.Save() = %v, want %v", string(got), string(want))
			}
		})
	}
}

func copyFile(path string) (pathToCopy string) {
	filename := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	_, _ = os.Stat(filename)
	if strings.Index(path, "_copy") == -1 {
		pathToCopy = strings.Replace(path, filename, filename+"_copy", 1)
	} else {
		pathToCopy = strings.Replace(path, filename, strings.TrimSuffix(filename, "_copy"), 1)
	}
	if err := os.Rename(path, pathToCopy); err != nil {
		log.Println(err)
	}
	return pathToCopy
}
func TestConf_SetEnvs(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args
		ps      string
		db      string
		wantErr bool
	}{
		{"Current", args{"config.yml"}, `F:\GitLab\dreamkeepers-psd`, `F:\GitLab\dreamkeepers-dat`, false},
		{"Default", args{".default_config.yml"}, filepath.Join(user.HomeDir, "dreamkeepers-psd"), filepath.Join(user.HomeDir, "dreamkeepers-dat"), false},

		{"NonExistent", args{"fake_config.yml"}, filepath.Join(user.HomeDir, "dreamkeepers-psd"), filepath.Join(user.HomeDir, "dreamkeepers-dat"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Config.Load(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Conf.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := Config.setEnvs(); err != nil {
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
