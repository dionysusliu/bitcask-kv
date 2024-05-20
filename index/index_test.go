package index

import (
	"bitcask-go/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testIndexerPutNilKey(t *testing.T, indexer *Indexer) {
	assertion := assert.New(t)
	bt := NewBTreeIndexer()

	// edge case: put nil key and retrieve
	res := bt.Put(nil, &data.LogRecordPos{
		Fid:    0,
		Offset: 100,
	})
	assertion.Equal(res, true, "Indexer should support nil key")
}

func testIndexerGetNilKey(t *testing.T, indexer *Indexer) {
	assertion := assert.New(t)
	bt := NewBTreeIndexer()

	// edge case: put nil key and retrieve
	bt.Put(nil, &data.LogRecordPos{
		Fid:    0,
		Offset: 100,
	})
	resFromNilKey := bt.Get(nil)
	assertion.Equal(resFromNilKey.Fid, uint32(0), "Retrieval Result doesn't match")
	assertion.Equal(resFromNilKey.Offset, int64(100), "Retrieval Result doesn't match")
}
