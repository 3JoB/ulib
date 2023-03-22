package fsutil

import (
	"bufio"
	"os"

	"github.com/3JoB/unsafeConvert"
)

type nn struct {
	os     *os.File
	writer *bufio.Writer
}

// Higher performance sustainable file write operations.
//
// Example Files: `writer_test.go`
func NewWriter(path string) (*nn, error) {
	n := &nn{}
	var err error
	n.os, err = OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	n.writer = bufio.NewWriter(n.os)
	return n, nil
}

// Write data of type `any` to the buffer (automatically checked)
func (n *nn) Add(w any) (err error) {
	switch s := w.(type) {
	case string:
		err = n.AddString(s)
	case []byte:
		err = n.AddBytes(s)
	default:
		_, err = n.writer.Write(unsafeConvert.BytesReflect(w.(string)))
	}
	return
}

// Write data of type `[]byte` to the buffer
func (n *nn) AddBytes(w []byte) error {
	_, err := n.writer.Write(w)
	return err
}

// Write data of type `String` to the buffer
func (n *nn) AddString(w string) error {
	_, err := n.writer.Write(unsafeConvert.BytesReflect(w))
	return err
}

// Flush writes any buffered data to the underlying io.Writer.
func (n *nn) Flush() error {
	return n.writer.Flush()
}

// Write the data in the buffer to the file and close the IO channel.
//
// Tips: After this operation, please do not continue to operate on the previous pointer!
func (n *nn) Close() error {
	if err := n.Flush(); err != nil {
		return err
	}
	if err := n.os.Close(); err != nil {
		return err
	}
	n.os = nil
	return nil
}
