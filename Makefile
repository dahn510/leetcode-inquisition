SHELL = /bin/sh
SOURCES := main.go

prefix ?= /usr/local
bindir ?= $(prefix)/bin

.PHONY: all
all: build

.PHONY: install
install: build
	install -Dm755 ./leetcode-inquisition $(DESTDIR)$(bindir)/leetcode-inquisition

.PHONY: build
build: $(SOURCES)
	go build

.PHONY: clean
clean:
	rm -f ./leetcode-inquisition
