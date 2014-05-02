package golog

type Severity int

const (
	ERROR Severity = 1 << iota
	INFO
	DEBUG
	ALL Severity = ^0
)
