.DEFAULT_GOAL := build

build :
	go build

test:
	go test ./...

test-cov:
	go test ./... -coverprofile=coverage.out

test-fuzz:
	go test ./internal/paths -fuzz=Fuzz --fuzztime 10s

test-verbose:
	go test ./... -v

test-race:
	go test ./... -race

clean:
	rm -f coverage.out
