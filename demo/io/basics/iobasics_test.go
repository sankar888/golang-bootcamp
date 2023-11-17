// Package iobasics introduces the basic io concepts and abstractions available in golang
// It also provides basic file read, write examples
// golang io is to be leared with os and net package. Detailed learning and examples are provided in separate files
package iobasics

/*
 Q: I/O is basically about read and write to various sources and sinks. What sunc sources are available for I/O ?
 A: Primary sources/sinks involve Filesystem, Network - tcp / udp port, Linux socket.

 Q: What is the basic abstraction for reading and writing from various sources ?
 A: golang io package introduces basic interfaces which defines the basic abstraction of io. The basic Abstractions are

type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer is the interface that wraps the basic Write method. Write writes len(p) bytes from p to the underlying data stream.
type Writer interface {
	Write(p []byte) (n int, err error)
}

// Closer is the interface that wraps the basic Close method. The behavior of Close after the first call is undefined. Specific implementations may document their own behavior.
type Closer interface {
	Close() error
}

// Seeker is the interface that wraps the basic Seek method. Seek sets the offset for the next Read or Write to offset, interpreted according to whence:
// SeekStart means relative to the start of the file,
// SeekCurrent means relative to the current offset, and
// SeekEnd means relative to the end
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

In addition to it, we have ReadAt, WriteAt which reads and writes to a specific offset.
type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}

type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}
*/
