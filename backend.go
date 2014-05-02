package golog

type Backend interface {
	Error(msg string)
	Info(msg string)
	Debug(msg string)
}
