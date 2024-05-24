package fio

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestNewFileIO(t *testing.T) {
	// create a new file io manager, by normally it should not return error
	filePath := filepath.Join("/tmp", "b.data")
	fio, err := NewFileIOManger(filePath)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	// clear and close the file
	if err := fio.Clear(); err != nil {
		t.Error(err)
	}
	if err := fio.Close(); err != nil {
		t.Error(err)
	}
}

func TestFileIOReadWrite(t *testing.T) {
	// create a new file io manager
	// write a few bytes, then read and check if it's the same
	fio, err := NewFileIOManger(filepath.Join("/tmp", "b.data"))
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	// write bytes
	n, err := fio.Write([]byte(""))
	assert.Equal(t, 0, n)
	assert.Nil(t, err)

	n, err = fio.Write([]byte("hello,world!\n"))
	assert.Equal(t, 13, n)
	assert.Nil(t, err)

	n, err = fio.Write([]byte("this is bitcask-kv\n"))
	assert.Equal(t, 19, n)
	assert.Nil(t, err)

	// read bytes
	rBuf := make([]byte, 1024)
	n, err = fio.Read(rBuf, 0)
	if err.Error() != "EOF" {
		assert.Nil(t, err)
	}
	rString := string(rBuf)
	assert.Equal(t, "hello,world!\nthis is bitcask-kv\n", rString[:32])

	// clear and close the file
	if err := fio.Clear(); err != nil {
		t.Error(err)
	}
	if err := fio.Close(); err != nil {
		t.Error(err)
	}
}

func TestFileIO_AppendLogRecord(t *testing.T) {
	// create a new file io manager
	// write a few bytes, then read and check if it's the same
	fio, err := NewFileIOManger(filepath.Join("/tmp", "b.data"))
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	// append log record
	n, err := fio.AppendLogRecord([]byte("hello,world!\n"))
	assert.Equal(t, 13, n)
	assert.Nil(t, err)

	n, err = fio.AppendLogRecord([]byte("this is bitcask-kv\n"))
	assert.Equal(t, 19, n)
	assert.Nil(t, err)

	// read bytes
	rBuf := make([]byte, 1024)
	n, err = fio.Read(rBuf, 0)
	if err.Error() != "EOF" {
		assert.Nil(t, err)
	}
	rString := string(rBuf)
	assert.Equal(t, "hello,world!\nthis is bitcask-kv\n", rString[:32])

	// clear and close the file
	if err := fio.Clear(); err != nil {
		t.Error(err)
	}
	if err := fio.Close(); err != nil {
		t.Error(err)
	}
}
