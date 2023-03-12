package hash

/*import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"os"
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

func (h *HashF) hashWithOs(hs hash.Hash) string {
	_, _ = io.Copy(hs, h.Os)
	if h.Close {
		h.Os.Close()
	}
	return hex.EncodeToString(hs.Sum(nil))
}

func (h *HashF) MD5() string {
	return h.hashWithOs(md5.New())
}

func (h *HashF) SHA1() string {
	return h.hashWithOs(sha1.New())
}

func (h *HashF) SHA256() string {
	return h.hashWithOs(sha256.New())
}

func (h *HashF) CRC32() string {
	hs := crc32.NewIEEE()
	_, _ = io.Copy(hs, h.Os)
	if h.Close {
		h.Os.Close()
	}
	return fmt.Sprint(hs.Sum32())
}
*/