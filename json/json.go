package json

import (
	"context"

	"github.com/3JoB/unsafeConvert"
	gjs "github.com/goccy/go-json"
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
	return unsafeConvert.String(m.data)
}

func (m *marshal) Bytes() []byte {
	return m.data
}

func Unmarshal(data []byte, str any) error {
	return gjs.Unmarshal(data, str)
}

func UnmarshalString(data string, str any) error {
	return gjs.Unmarshal(unsafeConvert.Bytes(data), str)
}

func UnmarshalContext(ctx context.Context, data []byte, v any) error {
	return gjs.UnmarshalContext(ctx, data, v)
}

func UnmarshalStringContext(ctx context.Context, data string, v any) error {
	return gjs.UnmarshalContext(ctx, unsafeConvert.Bytes(data), v)
}
