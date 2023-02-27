package crypt

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func SHA256(data string) string {
	h := sha256.New()
	if _, err := h.Write([]byte(data)); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func SHA512(data string) string {
	h := sha512.New()
	if _, err := h.Write([]byte(data)); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func MD5(data string) string {
	h := md5.New()
	if _, err := h.Write([]byte(data)); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
