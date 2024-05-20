package index

import (
	"bitcask-go/data"
	"bytes"
	"github.com/google/btree"
	"sync"
)

// BTree is a btree implementation of Our Indexer
// credit to google: "github.com/google/btree"
type BTree struct {
	lock *sync.RWMutex
	tree *btree.BTree
}

// Interface for google-btree's items
func (ai Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}

// add new key-value pair in btree
// return true on success, otherwise return false
// this method is thread-safe
func (bt *BTree) Put(key []byte, pos *data.LogRecordPos) bool {
	item := Item{key, pos}
	bt.lock.Lock()
	bt.tree.ReplaceOrInsert(item)
	bt.lock.Unlock()
	return true
}

// try to get the value from provided key
// if K-V pair exists, return the value, otherwise return nil
// this method is thread-safe
func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	keySearch := Item{key: key}
	bt.lock.Lock()
	item := bt.tree.Get(keySearch)
	bt.lock.Unlock()
	if item == nil {
		return nil
	}
	return item.(*Item).pos
}

// try to delete a key in storage
// if K-V pair exists and deleted successfully, return true, otherwise return false
// this method is thread safe
func (bt *BTree) Del(key []byte) bool {
	keySearch := Item{key: key}
	bt.lock.Lock()
	item := bt.tree.Delete(keySearch)
	bt.lock.Unlock()
	if item == nil {
		return false
	}
	return true
}
