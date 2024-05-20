package data

// logRecordPos is storage representation of each in-disk log
type LogRecordPos struct {
	Fid    uint32 // unique identifier of file
	Offset int64  // start position of this log
}
