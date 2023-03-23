package crypt

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"strings"

	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/hex"
)

func SHA256(data string) string {
	return Crypt(sha256.New(), data)
}

func SHA512(data string) string {
	return Crypt(sha512.New(), data)
}

func MD5(data string) string {
	return Crypt(md5.New(), data)
}

func MD5Str(s string) string {
	h := md5.New()
	h.Write(unsafeConvert.BytesReflect(s))
	result := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	result = result[8:24]
	return result
}

func Crypt(h hash.Hash, d string) string {
	if _, err := h.Write(unsafeConvert.BytesReflect(d)); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
