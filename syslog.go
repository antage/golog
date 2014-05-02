package golog

type Facility int

const (
	// Facility.

	// From /usr/include/sys/syslog.h.
	// These are the same up to LOG_FTP on Linux, BSD, and OS X.
	SYSLOG_KERN Facility = iota << 3
	SYSLOG_USER
	SYSLOG_MAIL
	SYSLOG_DAEMON
	SYSLOG_AUTH
	SYSLOG_SYSLOG
	SYSLOG_LPR
	SYSLOG_NEWS
	SYSLOG_UUCP
	SYSLOG_CRON
	SYSLOG_AUTHPRIV
	SYSLOG_FTP
	_ // unused
	_ // unused
	_ // unused
	_ // unused
	SYSLOG_LOCAL0
	SYSLOG_LOCAL1
	SYSLOG_LOCAL2
	SYSLOG_LOCAL3
	SYSLOG_LOCAL4
	SYSLOG_LOCAL5
	SYSLOG_LOCAL6
	SYSLOG_LOCAL7
)
