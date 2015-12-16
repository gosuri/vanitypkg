PREFIX = bin
PROG = vanitypkg

test:
	go test ./...

build:
	go build -o $(PREFIX)/$(PROG) ./cmd

.PHONY: build dist test clean
