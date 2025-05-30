package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"hash/fnv"
	"io"

	"github.com/3JoB/unsafeConvert"
)

func NewReader(r io.Reader, opt *Opt) string {
	var hs hash.Hash
	var h func() hash.Hash
	switch opt.Crypt {
	case MD5:
		h = md5.New
	case SHA1:
		h = sha1.New
	case SHA224:
		h = sha256.New224
	case SHA256:
		h = sha256.New
	case SHA384:
		h = sha512.New384
	case SHA512_224:
		h = sha512.New512_224
	case SHA512_256:
		h = sha512.New512_256
	case SHA512:
		h = sha512.New
	case Fnv128:
		h = fnv.New128
	case Fnv128a:
		h = fnv.New128a
	default:
		return ""
	}
	if opt.HMACKey != "" {
		hs = hmac.New(h, unsafeConvert.BytePointer(opt.HMACKey))
	} else {
		hs = h()
	}
	return hashrl(r, hs)
}
