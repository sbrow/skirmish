package deck

import (
	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/skir/internal/base"
	"github.com/sbrow/skirmish/skir/internal/card"
)

var CmdDeck = &base.Command{
	UsageLine: "deck [leader name]",
	Short:     "print information for every card in a deck",
	Long:      `Equivilant to many calls of 'card'`,
	Run:       Run,
}

func Run(cmd *base.Command, args []string) {
	rows, err := skirmish.Query("SELECT name from cards where cards.leader=$1 ORDER BY name ASC", args[0])
	if err != nil {
		base.Fatalf("%s", err)
	}

	var crd string
	for rows.Next() {
		if err := rows.Scan(&crd); err != nil {
			base.Fatalf("%s", err)
		}
		card.CmdCard.Run(card.CmdCard, []string{crd})
	}
}
