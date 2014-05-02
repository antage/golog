package golog

import (
	"fmt"
	"os"
	"sync"
)

type sink struct {
	mask      Severity
	formatter Formatter
	backend   Backend
}

type Logger interface {
	AddBackend(mask Severity, formatter Formatter, backend Backend)

	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})
}

type logger struct {
	sync.RWMutex
	sinks []sink
}

func NewLogger() Logger {
	return &logger{
		sinks: make([]sink, 0, 2),
	}
}

func (l *logger) AddBackend(mask Severity, formatter Formatter, backend Backend) {
	l.Lock()
	defer l.Unlock()

	newSink := sink{
		mask:      mask,
		formatter: formatter,
		backend:   backend,
	}
	l.sinks = append(l.sinks, newSink)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.RLock()
	defer l.RUnlock()

	msg := fmt.Sprintf(format, args...)
	for _, sink := range l.sinks {
		if sink.mask&ERROR == 0 {
			continue
		}

		msgFormatted := sink.formatter.FormatError(msg)
		sink.backend.Error(msgFormatted)
	}
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.RLock()
	defer l.RUnlock()

	msg := fmt.Sprintf(format, args...)
	for _, sink := range l.sinks {
		if sink.mask&INFO == 0 {
			continue
		}

		msgFormatted := sink.formatter.FormatInfo(msg)
		sink.backend.Info(msgFormatted)
	}
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.RLock()
	defer l.RUnlock()

	msg := fmt.Sprintf(format, args...)
	for _, sink := range l.sinks {
		if sink.mask&DEBUG == 0 {
			continue
		}

		msgFormatted := sink.formatter.FormatDebug(msg)
		sink.backend.Debug(msgFormatted)
	}
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.Errorf(format, args...)
	os.Exit(1)
}
