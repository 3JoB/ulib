package hash

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"hash/crc32"
	"hash/fnv"
	"io"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/fsutil"
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
	if opt.HMAC != nil {
		if opt.HMAC.Key == "" {
			opt.HMAC.Key = "ulib"
		}
		return hmac32(h, opt.HMAC.Key, fmt.Sprint(h.Sum32()))
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

func hmac32(hs hash.Hash32, key, rta string) string {
	m := hmac.New(sha512.New, unsafeConvert.BytesReflect(key))
	if _, err := m.Write(unsafeConvert.BytesReflect(rta)); err != nil {
		return ""
	}
	return hex.EncodeToString(m.Sum(nil))
}
