PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test

all: test

test:
	go test -v $(PKGS)
