package fio

// IOManager manages persistent storage of all K-V pairs
// it should monitor one folder where all Bitcask files stored
// it is not guaranteed to be thread-safe
type IOManager interface {
	// Read bytes of data from an IO object, and put in provided buffer
	Read([]byte, int64) (int, error)

	// Write bytes of data to an IO object
	Write([]byte) (int, error)

	// AppendLogRecord Append a log record to the end of the file
	// return insertion position on success, otherwise return error
	AppendLogRecord([]byte) (int, error)

	// Sync all data in indexer to disk
	Sync() error

	// Close the IO object
	Close() error
}
