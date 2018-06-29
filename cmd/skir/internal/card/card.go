package card

import (
	"flag"
	"fmt"
	"strings"

	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/cmd/skir/internal/base"
)

var formats = map[string]func(skirmish.Card) string{
	"string": func(c skirmish.Card) string {
		return fmt.Sprint(c)
	},
	"ue": func(c skirmish.Card) string {
		data, err := c.UEJSON(true)
		if err != nil {
			base.Fatalf(err.Error())
		}
		return string(data)
	},
	"xml": func(c skirmish.Card) string {
		data, err := c.XML()
		if err != nil {
			base.Fatalf(err.Error())
		}
		return string(data)
	},
}

var DefaultFormat = "string"

var CmdCard = &base.Command{
	UsageLine: "card [card name]",
	Short:     "return the text of a given card",
	Long:      ``,
}

var format *string

func init() {
	flags := flag.NewFlagSet("flags", flag.ContinueOnError)
	format = flags.String("fmt", "string", "fmt determines which format to output in.")
	CmdCard.Flag = *flags
	CmdCard.Run = CardRun
}

func CardRun(cmd *base.Command, args []string) {
	if err := cmd.Flag.Parse(args); err != nil {
		base.Errorf(err.Error())
	}
	args = cmd.Flag.Args()
	var card skirmish.Card
	name := strings.Join(args, " ")
	if name == "" {
		name = strings.Join(args, " ")
	}
	card, err := skirmish.Load(name)
	if err != nil {
		base.Fatalf(err.Error())
	}
	fmt.Println(formats[*format](card))
}
