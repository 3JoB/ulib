package hex

import (
	"encoding/hex"

	"github.com/3JoB/unsafeConvert"
)

// EncodeToString returns the hexadecimal encoding of src.
func EncodeToString(src []byte) string {
	dst := make([]byte, len(src)*2)
	hex.Encode(dst, src)
	return unsafeConvert.StringReflect(dst)
}

// DecodeString returns the bytes represented by the hexadecimal string s.
//
// DecodeString expects that src contains only hexadecimal
// characters and that src has even length.
// If the input is malformed, DecodeString returns
// the bytes decoded before the error.
func DecodeString(s string) ([]byte, error) {
	src := unsafeConvert.BytesReflect(s)
	// We can use the source slice itself as the destination
	// because the decode loop increments by one and then the 'seen' byte is not used anymore.
	n, err := hex.Decode(src, src)
	return src[:n], err
}
