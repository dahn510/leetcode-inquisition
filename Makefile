SHELL := bash
SOURCES := main.go

.PHONY: install
install: build
	install -Dm755 ./leetcode-inquisition $(DESTDIR)$(PREFIX)/leetcode-inquisition

.PHONY: build
build: $(SOURCES)
	go build

.PHONY: clean
clean:
	rm -f ./leetcode-inquisition
