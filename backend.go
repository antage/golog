package golog

type Backend interface {
	Error(msg string)
	Warning(msg string)
	Info(msg string)
	Debug(msg string)
}
