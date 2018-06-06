package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

// logger should be global
var logger *log.Logger

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

	// set globally via linker during compilation; in version.go
	params.Build.Timestamp = getBuildTimestamp()
	params.Build.Version = getBuildVersion()

	logger = log.New(ioutil.Discard, "null ", log.Lshortfile|log.LUTC|log.LstdFlags)
	if params.Build.Debug {
		logger = log.New(os.Stderr, "debug ", log.Lshortfile|log.LUTC|log.LstdFlags)
	}

	return params
}
