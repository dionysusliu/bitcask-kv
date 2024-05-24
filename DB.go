package bitcask

import (
	"bitcask-go/data"
	"sync"
	"sync/atomic"
)

type DB struct {
	mu         *sync.RWMutex
	activeFile *data.DBFile            // file for current WAL, read and write
	oldFiles   map[uint32]*data.DBFile // old WAL files, read only

	nextId atomic.Int32 // next file's name
}

func (db *DB) Put(key, value []byte) error {
	// validate key
	if len(key) == 0 {
		return ErrKeyNotFound
	}

	newLogRecord := &data.LogRecordData{
		Key:   key,
		Value: value,
		Type:  data.Normal,
	}

}

func (db *DB) appendLogRecord(logRecord *data.LogRecordData) error {
	// create active file if not exist
	if db.activeFile == nil {
		// create new file
	}
	// write log record to active file
}

func (db *DB) setActiveFile() *data.DBFile {

}
