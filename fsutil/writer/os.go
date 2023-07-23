package writer

import (
	"os"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/fsutil"
)

type oo struct {
	os *os.File
}

// Higher performance sustainable file write operations.
//
// Example Files: `writer_test.go`
func NewOSWriter(path string) (*oo, error) {
	n := &oo{}
	var err error
	n.os, err = fsutil.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return n, nil
}

// Write data of type `any` to the buffer (automatically checked)
func (n *oo) Add(w any) (err error) {
	switch s := w.(type) {
	case string:
		err = n.AddString(s)
	case []byte:
		err = n.AddBytes(s)
	default:
		_, err = n.os.Write(unsafeConvert.BytePointer(w.(string)))
	}
	return
}

// Write data of type `[]byte` to the buffer
func (n *oo) AddBytes(w []byte) error {
	_, err := n.os.Write(w)
	return err
}

// Write data of type `String` to the buffer
func (n *oo) AddString(w string) error {
	_, err := n.os.Write(unsafeConvert.BytePointer(w))
	return err
}

// Write the data in the buffer to the file and close the IO channel.
//
// Tips: After this operation, please do not continue to operate on the previous pointer!
func (n *oo) Close() error {
	if err := n.os.Close(); err != nil {
		return err
	}
	return nil
}
