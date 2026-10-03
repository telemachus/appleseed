.DEFAULT_GOAL := test

fmt:
	golangci-lint fmt --no-config -Egofumpt

# Pinned to go1.27.1: staticcheck can't read go1.27.2's export data yet.
# Remove once a staticcheck release supports it.
staticcheck:
	GOTOOLCHAIN=go1.27.1 staticcheck ./...

revive:
	revive -config revive.toml ./...

golangci:
	golangci-lint run

lint: staticcheck revive golangci

test:
	go test -shuffle on .

testv:
	go test -shuffle on -v .

testr:
	go test -race -shuffle on .

build: lint
	go build .

clean:
	go clean

.PHONY: fmt staticcheck revive golangci lint build test testv testr clean
