package index

import (
	"bitcask-go/data"
)

// Indexer is interface definition for in-memory index
// it stores the pair <key, log-position-in-file>
// all methods should be thread-safe under read / write by default, except for explicit notice
type Indexer interface {
	// Put add K-V pair in indexer
	// return true on success, otherwise return false
	Put(key []byte, pos *data.LogRecordPos) bool

	// Get get value with provided key
	// return value if found, otherwise return nil
	Get(key []byte) *data.LogRecordPos

	// Del delete an K-V pair in indexer
	// return true if K-V doesn't exist or being successfully deleted, return false otherwise
	Del(key []byte) bool
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}
