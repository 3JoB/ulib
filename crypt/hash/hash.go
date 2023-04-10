package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"strings"

	"github.com/3JoB/unsafeConvert"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"

	"github.com/3JoB/ulib/hex"
)

func SHA3_224(data []byte) hash.Hash {
	return Crypt(sha3.New224(), data)
}

func SHA3_256(data []byte) hash.Hash {
	return Crypt(sha3.New256(), data)
}

func SHA3_384(data []byte) hash.Hash {
	return Crypt(sha3.New384(), data)
}

func SHA3_512(data []byte) hash.Hash {
	return Crypt(sha3.New512(), data)
}

func SHA224(data []byte) hash.Hash {
	return Crypt(sha256.New224(), data)
}

func SHA256(data []byte) hash.Hash {
	return Crypt(sha256.New(), data)
}

func SHA384(data []byte) hash.Hash {
	return Crypt(sha512.New384(), data)
}

func SHA512(data []byte) hash.Hash {
	return Crypt(sha512.New(), data)
}

func SHA512_224(data []byte) hash.Hash {
	return Crypt(sha512.New512_224(), data)
}

func SHA512_256(data []byte) hash.Hash {
	return Crypt(sha512.New512_256(), data)
}

func MD5(data []byte) hash.Hash {
	return Crypt(md5.New(), data)
}

func HashBcrypt(password string) []byte {
	hash, _ := bcrypt.GenerateFromPassword(unsafeConvert.BytesReflect(password), bcrypt.DefaultCost)
	return hash
}

func CorrectBcrypt(hash []byte, password string) bool {
	return bcrypt.CompareHashAndPassword(hash, unsafeConvert.BytesReflect(password)) == nil
}

func MD5Str(data []byte) string {
	return strings.ToUpper(HexEncoding(MD5(data)))[8:24]
}

func Crypt(h hash.Hash, d []byte) hash.Hash {
	if _, err := h.Write(d); err != nil {
		return nil
	}
	return h
}

func HexEncoding(h hash.Hash) string {
	return hex.EncodeToString(h.Sum(nil))
}