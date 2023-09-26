package hash

import (
	"sync"

	"golang.org/x/crypto/sha3"
)

var (
	shake128pool = &sync.Pool{
		New: func() any {
			return sha3.NewShake128()
		},
	}
	shake256pool = &sync.Pool{
		New: func() any {
			return sha3.NewShake256()
		},
	}
)

func AcquireShake128() sha3.ShakeHash {
	return shake128pool.Get().(sha3.ShakeHash)
}

func AcquireShake256() sha3.ShakeHash {
	return shake256pool.Get().(sha3.ShakeHash)
}

func ReleaseShake128(s sha3.ShakeHash) {
	s.Reset()
	shake128pool.Put(s)
}

func ReleaseShake256(s sha3.ShakeHash) {
	s.Reset()
	shake256pool.Put(s)
}
