package index

import (
	"testing"
)

func TestBTreePutNilKey(t *testing.T) {
	var bt Indexer = NewBTreeIndexer()
	t.Run("TestBTree_PutNilKey", func(t *testing.T) {
		testIndexerPutNilKey(t, &bt)
	})
}

func TestBTreeGetNilKey(t *testing.T) {
	var bt Indexer = NewBTreeIndexer()
	t.Run("TestBTree_GetNilKey", func(t *testing.T) {
		testIndexerGetNilKey(t, &bt)
	})
}
