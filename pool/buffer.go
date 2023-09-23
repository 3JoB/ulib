package pool

import (
	"bytes"
	"sync"
)

var (
	bufferPool = &sync.Pool{
		New: newBuffer,
	}
)

func newBuffer() any {
	return &bytes.Buffer{}
}

func NewBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func ReleaseBuffer(b *bytes.Buffer) {
	b.Reset()
	bufferPool.Put(b)
}
