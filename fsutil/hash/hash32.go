package hash

import (
	"crypto/hmac"
	"crypto/sha512"
	"hash"
	"hash/crc32"
	"hash/fnv"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/hex"
)

func New32(path string, opt *HashOpt) string {
	var h hash.Hash32
	switch opt.Crypt {
	case CRC32:
		h = crc32.New(crc32.MakeTable(crc32.IEEE))
	case CRC32Castagnoli:
		h = crc32.New(crc32.MakeTable(crc32.Castagnoli))
	case CRC32Koopman:
		h = crc32.New(crc32.MakeTable(crc32.Koopman))
	case Fnv32:
		h = fnv.New32()
	case Fnv32a:
		h = fnv.New32a()
	default:
		return ""
	}
	if opt.HMACKey != "" {
		return hmac32(h, opt.HMACKey)
	}
	return hash32el(path, h)
}

func hmac32(hs hash.Hash32, key string) string {
	m := hmac.New(sha512.New, unsafeConvert.BytesReflect(key))
	if _, err := m.Write(unsafeConvert.BytesReflect(hex.EncodeToString(hs.Sum(nil)))); err != nil {
		return ""
	}
	return hex.EncodeToString(m.Sum(nil))
}
