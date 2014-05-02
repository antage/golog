package golog

import (
	"io"
	"os"
	"sync"
)

type writerBackend struct {
	sync.Mutex
	writer io.Writer
	flush  func()
}

func NewWriterBackend(writer io.Writer) Backend {
	backend := &writerBackend{
		writer: writer,
	}

	if file, ok := writer.(*os.File); ok {
		backend.flush = func() {
			file.Sync()
		}
	} else {
		backend.flush = func() {}
	}

	return backend
}

func (backend writerBackend) Error(msg string) {
	backend.Lock()
	defer backend.Unlock()

	backend.writer.Write([]byte(msg))
	backend.flush()
}

func (backend writerBackend) Info(msg string) {
	backend.Lock()
	defer backend.Unlock()

	backend.writer.Write([]byte(msg))
	backend.flush()
}

func (backend writerBackend) Debug(msg string) {
	backend.Lock()
	defer backend.Unlock()

	backend.writer.Write([]byte(msg))
	backend.flush()
}
