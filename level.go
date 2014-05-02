package golog

type Severity int

const (
	ERROR Severity = iota
	INFO
	DEBUG
	ALL = Severity(-1)
)
