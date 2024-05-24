package fio

import (
	"log"
	"os"
)

type FileIO struct {
	fd *os.File
}

func NewFileIOManger(path string) (*FileIO, error) {
	fd, err := os.OpenFile(
		path,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, err
	}
	return &FileIO{fd: fd}, nil
}

func (f *FileIO) Read(b []byte, off int64) (int, error) {
	return f.fd.ReadAt(b, off)
}

func (f *FileIO) Write(b []byte) (int, error) {
	return f.fd.Write(b)
}

func (f *FileIO) AppendLogRecord(b []byte) (int, error) {
	_, err := f.fd.Seek(0, 2)
	if err != nil {
		return 0, err
	}
	n, err := f.fd.Write(b)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func (f *FileIO) Sync() error {
	return f.fd.Sync()
}

func (f *FileIO) Close() error {
	return f.fd.Close()
}

func (f *FileIO) Clear() error {
	if _, err := f.fd.Seek(0, 0); err != nil {
		log.Fatalf("Failed to seek file: %v", err)
	}
	return f.fd.Truncate(0)
}
