package fio

// DiskManager manages persistent storage of all K-V pairs
// it should monitor one folder where all Bitcask files stored
// it should be thread-safe
type DiskManager interface {
	Read(Fid uint32, Offset int64) []byte

	Write(Fid uint32) int64

	Sync(Fid uint32) bool

	Merge() bool
}
