package main

import (
	"github.com/henderjon/logger"
)

func main() {
	infoLogger := logger.NewStdoutLogger(true)
	infoLogger.Log("is this this really an error", logger.Here())
}
