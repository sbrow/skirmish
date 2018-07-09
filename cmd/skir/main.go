//go:generate ./mkalldocs.sh
//	- Query the database
//  - Get a card from the database in X format.
// 	- Get all cards from the database into X format.
//
// PS
//
// Fills out a document (psd) template with data from a dataset (csv) file and
//
// 		skir ps
//		skir ps -card=$CARDNAME // Builds the psd for the given card
//		skir ps -deck=$DECKNAME
//		skir ps action $ACTIONNAME

package main

import (
	"flag"
	"log"
	"os"

	"github.com/sbrow/skirmish/cmd/skir/internal/base"
	"github.com/sbrow/skirmish/cmd/skir/internal/card"
	"github.com/sbrow/skirmish/cmd/skir/internal/export"
	"github.com/sbrow/skirmish/cmd/skir/internal/help"
	"github.com/sbrow/skirmish/cmd/skir/internal/sql"
	"github.com/sbrow/skirmish/cmd/skir/internal/version"
)

func init() {
	base.Commands = []*base.Command{
		export.CmdExport,
		card.CmdCard,
		sql.CmdSql,
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
