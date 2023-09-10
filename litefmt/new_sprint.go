package litefmt

import (
	"bytes"
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
