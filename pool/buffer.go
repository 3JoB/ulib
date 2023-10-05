package pool

import (
	"bytes"
	"errors"
	"sync"
)

type BufferClose struct {
	bytes.Buffer
}

var (
	bufferPool      = &sync.Pool{}
	bufferClosePool = &sync.Pool{}

	ErrPtr = errors.New("the incoming pointer cannot be nil")
)

func NewBuffer() *bytes.Buffer {
	r := bufferPool.Get()
	if r == nil {
		return &bytes.Buffer{}
	}
	return r.(*bytes.Buffer)
}

func ReleaseBuffer(b *bytes.Buffer) {
	b.Reset()
	bufferPool.Put(b)
}
