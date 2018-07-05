package skirmish

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Database struct {
	Dir string `yaml:"dir"`
}
type PS struct {
	Dir     string
	Deck    string
	NonDeck string `yaml:"non_deck"`
}

var Config *Conf

type Conf struct {
	PS       `yaml:"photoshop"`
	Database `yaml:"database"`
}

func DefaultCfg() *Conf {
	cfg := &Conf{}

	user, err := user.Current()
	if err != nil {
		log.Println("Couldn't get current user.")

	}
	var home string
	if user != nil {
		home = user.HomeDir
	}
	cfg.Database.Dir = filepath.Join(home, "dreamkeepers-dat")
	cfg.PS = PS{
		Dir:     filepath.Join(home, "dreamkeepers-psd"),
		Deck:    "Template009.1.psd",
		NonDeck: "Template009.1h.psd",
	}
	return cfg
}

func (c *Conf) Load(path string) error {
	if path == "" {
		path = "config.yml"
	}
	f, err := os.Open(path)
	if err != nil {
		log.Printf("config file \"%s\" was not found.", path)
		c = DefaultCfg()
		if path == "config.yml" {
			log.Println("Creating a new log file from default values.")
			return c.Save(path)
		}
		return err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &c)
	return err
}

func (c *Conf) Save(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	// defer f.Close()

	_, err = f.Write(data)
	f.Close()
	return err
}

func (c *Conf) setEnvs() error {
	if err := os.Setenv("SK_PS", c.PS.Dir); err != nil {
		return err
	}
	if err := os.Setenv("SK_SQL", c.Database.Dir); err != nil {
		return err
	}
	return nil
}
func init() {
	Config = &Conf{}
	if err := Config.Load(""); err != nil {
		log.Println(err)
	}
}
