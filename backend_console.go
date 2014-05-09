package golog

import (
	"os"
	"sync"
)

type consoleBackend struct {
	sync.Mutex
	stderr *os.File
	stdout *os.File
}

func NewConsoleBackend() Backend {
	backend := &consoleBackend{
		stderr: os.Stderr,
		stdout: os.Stdout,
	}

	return backend
}

func (backend consoleBackend) Error(msg string) {
	backend.Lock()
	defer backend.Unlock()

	backend.stderr.Write([]byte(msg))
	backend.stderr.Sync()
}

func (backend consoleBackend) Warning(msg string) {
	backend.Lock()
	defer backend.Unlock()

	backend.stderr.Write([]byte(msg))
	backend.stderr.Sync()
}

func (backend consoleBackend) Info(msg string) {
	backend.Lock()
	defer backend.Unlock()

	backend.stdout.Write([]byte(msg))
	backend.stdout.Sync()
}

func (backend consoleBackend) Debug(msg string) {
	backend.Lock()
	defer backend.Unlock()

	backend.stdout.Write([]byte(msg))
	backend.stdout.Sync()
}
