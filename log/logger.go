package log

import (
	"log"
)

// Logger is an optional custom logger.
var Logger StdLogger = log.Default()

// StdLogger interface for Standard Logger.
type StdLogger interface {
	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(format string, args ...interface{})
	Print(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})
}

// Fatal writes a log entry.
// It uses Logger if not nil, otherwise it uses the default log.Logger.
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// Fatalf writes a log entry.
// It uses Logger if not nil, otherwise it uses the default log.Logger.
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}

// Print writes a log entry.
// It uses Logger if not nil, otherwise it uses the default log.Logger.
func Print(args ...interface{}) {
	Logger.Print(args...)
}

// Println writes a log entry.
// It uses Logger if not nil, otherwise it uses the default log.Logger.
func Println(args ...interface{}) {
	Logger.Println(args...)
}

// Printf writes a log entry.
// It uses Logger if not nil, otherwise it uses the default log.Logger.
func Printf(format string, args ...interface{}) {
	Logger.Printf(format, args...)
}

// Warnf writes a log entry.
func Warnf(format string, args ...interface{}) {
	Printf(format, args...)
}

// Infof writes a log entry.
func Infof(format string, args ...interface{}) {
	Printf(format, args...)
}
