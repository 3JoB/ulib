package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/3JoB/ulib/crypt"
	"github.com/3JoB/unsafeConvert"
)

func c(h func() hash.Hash, data, key string) string {
	key = func() string{
		if key == "" {
			key = "ulib-hmac"
		}
		return key
	}()
	return crypt.Crypt(hmac.New(h, unsafeConvert.BytesReflect(key)), data)
}

func SHA256(data, key string) string {
	return c(sha256.New, key, data)
}

func SHA512(data, key string) string {
	return c(sha512.New, key, data)
}

func MD5(data, key string) string {
	return c(md5.New, key, data)
}
