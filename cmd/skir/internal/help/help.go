// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package help implements the ``skir help'' command.
package help

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/sbrow/skirmish/cmd/skir/internal/base"
)

// Help implements the 'help' command.
func Help(args []string) {
	if len(args) == 0 {
		PrintUsage(os.Stdout)
		// not exit 2: succeeded at 'skir help'.
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: skir help command\n\nToo many arguments given.\n")
		os.Exit(2) // failed at 'skir help'
	}

	arg := args[0]

	// 'skir help documentation' generates doc.go.
	if arg == "documentation" {
		fmt.Println("// Copyright 2011 The Go Authors. All rights reserved.")
		fmt.Println("// Use of this source code is governed by a BSD-style")
		fmt.Println("// license that can be found in the LICENSE file.")
		fmt.Println()
		fmt.Println("// DO NOT EDIT THIS FILE. GENERATED BY mkalldocs.sh.")
		fmt.Println("// Edit the documentation in other files and rerun mkalldocs.sh to generate this one.")
		fmt.Println()
		buf := new(bytes.Buffer)
		PrintUsage(buf)
		usage := &base.Command{Long: buf.String()}
		tmpl(&commentWriter{W: os.Stdout}, documentationTemplate, append([]*base.Command{usage}, base.Commands...))
		fmt.Println("package main")
		return
	}

	for _, cmd := range base.Commands {
		if cmd.Name() == arg {
			tmpl(os.Stdout, helpTemplate, cmd)
			// not exit 2: succeeded at 'skir help cmd'.
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown help topic %#q. Run 'skir help'.\n", arg)
	os.Exit(2) // failed at 'skir help cmd'
}

var usageTemplate = `Skir is a tool for developing the Skirmish card game.

Usage:

	skir command [arguments]

The commands are:
{{range .}}{{if .Runnable}}
	{{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use "skir help [command]" for more information about a command.

`

// Additional help topics:
// {{range .}}{{if not .Runnable}}
// 	{{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

// Use "skir help [topic]" for more information about that topic.

// `

var helpTemplate = `{{if .Runnable}}usage: skir {{.UsageLine}}

{{end}}{{.Long | trim}}
`

var documentationTemplate = `{{range .}}{{if .Short}}{{.Short | capitalize}}

{{end}}{{if .Runnable}}Usage:

	skir {{.UsageLine}}

{{end}}{{.Long | trim}}


{{end}}`

// commentWriter writes a Go comment to the underlying io.Writer,
// using line comment form (//).
type commentWriter struct {
	W            io.Writer
	wroteSlashes bool // Wrote "//" at the beginning of the current line.
}

func (c *commentWriter) Write(p []byte) (int, error) {
	var n int
	for i, b := range p {
		if !c.wroteSlashes {
			s := "//"
			if b != '\n' {
				s = "// "
			}
			if _, err := io.WriteString(c.W, s); err != nil {
				return n, err
			}
			c.wroteSlashes = true
		}
		n0, err := c.W.Write(p[i : i+1])
		n += n0
		if err != nil {
			return n, err
		}
		if b == '\n' {
			c.wroteSlashes = false
		}
	}
	return len(p), nil
}

// An errWriter wraps a writer, recording whether a write error occurred.
type errWriter struct {
	w   io.Writer
	err error
}

func (w *errWriter) Write(b []byte) (int, error) {
	n, err := w.w.Write(b)
	if err != nil {
		w.err = err
	}
	return n, err
}

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace, "capitalize": capitalize})
	template.Must(t.Parse(text))
	ew := &errWriter{w: w}
	err := t.Execute(ew, data)
	if ew.err != nil {
		// I/O error writing. Ignore write on closed pipe.
		if strings.Contains(ew.err.Error(), "pipe") {
			os.Exit(1)
		}
		base.Fatalf("writing output: %v", ew.err)
	}
	if err != nil {
		panic(err)
	}
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[n:]
}

func PrintUsage(w io.Writer) {
	bw := bufio.NewWriter(w)
	tmpl(bw, usageTemplate, base.Commands)
	bw.Flush()
}
