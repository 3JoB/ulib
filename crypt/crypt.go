package crypt

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

func SHA256(data string) string {
	if b, err := Crypt(sha256.New(), data); err != nil {
		return ""
	} else {
		return b
	}
}

func SHA512(data string) string {
	if b, err := Crypt(sha512.New(), data); err != nil {
		return ""
	} else {
		return b
	}
}

func MD5(data string) string {
	if b, err := Crypt(md5.New(), data); err != nil {
		return ""
	} else {
		return b
	}
}

func Crypt(h hash.Hash, d string) (string, error) {
	if _, err := h.Write([]byte(d)); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}