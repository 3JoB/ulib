package litefmt_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/3JoB/ulib/litefmt"
)

var dc = []string{
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

var ll = []string{
	"aavewtrhjvtrvtrs",
	"aghasbvaewibavcwe",
}

func Benchmark_String_For(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r := ""
		for _, e := range dc {
			r = r + e
		}
	}
}

func Benchmark_Strings_Join(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = strings.Join(dc, " ")
	}
}

func Benchmark_LiteFMT_Sprint(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = litefmt.Sprint(dc...)
	}
}

func Benchmark_LiteFMT_SprintP(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = litefmt.SprintP(dc...)
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

func Benchmark_L_LITEFMT_Sprint(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = litefmt.Sprint(ll...)
	}
}

func Benchmark_L_FMT_Sprint(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(ll)
	}
}
