TARGET   ?= alfred
PREFIX   ?= /usr/local
GO       ?= go
GOFLAGS  ?=

build-cli: resources
	@echo "Building Alfred CLI tool..."
	$(GO) $(GOFLAGS) build -o $(TARGET) tool/alfred.go

install-cli:
	@echo "Installing Alfred CLI tool..."
	@cp alfred $(PREFIX)/bin/

run:
	@echo "Starting Go-Web framework"
	$(GO) $(GOFLAG) run .

build:
	@echo "Building Go-Web framework"
	$(GO) $(GOFLAG) build -o goweb .

.PHONY: all build build_with_race_detector resources install docker test html_coverage benchmark fmt clean