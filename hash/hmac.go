package hash

import (
	"crypto/hmac"
	hs "hash"

	"github.com/3JoB/unsafeConvert"
	"golang.org/x/crypto/sha3"

	"github.com/3JoB/ulib/hex"
	"github.com/3JoB/ulib/internal/hash"
)

var (
	pubkey = unsafeConvert.BytePointer("ulib-hmac")
)

func c(h func() hs.Hash, data, key []byte) *hash.Hash {
	if key == nil {
		key = pubkey
	}
	return hash.Crypt(hmac.New(h, key), data)
}

func CreateHMAC(data, key []byte, h func() hs.Hash) *hash.Hash {
	return c(h, key, data)
}

// NewShake128 creates a new SHAKE128 variable-output-length ShakeHash.
// Its generic security strength is 128 bits against all attacks
// if at least 32 bytes of its output are used.
func Shake128(data []byte, bits int) string {
	shake := sha3.NewShake128()
	shake.Write(data)
	if bits > 128 {
		bits = 128
	} else if bits < 32 {
		bits = 32
	}
	h := make([]byte, bits)
	shake.Read(h)
	return hex.EncodeToString(h)
}

// NewShake256 creates a new SHAKE256 variable-output-length ShakeHash.
// Its generic security strength is 256 bits against all attacks
// if at least 64 bytes of its output are used.
func Shake256(data []byte, bits int) string {
	shake := sha3.NewShake256()
	shake.Write(data)
	if bits > 256 {
		bits = 256
	} else if bits < 64 {
		bits = 64
	}
	h := make([]byte, bits)
	shake.Read(h)
	return hex.EncodeToString(h)
}
