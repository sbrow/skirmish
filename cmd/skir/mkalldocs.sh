#!/bin/bash
# Copyright 2012 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e

go build -o go.latest
./go.latest help documentation >alldocs.go
gofmt -w alldocs.go
rm go.latest

godoc2md -template ../../.doc.template github.com/sbrow/skirmish/cmd/skir > README.md
# godoc2md github.com/sbrow/skirmish/cmd/skir > README.md