package writer

import (
	"bufio"
	"io"
	"os"

	"github.com/3JoB/unsafeConvert"
)

// Buffer represents a type used for buffered writing with configurable buffer limits and support for various data types.
type Buffer struct {
	buffer    int
	maxbuffer int
	os        *os.File
	writer    *bufio.Writer
}

// NewWriter initializes and returns a new instance of Buffer with a buffered writer, writing to the specified file path.
// If the file cannot be opened or created, it returns an error.
func NewWriter(path string) (*Buffer, error) {
	n := &Buffer{}
	var err error
	n.os, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	n.writer = bufio.NewWriter(n.os)
	return n, nil
}

// CopyTo transfers data from the provided io.Reader to the internal buffered writer. Returns the number of bytes written and an error.
func (n *Buffer) CopyTo(to io.Reader) (written int64, err error) {
	return io.Copy(n.writer, to)
}

// CopyIn flushes the internal buffer and copies data from the internal file to the provided Writer.
func (n *Buffer) CopyIn(in io.Writer) (written int64, err error) {
	if err := n.Flush(); err != nil {
		return 0, err
	}
	return io.Copy(in, n.os)
}

// Add writes the provided data to the internal writer. It supports string, []byte, or other types convertible to string.
func (n *Buffer) Add(w any) (err error) {
	switch s := w.(type) {
	case string:
		err = n.AddString(s)
	case []byte:
		err = n.AddBytes(s)
	default:
		_, err = n.writer.Write(unsafeConvert.BytePointer(w.(string)))
	}
	if err == nil {
		if n.maxbuffer == 0 {
			return
		}
		if n.Buffered() > n.maxbuffer {
			err = n.Flush()
		}
	}
	return
}

// MaxBuffer sets the maximum buffer size for the internal buffered writer.
func (n *Buffer) MaxBuffer(max int) {
	n.maxbuffer = max
}

// Buffered returns the number of bytes currently stored in the internal buffer of the writer.
func (n *Buffer) Buffered() int {
	n.buffer = n.writer.Buffered()
	return n.buffer
}

// AddBytes writes the provided byte slice to the internal buffered writer and returns an error if the write fails.
func (n *Buffer) AddBytes(w []byte) error {
	_, err := n.writer.Write(w)
	return err
}

// AddString writes the provided string to the internal buffered writer and returns an error if the write operation fails.
func (n *Buffer) AddString(w string) error {
	_, err := n.writer.Write(unsafeConvert.BytePointer(w))
	return err
}

// Flush writes any buffered data to the underlying writer and clears the buffer. Returns an error if the flush fails.
func (n *Buffer) Flush() error {
	if n.Buffered() == 0 {
		return nil
	}
	return n.writer.Flush()
}

// Close finalizes the buffered writing operations by flushing the buffer and closing the associated file. Returns an error if any fail.
func (n *Buffer) Close() error {
	if err := n.Flush(); err != nil {
		return err
	}
	if err := n.os.Close(); err != nil {
		return err
	}
	return nil
}
