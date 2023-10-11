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
	bufferPool      = &sync.Pool{
		New: func() any {
			return &bytes.Buffer{}
		},
	}
	bufferClosePool = &sync.Pool{
		New: func() any {
			return &BufferClose{}
		},
	}

	ErrPtr = errors.New("the incoming pointer cannot be nil")
)

func NewBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func ReleaseBuffer(b *bytes.Buffer) {
	if b != nil {
		b.Reset()
		bufferPool.Put(b)
	}
}
