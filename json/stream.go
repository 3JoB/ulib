package json

import (
	"errors"

	"github.com/3JoB/unsafeConvert"
	js "github.com/goccy/go-json"
)

// RawMessage is a raw encoded JSON value. It implements Marshaler and Unmarshaler and can be used to delay JSON decoding or precompute a JSON encoding.
type RawMessage []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m RawMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return unsafeConvert.Bytes("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

var _ js.Marshaler = (*RawMessage)(nil)
var _ js.Unmarshaler = (*RawMessage)(nil)
