package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/3JoB/unsafeConvert"

	uch "github.com/3JoB/ulib/crypt/hash"
)

var pubkey string = "ulib-hmac"

func c(h func() hash.Hash, data, key string) string {
	if key == "" {
		key = pubkey
	}
	return uch.Select(hmac.New(h, unsafeConvert.BytesReflect(key)), data)
}

func SHA224(data, key string) string {
	return c(sha256.New224, key, data)
}

func SHA256(data, key string) string {
	return c(sha256.New, key, data)
}

func SHA384(data, key string) string {
	return c(sha512.New384, key, data)
}

func SHA512_224(data, key string) string {
	return c(sha512.New512_224, key, data)
}

func SHA512_256(data, key string) string {
	return c(sha512.New512_256, key, data)
}

func SHA512(data, key string) string {
	return c(sha512.New, key, data)
}

func MD5(data, key string) string {
	return c(md5.New, key, data)
}
