// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package version implements the ``go version'' command.
package version

import (
	"fmt"

	"github.com/sbrow/skirmish/cmd/skir/internal/base"
)

var CmdVersion = &base.Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print skir version",
	Long:      `Version prints the skir version.`,
}

const Version = "v0.11.1-7-ge8872bb"

func runVersion(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}
	fmt.Printf("skir version %s", Version)
}
