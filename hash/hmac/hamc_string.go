package hmac

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"

	"github.com/3JoB/unsafeConvert"
	"golang.org/x/crypto/sha3"

	"github.com/3JoB/ulib/internal/hash"
)

func SHA3_224S(data, key string) *hash.Hash {
	return c(sha3.New224, unsafeConvert.BytePointer(key), unsafeConvert.BytePointer(data))
}

func SHA3_256S(data, key string) *hash.Hash {
	return c(sha3.New256, unsafeConvert.BytePointer(key), unsafeConvert.BytePointer(data))
}

func SHA3_384S(data, key string) *hash.Hash {
	return c(sha3.New384, unsafeConvert.BytePointer(key), unsafeConvert.BytePointer(data))
}

func SHA3_512S(data, key string) *hash.Hash {
	return c(sha3.New512, unsafeConvert.BytePointer(key), unsafeConvert.BytePointer(data))
}

func SHA256S(data, key string) *hash.Hash {
	return c(sha256.New, unsafeConvert.BytePointer(key), unsafeConvert.BytePointer(data))
}

func SHA224S(data, key string) *hash.Hash {
	return c(sha256.New224, unsafeConvert.BytePointer(key), unsafeConvert.BytePointer(data))
}

func SHA384S(data, key string) *hash.Hash {
	return c(sha512.New384, unsafeConvert.BytePointer(key), unsafeConvert.BytePointer(data))
}

func SHA512_224S(data, key string) *hash.Hash {
	return c(sha512.New512_224, unsafeConvert.BytePointer(key), unsafeConvert.BytePointer(data))
}

func SHA512_256S(data, key string) *hash.Hash {
	return c(sha512.New512_256, unsafeConvert.BytePointer(key), unsafeConvert.BytePointer(data))
}

func SHA512S(data, key string) *hash.Hash {
	return c(sha512.New, unsafeConvert.BytePointer(key), unsafeConvert.BytePointer(data))
}

func MD5S(data, key string) *hash.Hash {
	return c(md5.New, unsafeConvert.BytePointer(key), unsafeConvert.BytePointer(data))
}

// NewShake128 creates a new SHAKE128 variable-output-length ShakeHash.
// Its generic security strength is 128 bits against all attacks
// if at least 32 bytes of its output are used.
func Shake128S(data string, bits int) string {
	return unsafeConvert.StringSlice(Shake128(unsafeConvert.ByteSlice(data), bits))
}

// NewShake256 creates a new SHAKE256 variable-output-length ShakeHash.
// Its generic security strength is 256 bits against all attacks
// if at least 64 bytes of its output are used.
func Shake256S(data string, bits int) string {
	return unsafeConvert.StringSlice(Shake256(unsafeConvert.ByteSlice(data), bits))
}
