package card

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/sbrow/skirmish"
	"github.com/sbrow/skirmish/skir/internal/base"
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

// DefaultFormat is the format to use when none is passed.
var DefaultFormat = "string"

var CmdCard = &base.Command{
	UsageLine: "card [-fmt=[format]] [card name]",
	Short:     "show information about a specific card",
	Long: `Card prints data for the given card to standard output.
	
The -fmt flag can be used to alter the output format.`,
}

var format *string

func init() {
	flags := flag.NewFlagSet("flags", flag.ContinueOnError)
	format = flags.String("fmt", "string", "fmt determines which format to output in.")
	CmdCard.Flag = *flags
	CmdCard.Run = Run
	CmdCard.Long += " The valid formats are:"
	fmts := make([]string, len(formats))
	i := 0
	for f := range formats {
		fmts[i] = f
		i++
	}
	sort.Strings(fmts)
	for i, format := range fmts {
		if i+1 == len(fmts) {
			CmdCard.Long += fmt.Sprintf(` and "%s".`, format)
		} else {
			CmdCard.Long += fmt.Sprintf(` "%s",`, format)
		}
	}
}

func Run(cmd *base.Command, args []string) {
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
