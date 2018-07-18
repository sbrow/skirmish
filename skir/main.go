//go:generate ./mkalldocs.sh

package main

import (
	"flag"
	"log"
	"os"

	"github.com/sbrow/skirmish/skir/internal/base"
	"github.com/sbrow/skirmish/skir/internal/card"
	"github.com/sbrow/skirmish/skir/internal/dump"
	"github.com/sbrow/skirmish/skir/internal/export"
	"github.com/sbrow/skirmish/skir/internal/help"
	"github.com/sbrow/skirmish/skir/internal/ps"
	"github.com/sbrow/skirmish/skir/internal/recover"
	"github.com/sbrow/skirmish/skir/internal/sql"
	"github.com/sbrow/skirmish/skir/internal/version"
)

func init() {
	base.Commands = []*base.Command{
		card.CmdCard,
		dump.CmdDump,
		export.CmdExport,
		ps.CmdPS,
		recover.CmdRecover,
		sql.CmdSQL,
		version.CmdVersion,
	}
}

func main() {
	flag.Usage = base.Usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) < 1 {
		base.Usage()
	}
	if args[0] == "help" {
		help.Help(args[1:])
		return
	}
	for _, cmd := range base.Commands {
		if cmd.Name() == args[0] && cmd.Runnable() {
			cmd.Flag.Usage = func() { cmd.Usage() }
			if cmd.CustomFlags {
				args = args[1:]
			} else {
				err := cmd.Flag.Parse(args[1:])
				if err != nil {
					log.Println(err)
				}
				args = cmd.Flag.Args()
			}
			cmd.Run(cmd, args)
			base.Exit()
			return
		}
	}
}

func init() {
	base.Usage = mainUsage
}

func mainUsage() {
	help.PrintUsage(os.Stderr)
	os.Exit(2)
}
