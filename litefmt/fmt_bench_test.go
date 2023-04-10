package litefmt_test

import (
	"fmt"
	"testing"

	"github.com/3JoB/ulib/litefmt"
)

var dc []string = []string{
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
	"aaaaa",
	"bbbbb",
	"ccccc",
}

func Benchmark_LiteFMT_Sprint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = litefmt.Sprint(dc...)
	}
}

func Benchmark_FMT_Sprint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(dc)
	}
}

func Benchmark_LiteFMT_Sprintln(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = litefmt.Sprintln(dc...)
	}
}

func Benchmark_FMT_Sprintln(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintln(dc)
	}
}