BIN=gostart
HEAD=$(shell git describe --dirty --long --tags 2> /dev/null  || git rev-parse --short HEAD)
TIMESTAMP=$(shell date '+%Y-%m-%dT%H:%M:%S %z %Z')
DEPLOYMENT_PATH=bin/release/$(BIN)/$(BIN)-$(HEAD)

LDFLAGS="-X main.buildVersion=$(HEAD) -X 'main.buildTimestamp=$(TIMESTAMP)' -X 'main.compiledBy=$(shell go version)'" # `-s -w` removes some debugging info that might not be necessary in production (smaller binaries)

all: print

.PHONY: build
build: clean darwin64 linux64

.PHONY: dep
dep:
	go mod vendor

clean:
	rm -f $(BIN) $(BIN)-*

.PHONY: local
local:
	go build -ldflags $(LDFLAGS) -o $(BIN) ./cmd/

darwin64:
	env GOOS=darwin GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(BIN)-darwin64-$(HEAD) ./cmd/

linux64:
	env GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(BIN)-linux64-$(HEAD) ./cmd/

docker:
	env GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(BIN) ./cmd/
	docker build -t "$(BIN):$(HEAD)" .

print: build
	$(info aws s3 cp "$(BIN)-linux64-$(HEAD)" "$(DEPLOYMENT_PATH)" --sse AES256)

.PHONY: test
test:
	go test -coverprofile=coverage.out -covermode=count

.PHONY: race
race:
	go test -race

.PHONY: test-report
test-report: test
	go tool cover -html=coverage.out

