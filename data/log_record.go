package data

// logRecordPos is storage representation of each in-disk log
type LogRecordPos struct {
	Fid    uint32 // unique identifier of file
	Offset int64  // start position of this log
}

type LogRecordData struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}

type LogRecordType byte

const (
	Normal LogRecordType = iota
	Deleted
)
