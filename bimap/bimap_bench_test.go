package bimap

import "testing"

func BenchmarkInsert(b *testing.B) {
	bm := New[int, string]()
	for i := 0; i < b.N; i++ {
		bm.Insert(i, "value")
	}
}

func BenchmarkGet(b *testing.B) {
	bm := New[int, string]()
	bm.Insert(1, "a")
	for i := 0; i < b.N; i++ {
		bm.Get(1)
	}
}

func BenchmarkDelete(b *testing.B) {
	bm := New[int, string]()
	bm.Insert(1, "a")
	for i := 0; i < b.N; i++ {
		bm.Delete(1)
		bm.Insert(1, "a")
	}
}
