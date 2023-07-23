package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"strings"

	"github.com/3JoB/unsafeConvert"
	"golang.org/x/crypto/sha3"

	"github.com/3JoB/ulib/internal/hash"
)

func SHA3_224S(data string) *hash.Hash {
	return hash.Crypt(sha3.New224(), unsafeConvert.BytePointer(data))
}

func SHA3_256S(data string) *hash.Hash {
	return hash.Crypt(sha3.New256(), unsafeConvert.BytePointer(data))
}

func SHA3_384S(data string) *hash.Hash {
	return hash.Crypt(sha3.New384(), unsafeConvert.BytePointer(data))
}

func SHA3_512S(data string) *hash.Hash {
	return hash.Crypt(sha3.New512(), unsafeConvert.BytePointer(data))
}

func SHA224S(data string) *hash.Hash {
	return hash.Crypt(sha256.New224(), unsafeConvert.BytePointer(data))
}

func SHA256S(data string) *hash.Hash {
	return hash.Crypt(sha256.New(), unsafeConvert.BytePointer(data))
}

func SHA384S(data string) *hash.Hash {
	return hash.Crypt(sha512.New384(), unsafeConvert.BytePointer(data))
}

func SHA512S(data string) *hash.Hash {
	return hash.Crypt(sha512.New(), unsafeConvert.BytePointer(data))
}

func SHA512_224S(data string) *hash.Hash {
	return hash.Crypt(sha512.New512_224(), unsafeConvert.BytePointer(data))
}

func SHA512_256S(data string) *hash.Hash {
	return hash.Crypt(sha512.New512_256(), unsafeConvert.BytePointer(data))
}

func MD5S(data string) *hash.Hash {
	return hash.Crypt(md5.New(), unsafeConvert.BytePointer(data))
}

func MD5StrS(data string) string {
	return strings.ToUpper(MD5(unsafeConvert.BytePointer(data)).Hex())[8:24]
}
