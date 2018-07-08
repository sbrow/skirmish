package sql

import (
	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/cmd/skir/internal/base"
)

// TODO(sbrow): Fix
var CmdSql = &base.Command{
	UsageLine: "sql [PSQL query]",
	Short:     "query the database",
	Long:      `'Skir sql' queries the database for any desired information.`,
	Run: func(cmd *base.Command, args []string) {
		dbname := skirmish.DefaultCfg().DB.Name
		user := skirmish.DefaultCfg().DB.User
		base.Run(append([]string{"psql", "-d", dbname, "-U", user, "-c"}, args...))
	},
}
