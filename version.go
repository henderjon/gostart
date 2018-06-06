package main

// these vars are built at compile time, DO NOT ALTER
var (
	// Version adds build information
	buildVersion string
	// BuildTimestamp adds build information
	buildTimestamp string
)

func getBuildVersion() string {
	return buildVersion
}

func getBuildTimestamp() string {
	return buildTimestamp
}

// assuming git ...
// HEAD=$(shell git describe --dirty --long --tags 2> /dev/null  || git rev-parse --short HEAD)
// TIMESTAMP=$(shell date '+%Y-%m-%dT%H:%M:%S %z %Z')
// LDFLAGS="-X main.buildVersion=$(HEAD) -X 'main.buildTimestamp=$(TIMESTAMP)'"
// go build -ldflags $(LDFLAGS)
