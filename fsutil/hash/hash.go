package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"hash/fnv"

	"github.com/3JoB/unsafeConvert"
)

const (
	MD5             = iota
	SHA1            // New()
	SHA224          // New()
	SHA256          // New()
	SHA384          // New()
	SHA512_224      // New()
	SHA512_256      // New()
	SHA512          // New()
	Fnv128          // New()
	Fnv128a         // New()
	CRC32           // New32()
	CRC32Castagnoli // New32()
	CRC32Koopman    // New32()
	Fnv32           // New32()
	Fnv32a          // New32()
	CRC64           // New64()
	CRC64ECMA       // New64()
	Fnv64           // New64()
	Fnv64a          // New64()
)

type HashOpt struct {
	HMACKey string
	Crypt   int
}

func New(path string, opt *HashOpt) string {
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
		hs = hmac.New(h, unsafeConvert.BytesReflect(opt.HMACKey))
	} else {
		hs = h()
	}
	return hashel(path, hs)
}
