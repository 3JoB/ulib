package hash

import (
	"hash"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/hex"
)

type Hash struct {
	Data []byte
}

func Crypt(h hash.Hash, d []byte) *Hash {
	h.Write(d)
	hs := &Hash{
		Data: h.Sum(nil),
	}
	return hs
}

func (h *Hash) Sum() []byte {
	return h.Data
}

func (h *Hash) String() string {
	return unsafeConvert.StringReflect(h.Data)
}

func (h *Hash) Bytes() []byte {
	return h.Data
}

func (h *Hash) Hex() string {
	return hex.EncodeToString(h.Data)
}
