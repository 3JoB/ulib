package pool

func NewBufferClose() *BufferClose {
	r := bufferClosePool.Get()
	if r == nil {
		return &BufferClose{}
	}
	return r.(*BufferClose)
}

func (b *BufferClose) Close() error {
	b.Reset()
	bufferClosePool.Put(b)
	return nil
}
