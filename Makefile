TARGET   ?= alfred
PREFIX   ?= /usr/local
GO       ?= go
GOFLAGS  ?=

all: build

build: resources
	$(GO) $(GOFLAGS) build -o $(TARGET) tool/alfred.go

install:
	@cp alfred $(PREFIX)/bin/

.PHONY: all build build_with_race_detector resources install docker test html_coverage benchmark fmt clean