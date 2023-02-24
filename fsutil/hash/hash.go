package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
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

func (h *Hash) MD5() (string, error) {
	f, err := os.Open(h.File)
	if err != nil {
		f.Close()
		return "", err
	}
	defer f.Close()
	hs := md5.New()
	_, _ = io.Copy(hs, f)
	return hex.EncodeToString(hs.Sum(nil)), nil
}

func (h *Hash) SHA1() (string, error) {
	f, err := os.Open(h.File)
	if err != nil {
		f.Close()
		return "", err
	}
	defer f.Close()
	hs := sha1.New()
	_, _ = io.Copy(hs, f)
	return hex.EncodeToString(hs.Sum(nil)), nil
}

func (h *Hash) SHA256() (string, error) {
	f, err := os.Open(h.File)
	if err != nil {
		f.Close()
		return "", err
	}
	defer f.Close()
	hs := sha256.New()
	_, _ = io.Copy(hs, f)
	return hex.EncodeToString(hs.Sum(nil)), nil
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
