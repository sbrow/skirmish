// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package version implements the ``go version'' command.
package version

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/sbrow/skirmish/cmd/skir/internal/base"
)

var CmdVersion = &base.Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print skir version",
	Long:      `Version prints the skir version.`,
}

func runVersion(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}

	fmt.Printf("skir version %s", Version())
}

func Version() string {
	cmd := exec.Command("git", "describe", "--tags")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		base.Errorf("%v", err)
	}
	return out.String()
}
