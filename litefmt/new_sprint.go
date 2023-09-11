package litefmt

import (
	"bytes"
	"sync"
)

var (
	bytesPool = sync.Pool{
		New: func() any {
			return &bytes.Buffer{}
		},
	}
)

func psp_acquire() *bytes.Buffer {
	p := bytesPool.Get()
	if p == nil {
		return &bytes.Buffer{}
	}
	return p.(*bytes.Buffer)
}

func psp_release(b *bytes.Buffer) {
	b.Reset()
	bytesPool.Put(b)
}
