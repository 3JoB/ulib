package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"strings"

	"github.com/3JoB/unsafeConvert"
	"golang.org/x/crypto/bcrypt"

	"github.com/3JoB/ulib/hex"
)

func SHA256[T string | []byte](data T) string {
	return Select(sha256.New(), data)
}

func SHA224[T string | []byte](data T) string {
	return Select(sha256.New224(), data)
}

func SHA384[T string | []byte](data T) string {
	return Select(sha512.New384(), data)
}

func SHA512_224[T string | []byte](data T) string {
	return Select(sha512.New512_224(), data)
}

func SHA512_256[T string | []byte](data T) string {
	return Select(sha512.New512_256(), data)
}

func SHA512[T string | []byte](data T) string {
	return Select(sha512.New(), data)
}

func MD5[T string | []byte](data T) string {
	return Select(md5.New(), data)
}

func HashBcrypt(password string) []byte {
	hash, _ := bcrypt.GenerateFromPassword(unsafeConvert.BytesReflect(password), bcrypt.DefaultCost)
	return hash
}

func CorrectBcrypt(hash []byte, password string) bool {
	return bcrypt.CompareHashAndPassword(hash, unsafeConvert.BytesReflect(password)) == nil
}

func MD5Str[T string | []byte](data T) string {
	return strings.ToUpper(MD5(data))[8:24]
}

func Select(h hash.Hash, data any) string {
	switch s := data.(type) {
	case string:
		return Crypt(h, unsafeConvert.BytesReflect(s))
	case []byte:
		return Crypt(h, s)
	}
	return ""
}

func Crypt(h hash.Hash, d []byte) string {
	if _, err := h.Write(d); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
