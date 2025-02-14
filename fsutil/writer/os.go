package writer

import (
	"os"

	"github.com/3JoB/unsafeConvert"
)

// Os represents a wrapper around an *os.File object for performing write operations.
type Os struct {
	os *os.File
}

// NewOSWriter creates and returns a new Os instance for writing to the specified file path with necessary permissions.
func NewOSWriter(path string) (*Os, error) {
	n := &Os{}
	var err error
	n.os, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return n, nil
}

// Add writes the provided value to the underlying os. Supports string, byte slice, or defaults to string conversion.
func (n *Os) Add(w any) (err error) {
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

// AddBytes writes the provided byte slice to the underlying file and returns an error if the write operation fails.
func (n *Os) AddBytes(w []byte) error {
	_, err := n.os.Write(w)
	return err
}

// AddString writes the given string to the underlying file and returns an error if the write operation fails.
func (n *Os) AddString(w string) error {
	_, err := n.os.Write(unsafeConvert.BytePointer(w))
	return err
}

// Close closes the underlying os.File associated with the Os instance and returns any error encountered.
func (n *Os) Close() error {
	if err := n.os.Close(); err != nil {
		return err
	}
	return nil
}
