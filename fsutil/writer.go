package fsutil

import (
	"bufio"
	"os"

	"github.com/3JoB/unsafeConvert"
)

type nn struct {
	os  *os.File
	writer *bufio.Writer
}

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