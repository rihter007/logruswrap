package logruswrap_test

import (
	"github.com/rihter007/logruswrap"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNilInput(t *testing.T) {
	l := logruswrap.WrapPrintfLogger(nil)
	require.NotNil(t, l)
	l.Infof("some log that will not appear anywhere")
}

type dummyLogger struct {
	logruswrap.PrintfLogger

	debugf func(format string, args ...interface{})
}

func (l *dummyLogger) Debugf(format string, args ...interface{}) {
	l.debugf(format, args...)
}

func TestWrappedInput(t *testing.T) {
	var debugInvoked bool
	initialLogger := &dummyLogger{
		debugf: func(format string, args ...interface{}) {
			debugInvoked = true
		},
	}

	l := logruswrap.WrapPrintfLogger(initialLogger)
	require.NotNil(t, l)
	l.Debugf("some input")

	require.True(t, debugInvoked)
}
