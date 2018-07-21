// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package version implements the ``go version'' command.
package version

import (
	"fmt"

	"github.com/sbrow/skirmish/skir/internal/base"
)

var CmdVersion = &base.Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print skir version",
	Long:      `Version prints the skir version.`,
}

const Version = "v0.13.1-1-g9b0626e"

func runVersion(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
		base.Exit()
	}
	fmt.Printf("skir version %s", Version)
}
