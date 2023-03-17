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
	if err := fsutil.Remove("test.ulib.io"); err != nil {
		panic(err)
	}
}

func Benchmark_UlibWriter_Write(b *testing.B) {
	b.ResetTimer()
	w, _ := fsutil.NewWriter("test.ulibw.io")
	for i := 0; i < b.N; i++ {
		w.Add("test.ulibw.io")
	}
	w.Close()
	if err := fsutil.Remove("test.ulibw.io"); err != nil {
		panic(err)
	}
}

func Benchmark_UlibWriter_Strings_Write(b *testing.B) {
	b.ResetTimer()
	w, _ := fsutil.NewWriter("test.ulibws.io")
	for i := 0; i < b.N; i++ {
		w.AddString("test.ulibws.io")
	}
	w.Close()
	if err := fsutil.Remove("test.ulibws.io"); err != nil {
		panic(err)
	}
}

func Benchmark_Basic_Write(b *testing.B) {
	b.ResetTimer()
	f, _ := fsutil.Open("test.basic.io")
	for i := 0; i < b.N; i++ {
		f.Write([]byte("What fuck????"))
	}
	f.Close()
	if err := fsutil.Remove("test.basic.io"); err != nil {
		panic(err)
	}
}
