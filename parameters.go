package gostart

import (
	"flag"
	"fmt"
	"os"
)

// BuildParams is a set of values passed at runtime to define the operation params of the application
type BuildParams struct {
	Timestamp string
	Version   string
	Compiler  string
	Debug     bool
}

// GetOptParameters is a set of values passed at runtime to define the operation params of the application
type GetOptParameters struct {
	Build BuildParams
	Help  bool
}

const doc = `
%s is a simple analytics tool built by myON for myON.

version:  %s
compiled: %s
built:    %s

Usage: %s -cookie-salt <jwt-secret> [option [option]...]

Options:
`

// GetParams parses CLI args into values used by the application
func GetParams(buildVersion, buildTimestamp, compiledBy string) *GetOptParameters {
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			doc,
			os.Args[0],
			buildVersion,
			compiledBy,
			buildTimestamp,
			os.Args[0],
		)
		flag.PrintDefaults()
	}

	params := &GetOptParameters{}
	flag.BoolVar(&params.Build.Debug, "debug", false, "once more, with feeling")
	flag.BoolVar(&params.Help, "help", false, "show this message")
	flag.Parse()

	if params.Help {
		fmt.Println("built:", buildTimestamp)
		fmt.Println("version:", buildVersion)
		flag.PrintDefaults()
		os.Exit(0)
	}

	// value, ok := os.LookupEnv("")

	// set globally via linker during compilation; in version.go
	params.Build.Timestamp = buildTimestamp
	params.Build.Version = buildVersion
	params.Build.Compiler = compiledBy

	// only do this once
	if params.Help {
		flag.Usage()
		os.Exit(1)
	}

	return params
}
