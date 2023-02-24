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

type Hash struct {
	File string
}

func New(path string) *Hash {
	return &Hash{
		File: path,
	}
}

func ckpt(v string, h hash.Hash) (string, error) {
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
	return ckpt(h.File, md5.New())
}

func (h *Hash) SHA1() (string, error) {
	return ckpt(h.File, sha1.New())
}

func (h *Hash) SHA256() (string, error) {
	return ckpt(h.File, sha256.New())
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
	return cast.ToString(hs.Sum32()), nil
}
