package index

import (
	"bitcask-go/data"
)

// Indexer is interface definition for in-memory index
// it stores the pair <key, log-position-in-file>
type Indexer interface {
	Put(key []byte, pos *data.LogRecordPos) bool
	Get(key []byte) *data.LogRecordPos
	Del(key []byte) bool
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}
