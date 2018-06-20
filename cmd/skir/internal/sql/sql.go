package sql

import (
	"github.com/sbrow/skirmish/cmd/skir/internal/base"
)

// TODO: Fix
var CmdSql = &base.Command{
	UsageLine: "sql [PSQL query]",
	Short:     "query the database",
	Long:      `'Skir sql' queries the database for any desired information.`,
	Run: func(cmd *base.Command, args []string) {
		base.Run(append([]string{"psql", "-U", "postgres", "-c"}, args...))
	},
}
