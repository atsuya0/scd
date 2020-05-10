test := go test -v -cover -parallel 4

.PHONY: build install test format

build: format
	@go build

install:
	@go install

test:
	@$(test) ./cmd

format:
	@goimports -w cmd main.go
