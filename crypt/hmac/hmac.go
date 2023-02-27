package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func SHA256(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	if _, err := h.Write([]byte(data)); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func SHA512(data, key string) string {
	h := hmac.New(sha512.New, []byte(key))
	if _, err := h.Write([]byte(data)); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func MD5(data, key string) string {
	h := hmac.New(md5.New, []byte(key))
	if _, err := h.Write([]byte(data)); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
