package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"hash/crc32"
	"io"
	"os"

	"github.com/spf13/cast"
)

type HashF struct {
	Os    *os.File
	Close bool
}

func NewWithOs(o *os.File) *HashF {
	return &HashF{Os: o, Close: true}
}

func (h *HashF) DisableAutoClose() *HashF {
	h.Close = false
	return h
}

func hashWithOs(fs *os.File, close bool, h hash.Hash) (string, error) {
	_, _ = io.Copy(h, fs)
	if close {
		fs.Close()
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (h *HashF) MD5() (string, error) {
	return hashWithOs(h.Os, h.Close, md5.New())
}

func (h *HashF) SHA1() (string, error) {
	return hashWithOs(h.Os, h.Close, sha1.New())
}

func (h *HashF) SHA256() (string, error) {
	return hashWithOs(h.Os, h.Close, sha256.New())
}

func (h *HashF) CRC32() (string, error) {
	hs := crc32.NewIEEE()
	_, _ = io.Copy(hs, h.Os)
	if h.Close {
		h.Os.Close()
	}
	return cast.ToString(hs.Sum32()), nil
}
