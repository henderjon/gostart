package main

import (
	"os"

	"github.com/henderjon/errors"
	"github.com/henderjon/logger"
)

func main() {
	infoLogger := logger.New(os.Stderr, logger.Info)

	e := errors.New("is this this really an error", errors.New("hello, world"))
	infoLogger.Println(e)
}
