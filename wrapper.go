package logruswrap

import (
	"github.com/sirupsen/logrus"
)

type PrintfLogger interface {
	Logf(level logrus.Level, format string, args ...interface{})
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

type emptyPrintfLogger struct {
}

func (l *emptyPrintfLogger) Logf(level logrus.Level, format string, args ...interface{}) {
}

func (l *emptyPrintfLogger) Tracef(format string, args ...interface{}) {
}

func (l *emptyPrintfLogger) Debugf(format string, args ...interface{}) {
}

func (l *emptyPrintfLogger) Infof(format string, args ...interface{}) {
}

func (l *emptyPrintfLogger) Printf(format string, args ...interface{}) {
}

func (l *emptyPrintfLogger) Warnf(format string, args ...interface{}) {
}

func (l *emptyPrintfLogger) Warningf(format string, args ...interface{}) {
}

func (l *emptyPrintfLogger) Errorf(format string, args ...interface{}) {
}

func (l *emptyPrintfLogger) Fatalf(format string, args ...interface{}) {
}

func (l *emptyPrintfLogger) Panicf(format string, args ...interface{}) {
}

// Wraps PrintfLogger object, if the object is nil, a wrapper that does nothing is returned
func WrapPrintfLogger(l PrintfLogger) PrintfLogger {
	if l == nil {
		return &emptyPrintfLogger{}
	}
	return l
}
