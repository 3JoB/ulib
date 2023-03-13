package hash

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"hash/crc64"
	"hash/fnv"
	"io"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/fsutil"
)

func New64(path string, opt *HashOpt) string {
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
	if opt.HMAC != nil {
		if opt.HMAC.Key == "" {
			opt.HMAC.Key = "ulib"
		}
		return hmac64(h, opt.HMAC.Key, fmt.Sprint(h.Sum64()))
	}
	f, err := fsutil.Open(path)
	if err != nil {
		f.Close()
		return ""
	}
	defer f.Close()
	_, _ = io.Copy(h, f)
	return hex.EncodeToString(h.Sum(nil))
}

func hmac64(hs hash.Hash64, key, rta string) string {
	m := hmac.New(sha512.New, unsafeConvert.BytesReflect(key))
	if _, err := m.Write(unsafeConvert.BytesReflect(rta)); err != nil {
		return ""
	}
	return hex.EncodeToString(m.Sum(nil))
}
