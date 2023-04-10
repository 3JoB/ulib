package json

import (
	"github.com/3JoB/go-json"
	"github.com/3JoB/unsafeConvert"

	"github.com/3JoB/ulib/err"
)

var ErrNilPointer error = &err.Err{Op: "json.RawMessage", Err: "UnmarshalJSON on nil pointer"}

// RawMessage is a raw encoded JSON value. It implements Marshaler and Unmarshaler and can be used to delay JSON decoding or precompute a JSON encoding.
type RawMessage []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m RawMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return unsafeConvert.BytesReflect("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return ErrNilPointer
	}
	*m = append((*m)[0:0], data...)
	return nil
}

var _ json.Marshaler = (*RawMessage)(nil)
var _ json.Unmarshaler = (*RawMessage)(nil)
