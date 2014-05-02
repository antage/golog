package golog

import (
	"log/syslog"
)

type syslogBackend struct {
	syslog *syslog.Writer
}

func NewSyslogBackend(facility Facility, tag string) (Backend, error) {
	syslogWriter, err := syslog.New(syslog.Priority(int(syslog.LOG_EMERG)|int(facility)), tag)
	if err != nil {
		return nil, err
	}

	return &syslogBackend{
		syslog: syslogWriter,
	}, nil
}

func (backend syslogBackend) Error(msg string) {
	backend.syslog.Err(msg)
}

func (backend syslogBackend) Info(msg string) {
	backend.syslog.Info(msg)
}

func (backend syslogBackend) Debug(msg string) {
	backend.syslog.Debug(msg)
}
