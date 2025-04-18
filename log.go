package wwabago

import (
	"errors"
	stdlog "log"
	"os"
)

// BotLogger is an interface that represents the required methods to log data.
//
// Instead of requiring the standard logger, we can just specify the methods we
// use and allow users to pass anything that implements these.
type wabaClientLogger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

var log wabaClientLogger = stdlog.New(os.Stderr, "", stdlog.LstdFlags)

// SetLogger specifies the logger that the package should use.
func SetLogger(logger wabaClientLogger) error {
	if logger == nil {
		return errors.New("logger is nil")
	}
	log = logger
	return nil
}