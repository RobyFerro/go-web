TARGET   ?= julian
PREFIX   ?= /usr/local
GO       ?= go
GOFLAGS  ?=

build-cli: resources
	@echo "Building Julian CLI tool..."
	$(GO) $(GOFLAGS) build -o $(TARGET) tool/julian.go

install-cli:
	@echo "Installing Julian CLI tool..."
	@cp julian $(PREFIX)/bin/

run:
	@echo "Starting Go-Web framework"
	$(GO) $(GOFLAG) run .

build:
	@echo "Building Go-Web framework"
	$(GO) $(GOFLAG) build -o goweb .

.PHONY: all build build_with_race_detector resources install docker test html_coverage benchmark fmt clean