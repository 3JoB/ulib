package json

import (
	"context"

	gjs "github.com/goccy/go-json"

	"github.com/3JoB/ulib/reflect"
)

type marshal struct {
	err  error
	data []byte
}

func Marshal(a any) *marshal {
	m := new(marshal)
	m.data, m.err = gjs.Marshal(a)
	return m
}

func (m *marshal) String() string {
	return reflect.String(m.data)
}

func (m *marshal) Bytes() []byte {
	return m.data
}

func Unmarshal(data []byte, str any) error {
	return gjs.Unmarshal(data, str)
}

func UnmarshalString(data string, str any) error {
	return gjs.Unmarshal(reflect.Bytes(data), str)
}

func UnmarshalContext(ctx context.Context, data []byte, v any) error {
	return gjs.UnmarshalContext(ctx, data, v)
}

func UnmarshalStringContext(ctx context.Context, data string, v any) error {
	return gjs.UnmarshalContext(ctx, reflect.Bytes(data), v)
}
