package fsutil_test

import (
	"testing"

	"github.com/3JoB/ulib/fsutil"
)

func Benchmark_Ulib_Write(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fsutil.Write("test.ulib.io", "What fuck????")
	}
	fsutil.Remove("test.ulib.io")
}

func Benchmark_UlibWriter_Write(b *testing.B) {
	b.ResetTimer()
	w, _ := fsutil.NewWriter("test.ulibw.io")
	for i := 0; i < b.N; i++ {
		w.Add("test.ulibw.io")
	}
	w.Close()
	fsutil.Remove("test.ulibw.io")
}

func Benchmark_Basic_Write(b *testing.B) {
	b.ResetTimer()
	f, _ := fsutil.Open("test.basic.io")
	for i := 0; i < b.N; i++ {
		f.Write([]byte("What fuck????"))
	}
	f.Close()
	fsutil.Remove("test.basic.io")
}