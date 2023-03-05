package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"hash/crc32"
	"io"
	"os"

	"github.com/3JoB/unsafeConvert"
	"github.com/spf13/cast"
)

type Hash struct {
	File string
	HmacKey string
	Hmac bool
}

func NewWithPath(path string) *Hash {
	return &Hash{
		File: path,
	}
}

func (h *Hash) HMAC(key string) *Hash {
	h.Hmac = true
	if key == "" {
		key = "ulibHMAC"
	}
	h.HmacKey = key
	return h
}

func (h *Hash) readHash(hs hash.Hash) string {
	f, err := os.Open(h.File)
	if err != nil {
		f.Close()
		return ""
	}
	defer f.Close()
	_, _ = io.Copy(hs, f)
	return hex.EncodeToString(hs.Sum(nil))
}

func (h *Hash) MD5() string {
	if h.Hmac {
		return h.readHash(hmac.New(md5.New, unsafeConvert.BytesReflect(h.HmacKey)))
	}
	return h.readHash(md5.New())
}

func (h *Hash) SHA1() string {
	if h.Hmac {
		return h.readHash(hmac.New(sha1.New, unsafeConvert.BytesReflect(h.HmacKey)))
	}
	return h.readHash(sha1.New())
}

func (h *Hash) SHA256() string {
	if h.Hmac {
		return h.readHash(hmac.New(sha256.New, unsafeConvert.BytesReflect(h.HmacKey)))
	}
	return h.readHash(sha256.New())
}


func (h *Hash) SHA512() string{
	if h.Hmac {
		return h.readHash(hmac.New(sha512.New, unsafeConvert.BytesReflect(h.HmacKey)))
	}
	return h.readHash(sha512.New())
}

func (h *Hash) CRC32() string {
	f, err := os.Open(h.File)
	if err != nil {
		f.Close()
		return ""
	}
	defer f.Close()
	hs := crc32.NewIEEE()
	_, _ = io.Copy(hs, f)
	if h.Hmac{
		return crc32HMAC(hs, h.HmacKey, cast.ToString(hs.Sum32()))
	}
	return cast.ToString(hs.Sum32())
}

func crc32HMAC(hs hash.Hash32, key, rta string) string{
	m := hmac.New(sha512.New, unsafeConvert.BytesReflect(key))
	if _, err := m.Write(unsafeConvert.BytesReflect(rta)); err != nil {
		return ""
	}
	return hex.EncodeToString(m.Sum(nil))
}