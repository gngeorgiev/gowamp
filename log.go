package gowamp

import (
	glog "log"
	"os"
)

var (
	logFlags = glog.Ldate | glog.Ltime | glog.Lshortfile
	logger   Logger
)

// Logger is an interface compatible with log.Logger.
type Logger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type noopLogger struct{}

func (n noopLogger) Println(v ...interface{})               {}
func (n noopLogger) Printf(format string, v ...interface{}) {}

// setup logger for package, noop by default
func init() {
	if os.Getenv("DEBUG") != "" {
		logger = glog.New(os.Stderr, "", logFlags)
	} else {
		logger = noopLogger{}
	}
}

// Debug changes the log output to stderr
func Debug() {
	logger = glog.New(os.Stderr, "", logFlags)
}

// DebugOff changes the log to a noop logger
func DebugOff() {
	logger = noopLogger{}
}

// SetLogger allows users to inject their own logger instead of the default one.
func SetLogger(l Logger) {
	logger = l
}

func logErr(err error) error {
	if err == nil {
		return nil
	}
	if l, ok := logger.(*glog.Logger); ok {
		l.Output(2, err.Error())
	} else {
		logger.Println(err)
	}
	return err
}
