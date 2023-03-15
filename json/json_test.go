package json_test

import (
	"testing"

	ej "encoding/json"
	js "github.com/goccy/go-json"
	"github.com/3JoB/ulib/json"
)

type TestStruct struct {
	A string `json:"a"`
}

// Marshal
func Benchmark_Json_Marshal(b *testing.B) {
	b.ResetTimer()
	var tsc TestStruct
	for i := 0; i < b.N; i++ {
		tsc = TestStruct{A: "test"}
		_, _ = ej.Marshal(&tsc)
	}
}

func Benchmark_GoJson_Marshal(b *testing.B) {
	b.ResetTimer()
	var tsc TestStruct
	for i := 0; i < b.N; i++ {
		tsc = TestStruct{A: "test"}
		_, _ = js.Marshal(&tsc)
	}
}

// Marshal
func Benchmark_Ulib_Marshal(b *testing.B) {
	b.ResetTimer()
	var tsc TestStruct
	for i := 0; i < b.N; i++ {
		tsc = TestStruct{A: "test"}
		_ = json.Marshal(&tsc).String()
	}
}

// Unmarshal
func Benchmark_Json_Unmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := `{"a": "b"}`
		var tsc TestStruct
		if err := ej.Unmarshal([]byte(data), &tsc); err != nil {
			panic(err)
		}
	}
}

func Benchmark_GoJson_Unmarshal(b *testing.B) {
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		data := `{"a": "b"}`
		var tsc TestStruct
		if err := js.Unmarshal([]byte(data), &tsc); err != nil {
			panic(err)
		}
	}
}

// Unmarshal
func Benchmark_Ulib_Unmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := `{"a": "b"}`
		var tsc TestStruct
		if err := json.UnmarshalString(data, &tsc); err != nil {
			panic(err)
		}
	}
}
