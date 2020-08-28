TARGET   ?= alfred
PREFIX   ?= /usr/local
GO       ?= go
GOFLAGS  ?=

all: build

build: resources
	$(GO) $(GOFLAGS) build -o $(TARGET) tool/alfred.go

install:
	@cp alfred $(PREFIX)/bin/

gwf-run:
	$(GO) $(GOFLAG) run . server:run

gwf-build:
	$(GO) $(GOFLAG) build -o goweb .

.PHONY: all build build_with_race_detector resources install docker test html_coverage benchmark fmt clean