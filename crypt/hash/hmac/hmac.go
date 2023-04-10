package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/3JoB/unsafeConvert"
	"golang.org/x/crypto/sha3"

	uch "github.com/3JoB/ulib/crypt/hash"
)

var pubkey = unsafeConvert.BytesReflect("ulib-hmac")

func c(h func() hash.Hash, data, key []byte) hash.Hash {
	if key == nil {
		key = pubkey
	}
	return uch.Crypt(hmac.New(h, key), data)
}

func SHA3_224(data, key []byte) hash.Hash {
	return c(sha3.New224, key, data)
}

func SHA3_256(data, key []byte) hash.Hash {
	return c(sha3.New256, key, data)
}

func SHA3_384(data, key []byte) hash.Hash {
	return c(sha3.New384, key, data)
}

func SHA3_512(data, key []byte) hash.Hash {
	return c(sha3.New512, key, data)
}

func SHA256(data, key []byte) hash.Hash{
	return c(sha256.New, key, data)
}

func SHA224(data, key []byte) hash.Hash {
	return c(sha256.New224, key, data)
}

func SHA384(data, key []byte) hash.Hash {
	return c(sha512.New384, key, data)
}

func SHA512_224(data, key []byte) hash.Hash {
	return c(sha512.New512_224, key, data)
}

func SHA512_256(data, key []byte) hash.Hash {
	return c(sha512.New512_256, key, data)
}

func SHA512(data, key []byte) hash.Hash {
	return c(sha512.New, key, data)
}

func MD5(data, key []byte) hash.Hash {
	return c(md5.New, key, data)
}

// NewShake128 creates a new SHAKE128 variable-output-length ShakeHash. 
// Its generic security strength is 128 bits against all attacks 
// if at least 32 bytes of its output are used.
func Shake128(data []byte, bits int) []byte {
	shake := sha3.NewShake128()
	shake.Write(data)
	var mdata []byte
	if bits > 128 {
		bits = 128
	} else if bits < 32 {
		bits = 32
	}
	mdata = make([]byte, bits)
	shake.Read(mdata)
	shake.Reset()
	return mdata
}

// NewShake256 creates a new SHAKE256 variable-output-length ShakeHash. 
// Its generic security strength is 256 bits against all attacks 
// if at least 64 bytes of its output are used.
func Shake256(data []byte, bits int) []byte {
	shake := sha3.NewShake256()
	shake.Write(data)
	var mdata []byte
	if bits > 256 {
		bits = 128
	} else if bits < 64 {
		bits = 64
	}
	mdata = make([]byte, bits)
	shake.Read(mdata)
	shake.Reset()
	return mdata
}
