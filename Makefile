# Program parameters
VERSION=$(shell git describe --tags)

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GOGENERATE=$(GOCMD) generate

# BINARY_NAME=mybinary
# BINARY_WIN=$(BINARY_NAME).exe


default: version fmt build test clean install docs

all: version fmt build test clean install docs lint release

fast: version test install docs

version:
	sed -i -r 's/(const Version = ")([^"]*)(")/\1$(VERSION)\3/' ./skir/internal/version/version.go

fmt:
	goimports -w ./..	
	gofmt -s -w ./..

build: fmt
	$(GOBUILD) -v ./...

test: fmt
	$(GOTEST) -v -coverprofile=cover.out ./...

clean:
	$(GOCLEAN) ./...
	rm -f cover.out
	rm -f ./Unreal_JSONs/*.*

lint: fmt
	gometalinter.v2 --enable-gc --cyclo-over=15 ./...

install: fmt
	$(GOINSTALL) ./...

docs: fmt
	todos work
	$(GOGENERATE) ./...

release:
	echo \# v0.13.1 Release Notes > RELEASE.md
	git log v0.12.1...v0.13.0 --format=%s\n%b >> RELEASE.md