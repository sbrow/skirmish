package dump

import (
	"path/filepath"

	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/skir/internal/base"
)

var CmdDump = &base.Command{
	UsageLine: "dump",
	Short:     "save the current database to disk",
	Long: `Dump saves the current state of the database to the "skirmish_db.sql"
file in the dreamkeepers-dat repository.`,
	Run: func(cmd *base.Command, args []string) {
		if len(args) > 0 {
			cmd.Usage()
			base.Exit()
		}
		path := filepath.Join(skirmish.Cfg.DB.Dir, "skirmish_db.sql")
		skirmish.Dump(path)
	},
}
