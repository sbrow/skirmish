// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package version implements the ``go version'' command.
package version

import (
	"fmt"
	"os"

	"github.com/sbrow/skirmish/skir/internal/base"
)

var CmdVersion = &base.Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print skir version",
	Long:      `Version prints the installed version of skir.`,
}

const Version = "v0.13.0-30-g131a602"

func runVersion(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}
	fmt.Fprintf(os.Stderr, "skir version %s", Version)
}
