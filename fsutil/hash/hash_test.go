package hash_test

import (
	"testing"

	"github.com/3JoB/ulib/fsutil/hash"
)

func Benchmark_Hash_MD5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New("hash.go", &hash.HashOpt{Crypt: hash.MD5})
	}
}

func Benchmark_Hash_MD5_HMAC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New("hash.go", &hash.HashOpt{Crypt: hash.MD5, HMACKey: "12345"})
	}
}

func Benchmark_Hash_SHA1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New("hash.go", &hash.HashOpt{Crypt: hash.SHA1})
	}
}

func Benchmark_Hash_SHA1_HMAC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New("hash.go", &hash.HashOpt{Crypt: hash.SHA1, HMACKey: "12345"})
	}
}

func Benchmark_Hash_SHA224(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New("hash.go", &hash.HashOpt{Crypt: hash.SHA224})
	}
}

func Benchmark_Hash_SHA224_HMAC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New("hash.go", &hash.HashOpt{Crypt: hash.SHA224, HMACKey: "12345"})
	}
}

func Benchmark_Hash_Fnv128(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New("hash.go", &hash.HashOpt{Crypt: hash.Fnv128})
	}
}

func Benchmark_Hash_Fnv128_HMAC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New("hash.go", &hash.HashOpt{Crypt: hash.Fnv128, HMACKey: "12345"})
	}
}

func Benchmark_Hash_Fnv32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New32("hash.go", &hash.HashOpt{Crypt: hash.Fnv32})
	}
}

func Benchmark_Hash_Fnv32_HMAC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New32("hash.go", &hash.HashOpt{Crypt: hash.Fnv32, HMACKey: "12345"})
	}
}

func Benchmark_Hash_Fnv32a(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New32("hash.go", &hash.HashOpt{Crypt: hash.Fnv32a})
	}
}

func Benchmark_Hash_Fnv32a_HMAC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New32("hash.go", &hash.HashOpt{Crypt: hash.Fnv32a, HMACKey: "12345"})
	}
}

func Benchmark_Hash_CRC64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New64("hash.go", &hash.HashOpt{Crypt: hash.CRC64})
	}
}

func Benchmark_Hash_CRC64_HMAC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New64("hash.go", &hash.HashOpt{Crypt: hash.CRC64, HMACKey: "12345"})
	}
}

func Benchmark_Hash_CRC64ECMA(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash.New64("hash.go", &hash.HashOpt{Crypt: hash.CRC64ECMA})
	}
}

func Benchmark_Hash_CRC64ECMA_HMAC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New64("hash.go", &hash.HashOpt{Crypt: hash.CRC64ECMA, HMACKey: "12345"})
	}
}

func Benchmark_Hash_Fnv64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New64("hash.go", &hash.HashOpt{Crypt: hash.Fnv64})
	}
}

func Benchmark_Hash_Fnv64_HMAC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New64("hash.go", &hash.HashOpt{Crypt: hash.Fnv64, HMACKey: "12345"})
	}
}

func Benchmark_Hash_Fnv64a(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New64("hash.go", &hash.HashOpt{Crypt: hash.Fnv64a})
	}
}

func Benchmark_Hash_Fnv64a_HMAC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash.New64("hash.go", &hash.HashOpt{Crypt: hash.Fnv64a, HMACKey: "12345"})
	}
}
