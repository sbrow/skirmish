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
	UsageLine: "card [-fmt] [card name]",
	Short:     "show information about a specific card",
	Long: `Card prints data for the given card to standard output.
	
The -fmt flag can be used to alter the output format.`,
}

var format *string

func init() {
	flags := flag.NewFlagSet("flags", flag.ContinueOnError)
	format = flags.String("fmt", "string", "fmt determines which format to output in.")
	CmdCard.Flag = *flags
	CmdCard.Run = CardRun
	CmdCard.Long += " The valid formats are:"
	i := 0
	for f := range formats {
		if i+1 == len(formats) {
			CmdCard.Long += fmt.Sprintf(` and "%s".`, f)
		} else {
			CmdCard.Long += fmt.Sprintf(` "%s",`, f)
		}
		i++
	}
}

func CardRun(cmd *base.Command, args []string) {
	if err := cmd.Flag.Parse(args); err != nil {
		base.Errorf(err.Error())
	}
	args = cmd.Flag.Args()
	if len(args) == 0 {
		base.Run([]string{"skir", "help", "card"})
		return
	}
	var card skirmish.Card
	name := strings.Join(args, " ")
	if name == "" {
		name = strings.Join(args, " ")
	}
	card, err := skirmish.Load(name)
	if err != nil {
		base.Fatalf(err.Error())
	}
	f, ok := formats[*format]
	if !ok {
		base.Fatalf("format \"%s\" was not found.", *format)
	}
	fmt.Println(f(card))
}
