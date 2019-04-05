package gostart

import (
	"io/ioutil"
	"log"
	"os"
)

// logger should be global
var logger *log.Logger = log.New(ioutil.Discard, "null ", log.Lshortfile|log.LUTC|log.LstdFlags)

// New creates a new debug logger
func getDebugLogger(stderr bool) *log.Logger {
	if stderr {
		logger = log.New(os.Stderr, "debug ", log.Lshortfile|log.LUTC|log.LstdFlags)
	}
	return logger
}
