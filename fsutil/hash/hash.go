package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"hash/crc32"
	"io"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/fsutil"
)

type Crypt string

const (
	MD5 Crypt = "MD5"
	SHA1 Crypt = "SHA1"
	SHA224 Crypt = "SHA224"
	SHA256 Crypt = "SHA256"
	SHA384 Crypt = "SHA384"
	SHA512_224 Crypt = "SHA512_224"
	SHA512_256 Crypt = "SHA512_256"
	SHA512 Crypt = "SHA512"
	CRC32 Crypt = "CRC32"
)

type HashOpt struct{
	HMAC *HashHMAC
	Crypt Crypt
}

type HashHMAC struct {
	Key string
}

func New(path string, opt *HashOpt) string {
	var h func() hash.Hash
	var hs hash.Hash
	switch opt.Crypt {
	case MD5:
		h = md5.New
	case SHA1:
		h = sha1.New
	case SHA224:
		h = sha256.New224
	case SHA256:
		h = sha256.New
	case SHA384:
		h = sha512.New384
	case SHA512_224:
		h = sha512.New512_224
	case SHA512_256:
		h = sha512.New512_256
	case SHA512:
		h = sha512.New
	case CRC32:
		return c32(path, opt)
	default:
		return ""
	}
if opt.HMAC != nil {
			if opt.HMAC.Key == ""{
				opt.HMAC.Key = "ulib"
			}
			hs = hmac.New(h, unsafeConvert.BytesReflect(opt.HMAC.Key))
		} else {
			hs = h()
		}
	f, err := fsutil.Open(path)
	if err != nil {
		f.Close()
		return ""
	}
	defer f.Close()
	_, _ = io.Copy(hs, f)
	return hex.EncodeToString(hs.Sum(nil))
}

func c32(path string, opt *HashOpt) string {
	f, err := fsutil.Open(path)
	if err != nil {
		f.Close()
		return ""
	}
	defer f.Close()
	hs := crc32.NewIEEE()
	_, _ = io.Copy(hs, f)
	if opt.HMAC != nil {
		if opt.HMAC.Key == ""{
			opt.HMAC.Key = "ulib"
		}
		return crc32HMAC(hs, opt.HMAC.Key, fmt.Sprint(hs.Sum32()))
	}
	return fmt.Sprint(hs.Sum32())
}

func crc32HMAC(hs hash.Hash32, key, rta string) string {
	m := hmac.New(sha512.New, unsafeConvert.BytesReflect(key))
	if _, err := m.Write(unsafeConvert.BytesReflect(rta)); err != nil {
		return ""
	}
	return hex.EncodeToString(m.Sum(nil))
}
