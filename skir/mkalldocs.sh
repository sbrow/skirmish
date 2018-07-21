#!/bin/bash
# Copyright 2012 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e

go build -o skir.latest
./skir.latest help documentation >alldocs.go
gofmt -w alldocs.go
rm skir.latest

godoc2md -template ../.doc.template github.com/sbrow/skirmish/skir > README.md
branch=$(git rev-parse --abbrev-ref HEAD)
sed -i -r "s/branch=master/branch=$branch/g" README.md