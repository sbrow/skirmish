package recover

import (
	"fmt"
	"path/filepath"

	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/skir/internal/base"
)

// CmdRecover runs skirmish.Recover.
var CmdRecover = &base.Command{
	UsageLine: "recover",
	Short:     "reload the database from disk",
	Long: `Recover runs the skirmish_db.sql file from the dreamkeepers-dat repository
on the database, effectively resetting it to the most recently saved state. To overwrite this
file, see 'skir dump'.`,
	Run: func(cmd *base.Command, args []string) {
		if len(args) > 0 {
			cmd.Usage()
			base.Exit()
		}
		path := filepath.Join(skirmish.Cfg.DB.Dir, "skirmish_db.sql")
		result, err := skirmish.Recover(path)
		if err != nil {
			base.Fatalf("%s", err)
		}
		fmt.Printf("%v\n", result)
		base.Exit()
	},
}
