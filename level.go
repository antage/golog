package golog

type Severity int

const (
	ERROR Severity = 1 << iota
	WARNING
	INFO
	DEBUG
	ALL Severity = ^0
)
