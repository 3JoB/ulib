package hash

import (
	"crypto/md5"
	hash2 "hash"
	"strings"

	"github.com/3JoB/unsafeConvert"
	"golang.org/x/crypto/bcrypt"

	"github.com/3JoB/ulib/internal/hash"
)

func CreateHash(data []byte, h hash2.Hash) *hash.Hash {
	return hash.Crypt(h, data)
}

func Bcrypt(password string) []byte {
	bytes, _ := bcrypt.GenerateFromPassword(unsafeConvert.BytePointer(password), bcrypt.DefaultCost)
	return bytes
}

func CorrectBcrypt(hash []byte, password string) bool {
	return bcrypt.CompareHashAndPassword(hash, unsafeConvert.BytePointer(password)) == nil
}

func MD5(data []byte) string {
	return strings.ToUpper(CreateHash(data, md5.New()).Hex())[8:24]
}
