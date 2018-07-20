package skirmish

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/go-yaml/yaml"
)

func init() {
	_, file, _, _ := runtime.Caller(0) // TODO(sbrow): catch runtime.Caller error in config.go/init

	Cfg = &Config{}
	if err := Cfg.Load(filepath.Join(filepath.Dir(file), "config.yml")); err != nil {
		log.Println(err)
	}
	ImageDir = filepath.Join(os.Getenv("SK_PS"), "Images")
	Cfg.DB.Dir = filepath.Join(Cfg.DB.Dir)
}

type cfgDB struct {
	Dir  string // The default directory to call Load() and Recover() in.
	Host string // The server ip address.
	Port int    // the server port.
	Name string // The name of the database.
	User string // The user name to login as.
	SSL  bool   // Whether or not to use SSL.
}

type cfgPS struct {
	Dir     string // The directory of the Photoshop files.
	Deck    string // The name of the Deck Card Photoshop template file.
	NonDeck string `yaml:"non_deck"` // The name of the Non-Deck Card Photoshop template file.
}

// Cfg holds the currently loaded configuration settings.
var Cfg *Config

// LocalDB holds the configuration for connecting to the default, local postgres database.
// It is primarily used for running tests.
var LocalDB = &Config{
	DB: cfgDB{
		Host: "localhost",
		Port: 5432,
		Name: "postgres",
		User: "postgres",
		SSL:  false,
	},
}

// Config holds various configuration values for the program,
// namely the directories of other relevant git repositories:
// "dreamkeepers-dat" and "dreamkeepers-psd".
type Config struct {
	// PS holds configuration values related to Photoshop.
	PS struct {
		Dir     string // The directory of the Photoshop files.
		Deck    string // The name of the Deck Card Photoshop template file.
		NonDeck string `yaml:"non_deck"` // The name of the Non-Deck Card Photoshop template file.
	} `yaml:"photoshop"`
	// DB Holds configuration values related to the PSQL database.
	DB struct {
		Dir  string // The default directory to call Load() and Recover() in.
		Host string // The server ip address.
		Port int    // the server port.
		Name string // The name of the database.
		User string // The user name to login as.
		SSL  bool   // Whether or not to use SSL.
	} `yaml:"database"`
}

// DefaultCfg returns a full, basic Config.
func DefaultCfg() *Config {

	user, err := user.Current()
	if err != nil {
		log.Println("Couldn't get current user.")
	}
	var home string
	if user != nil {
		home = user.HomeDir
	}
	cfg := &Config{}
	cfg.PS = cfgPS{
		Dir:     filepath.Join(home, "dreamkeepers-psd"),
		Deck:    "Template009.1.psd",
		NonDeck: "Template009.1h.psd",
	}
	host := os.Getenv("PSQL_HOST")
	if host == "" {
		host = "localhost"
	}
	cfg.DB = cfgDB{
		Dir:  filepath.Join(home, "dreamkeepers-dat"),
		Host: host,
		Port: 5432,
		Name: "skirmish",
		User: "guest",
		SSL:  false,
	}
	return cfg
}

// DBArgs returns c.DB as a list of args that can be passed to Connect().
func (c Config) DBArgs() (host string, port int, DBName, user, sslmode string) {
	modes := map[bool]string{
		false: "disable",
		true:  "require",
	}
	d := c.DB
	return d.Host, d.Port, d.Name, d.User, modes[d.SSL]
}

// Load loads Config data from a YAML file at the given path.
// It returns an error if the file was not found.
func (c *Config) Load(path string) error {
	if path == "" {
		path = "config.yml"
	}
	f, err := os.Open(path)
	if err != nil {
		if filepath.Base(path) == "config.yml" {
			*c = *DefaultCfg()
			log.Println("Creating a new log file from default values.")
			return c.Save("config.yml")
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

// Save saves the Config to a YAML file at the given path.
func (c *Config) Save(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	if _, err = f.Write(data); err != nil {
		return err
	}
	return f.Close()
}
