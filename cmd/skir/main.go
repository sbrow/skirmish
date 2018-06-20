// PS
//
// Fills out a document (psd) template with data from a dataset (csv) file and
//
// 		skir ps
//		skir ps -card=$CARDNAME // Builds the psd for the given card
//		skir ps -deck=$DECKNAME
//		skir ps action $ACTIONNAME
//
// TODO: config file for non-programmers
//
//go:generate bash ./mkalldocs.sh

package main

import (
	"flag"
	"log"
	"os"

	"github.com/sbrow/skirmish/cmd/skir/internal/base"
	"github.com/sbrow/skirmish/cmd/skir/internal/build"
	"github.com/sbrow/skirmish/cmd/skir/internal/card"
	"github.com/sbrow/skirmish/cmd/skir/internal/help"
	"github.com/sbrow/skirmish/cmd/skir/internal/sql"
)

func init() {
	base.Commands = []*base.Command{
		build.CmdBuild,
		card.CmdCard,
		sql.CmdSql,
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
