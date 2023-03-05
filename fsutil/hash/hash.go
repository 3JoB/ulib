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

func readHash(v string, h hash.Hash) (string, error) {
	f, err := os.Open(v)
	if err != nil {
		f.Close()
		return "", err
	}
	defer f.Close()
	_, _ = io.Copy(h, f)
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (h *Hash) MD5() (string, error) {
	if h.Hmac {
		return readHash(h.File, hmac.New(md5.New, unsafeConvert.BytesReflect(h.HmacKey)))
	}
	return readHash(h.File, md5.New())
}

func (h *Hash) SHA1() (string, error) {
	if h.Hmac {
		return readHash(h.File, hmac.New(sha1.New, unsafeConvert.BytesReflect(h.HmacKey)))
	}
	return readHash(h.File, sha1.New())
}

func (h *Hash) SHA256() (string, error) {
	if h.Hmac {
		return readHash(h.File, hmac.New(sha256.New, unsafeConvert.BytesReflect(h.HmacKey)))
	}
	return readHash(h.File, sha256.New())
}


func (h *Hash) SHA512() (string, error) {
	if h.Hmac {
		return readHash(h.File, hmac.New(sha512.New, unsafeConvert.BytesReflect(h.HmacKey)))
	}
	return readHash(h.File, sha512.New())
}

func (h *Hash) CRC32() (string, error) {
	f, err := os.Open(h.File)
	if err != nil {
		f.Close()
		return "", err
	}
	defer f.Close()
	hs := crc32.NewIEEE()
	_, _ = io.Copy(hs, f)
	if h.Hmac{
		return crc32HMAC(hs, h.HmacKey, cast.ToString(hs.Sum32()))
	}
	return cast.ToString(hs.Sum32()), nil
}

func crc32HMAC(hs hash.Hash32, key, rta string) (string, error){
	m := hmac.New(sha512.New, unsafeConvert.BytesReflect(key))
	if _, err := m.Write(unsafeConvert.BytesReflect(rta)); err != nil {
		return "", err
	}
	return hex.EncodeToString(m.Sum(nil)), nil
}