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
		err = n.addString(s)
	case []byte:
		err = n.addBytes(s)
	default:
		err = n.addAny(w)
	}
	return
}

// Write data of type `[]byte` to the buffer
func (n *nn) AddBytes(w []byte) error {
	return n.addBytes(w)
}

// Write data of type `String` to the buffer
func (n *nn) AddString(w string) error {
	return n.addString(w)
}

func (n *nn) addAny(w any) error {
	n.writer.Write(unsafeConvert.BytesReflect(w.(string)))
	return nil
}

func (n *nn) addString(w string) error {
	n.writer.Write(unsafeConvert.BytesReflect(w))
	return nil
}

func (n *nn) addBytes(w []byte) error {
	n.writer.Write(w)
	return nil
}

// Write the data in the buffer to the file and close the IO channel.
//
// Tips: After this operation, please do not continue to operate on the previous pointer!
func (n *nn) Close() error {
	if err := n.writer.Flush(); err != nil {
		return err
	}
	if err := n.os.Close(); err != nil {
		return err
	}
	n.os = nil
	return nil
}
