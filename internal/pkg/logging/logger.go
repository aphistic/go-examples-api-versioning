package logging

import (
	"fmt"
)

type Logger interface {
	Log(message string, args ...interface{})
}

// NilLogger is a Logger that doesn't do anything and can be used
// as a sane default for assuming a Logger exists.
type NilLogger struct{}

// Use the compile-time type checking to make sure
// NilLogger is a Logger
var _ Logger = &NilLogger{}

func NewNilLogger() *NilLogger {
	return &NilLogger{}
}

func (nl *NilLogger) Log(message string, args ...interface{}) {
	// Do nothing since this is the nil logger
}

// StdoutLogger is a Logger that just writes a log message to STDOUT.
type StdoutLogger struct{}

// Use the compile-time type checking to make sure
// StdoutLogger is a Logger
func NewStdoutLogger() *StdoutLogger {
	return &StdoutLogger{}
}

func (nl *StdoutLogger) Log(message string, args ...interface{}) {
	fmt.Printf(message+"\n", args...)
}
