BIN=gostart
HEAD=$(shell git describe --dirty --long --tags 2> /dev/null  || git rev-parse --short HEAD)
TIMESTAMP=$(shell date '+%Y-%m-%dT%H:%M:%S %z %Z')
DEPLOYMENT_PATH=bin/release/$(BIN)/$(BIN)-$(HEAD)

LDFLAGS="-X main.buildVersion=$(HEAD) -X 'main.buildTimestamp=$(TIMESTAMP)'"

all: print

.PHONY: build
build: darwin64 linux64

.PHONY: dep
dep:
	go mod vendor

clean:
	rm -f $(BIN) $(BIN)-*

.PHONY: local
local:
	go build -ldflags $(LDFLAGS) -o $(BIN)

darwin64:
	env GOOS=darwin GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(BIN)-darwin64-$(HEAD)
	tar czvf $(BIN)-darwin64-$(HEAD).tgz $(BIN)-darwin64-$(HEAD)

linux64:
	env GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(BIN)-linux64-$(HEAD)
	tar czvf $(BIN)-linux64-$(HEAD).tgz $(BIN)-linux64-$(HEAD)

print: build
	$(info aws s3 cp "$(BIN)-linux64-$(HEAD)" "$(DEPLOYMENT_PATH)" --sse AES256)

.PHONY: test
test:
	go test -coverprofile=coverage.out -covermode=count

.PHONY: test-report
test-report: test
	go tool cover -html=coverage.out
