package pool

func NewBufferClose() *BufferClose {
	return bufferClosePool.Get().(*BufferClose)
}

func (b *BufferClose) Close() error {
	b.Reset()
	bufferClosePool.Put(b)
	return nil
}
