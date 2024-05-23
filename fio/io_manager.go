package fio

// DiskManager manages persistent storage of all K-V pairs
// it should monitor one folder where all Bitcask files stored
// it should be thread-safe
type IOManager interface {
	// Read bytes of data from an IO object, and put in provided buffer
	Read([]byte, int64) (int, error)

	// Write bytes of data to an IO object
	Write([]byte) (int, error)

	// Sync all data in indexer to disk
	Sync() error

	// Close the IO object
	Close() error
}
