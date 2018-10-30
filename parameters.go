package main

import (
	"flag"
	"fmt"
	"os"
)

type buildParams struct {
	Timestamp string
	Version   string
	Debug     bool
}

type getOptParameters struct {
	Build buildParams
	Help  bool
}

func getOptParams() *getOptParameters {
	params := &getOptParameters{}
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
	params.Build.Timestamp = getBuildTimestamp()
	params.Build.Version = getBuildVersion()

	return params
}
