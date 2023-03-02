package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"

	"github.com/3JoB/ulib/crypt"
)

func SHA256(data, key string) string {
	if b, err := crypt.Crypt(hmac.New(sha256.New, []byte(key)), data); err != nil {
		return ""
	} else {
		return b
	}
}

func SHA512(data, key string) string {
	if b, err := crypt.Crypt(hmac.New(sha512.New, []byte(key)), data); err != nil {
		return ""
	} else {
		return b
	}
}

func MD5(data, key string) string {
	if b, err := crypt.Crypt(hmac.New(md5.New, []byte(key)), data); err != nil {
		return ""
	} else {
		return b
	}
}
