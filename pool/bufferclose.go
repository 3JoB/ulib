package pool

import (
	"bytes"
	"sync"
)

type BufferClose struct {
	*bytes.Buffer
}

var bufferClosePool = &sync.Pool{
	New: newBufferClose,
}

func newBufferClose() any {
	return &BufferClose{}
}

func NewBufferClose() *BufferClose {
	return bufferClosePool.Get().(*BufferClose)
}

func (b *BufferClose) Close() error {
	b.Reset()
	bufferClosePool.Put(b)
	return nil
}
