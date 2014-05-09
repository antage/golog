package golog

import (
	"fmt"
)

type Formatter interface {
	FormatError(msg string) string
	FormatWarning(msg string) string
	FormatInfo(msg string) string
	FormatDebug(msg string) string
}

type ChainedFormatter struct {
	formatters []Formatter
}

func NewChainedFormatter(formatters ...Formatter) Formatter {
	return &ChainedFormatter{
		formatters: formatters,
	}
}

func (formatter ChainedFormatter) FormatError(msg string) string {
	result := msg
	for _, f := range formatter.formatters {
		result = f.FormatError(result)
	}
	return result
}

func (formatter ChainedFormatter) FormatWarning(msg string) string {
	result := msg
	for _, f := range formatter.formatters {
		result = f.FormatWarning(result)
	}
	return result
}

func (formatter ChainedFormatter) FormatInfo(msg string) string {
	result := msg
	for _, f := range formatter.formatters {
		result = f.FormatInfo(result)
	}
	return result
}

func (formatter ChainedFormatter) FormatDebug(msg string) string {
	result := msg
	for _, f := range formatter.formatters {
		result = f.FormatDebug(result)
	}
	return result
}

type NoopFormatter struct{}

func (_ NoopFormatter) FormatError(msg string) string {
	return msg
}

func (_ NoopFormatter) FormatWarning(msg string) string {
	return msg
}

func (_ NoopFormatter) FormatInfo(msg string) string {
	return msg
}

func (_ NoopFormatter) FormatDebug(msg string) string {
	return msg
}

type NewlineFormatter struct{}

func (formatter NewlineFormatter) FormatError(msg string) string {
	return fmt.Sprintf("%s\n", msg)
}

func (formatter NewlineFormatter) FormatWarning(msg string) string {
	return fmt.Sprintf("%s\n", msg)
}

func (formatter NewlineFormatter) FormatInfo(msg string) string {
	return fmt.Sprintf("%s\n", msg)
}

func (formatter NewlineFormatter) FormatDebug(msg string) string {
	return fmt.Sprintf("%s\n", msg)
}

type ColorFormatter struct{}

func (formatter ColorFormatter) FormatError(msg string) string {
	return fmt.Sprintf("%s%s%s", "\x1b[31m", msg, "\x1b[0m")
}

func (formatter ColorFormatter) FormatWarning(msg string) string {
	return fmt.Sprintf("%s%s%s", "\x1b[35m", msg, "\x1b[0m")
}

func (formatter ColorFormatter) FormatInfo(msg string) string {
	return fmt.Sprintf("%s%s%s", "\x1b[32m", msg, "\x1b[0m")
}

func (formatter ColorFormatter) FormatDebug(msg string) string {
	return msg
}

func DefaultConsoleFormatter() Formatter {
	return NewChainedFormatter(
		ColorFormatter{},
		NewlineFormatter{})
}
