package main

import (
	"github.com/grahambrooks/xpa/pkg/xpa"
	"os"
)

type OutputFormatter struct {
	ProgressFn    func(format string, args ...interface{})
	ErrorFn       func(format string, args ...interface{})
	ObservationFn func(observation *xpa.Analysis)
}

func (of *OutputFormatter) Progress(format string, args ...interface{}) {
	if of.ProgressFn != nil {
		of.ProgressFn(format, args...)
	}
}

func (of *OutputFormatter) Error(format string, args ...interface{}) {
	if of.ErrorFn != nil {
		of.ErrorFn(format, args...)
	}
}

func (of *OutputFormatter) Observation(observation *xpa.Analysis) {
	if of.ObservationFn != nil {
		of.ObservationFn(observation)
	}
}

func isTTY() bool {
	fileInfo, _ := os.Stdout.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}
