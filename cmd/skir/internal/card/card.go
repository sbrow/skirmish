package card

import (
	"fmt"
	"log"
	"strings"

	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/cmd/skir/internal/base"
)

var CmdCard = &base.Command{
	UsageLine: "card [card name]",
	Short:     "return the text of a given card",
	Long:      ``,
	Run: func(cmd *base.Command, args []string) {
		var card skirmish.Card
		name := args[0]
		if name == "" {
			name = strings.Join(args, " ")
		}
		card, err := skirmish.Load(name)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(card)
	},
}
