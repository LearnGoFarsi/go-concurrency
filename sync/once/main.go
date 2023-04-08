package main

import (
	"os"
	"sync"
)

type FileLogger struct {
	file *os.File
	once sync.Once
}

func NewFileLogger(path string) FileLogger {
	dir, _ := os.MkdirTemp(path, "log")
	file, _ := os.CreateTemp(dir, "log")

	return FileLogger{
		file: file,
		once: sync.Once{},
	}
}

func (f FileLogger) Close() {

	var err error
	f.once.Do(func() {
		err = f.file.Close()
	})

	if err != nil {
		panic(err)
	}
}

func main() {
	logger := NewFileLogger(".")
	logger.Close()

	logger.Close() // this should not cause any panic
}
