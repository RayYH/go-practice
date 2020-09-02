package interfaces

import "bytes"

type ReadWrite interface {
	Read(b bytes.Buffer) bool
	Write(b bytes.Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

// nested interface
type File interface {
	ReadWrite
	Lock
	Close()
}
