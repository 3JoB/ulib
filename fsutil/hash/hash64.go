package hash

import (
	"crypto/hmac"
	"crypto/sha512"
	"hash"
	"hash/crc64"
	"hash/fnv"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/hex"
)

func New64(path string, opt *Opt) string {
	var h hash.Hash64
	switch opt.Crypt {
	case CRC64:
		h = crc64.New(crc64.MakeTable(crc64.ISO))
	case CRC64ECMA:
		h = crc64.New(crc64.MakeTable(crc64.ECMA))
	case Fnv64:
		h = fnv.New64()
	case Fnv64a:
		h = fnv.New64a()
	default:
		return ""
	}
	if opt.HMACKey != "" {
		return hmac64(h, opt.HMACKey)
	}
	return hashel(path, h)
}

func hmac64(hs hash.Hash64, key string) string {
	m := hmac.New(sha512.New, unsafeConvert.BytePointer(key))
	if _, err := m.Write(unsafeConvert.BytePointer(hex.EncodeToString(hs.Sum(nil)))); err != nil {
		return ""
	}
	return hex.EncodeToString(m.Sum(nil))
}
