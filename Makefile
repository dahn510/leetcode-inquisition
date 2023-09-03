SHELL := bash
SOURCES := main.go

prefix = /usr/local
bindir = $(prefix)/bin


.PHONY: install
install: build
	install -Dm755 ./leetcode-inquisition $(bindir)/leetcode-inquisition

.PHONY: build
build: $(SOURCES)
	go build

.PHONY: clean
clean:
	rm -f ./leetcode-inquisition
