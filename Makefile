# Program parameters
VERSION=$(shell git describe --tags)
LAST_VERSION=v0.13.1

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GOGENERATE=$(GOCMD) generate

LDFLAGS="-X github.com/sbrow/skirmish.Version=$(VERSION)"

# BINARY_NAME=mybinary
# BINARY_WIN=$(BINARY_NAME).exe

default: fmt build test clean install docs

all: default lint release

fast: test install docs

fmt:
	goimports -w ./..	
	gofmt -s -w ./..

build: fmt
	$(GOBUILD) -v -ldflags $(LDFLAGS) ./...

test: fmt
	$(GOTEST) -v -coverprofile=cover.out ./...

clean:
	$(GOCLEAN) ./...
	rm -f ./*.out
	rm -f ./*/*.out

	rm -f ./*.test
	rm -f ./*/*.test

	rm -f ./Unreal_JSONs/*.json

lint: fmt
	gometalinter.v2 --enable-gc --cyclo-over=15 ./...

install:
	$(GOINSTALL) -ldflags=$(LDFLAGS) ./...

docs: fmt
	todos work
	$(GOGENERATE) ./...

release:
	echo "# $(VERSION) Release Notes" > RELEASE.md
	git log $(LAST_VERSION)...$(VERSION) --format=%s%b >> RELEASE.md