package bimap

import "testing"

func BenchmarkBaseInsert(b *testing.B) {
	biMap := NewBaseBiMap[int, string]()
	for i := 0; i < b.N; i++ {
		biMap.Insert(i, "value")
	}
}

func BenchmarkBaseGet(b *testing.B) {
	bm := NewBaseBiMap[int, string]()
	bm.Insert(1, "a")
	for i := 0; i < b.N; i++ {
		bm.Get(1)
	}
}

func BenchmarkBaseDelete(b *testing.B) {
	biMap := NewBaseBiMap[int, string]()
	for i := 0; i < b.N; i++ {
		biMap.Insert(i, "value")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		biMap.Delete(i)
	}
}
