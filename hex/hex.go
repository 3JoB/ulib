package hex

import (
	"encoding/hex"

	"github.com/3JoB/unsafeConvert"
)

// Encode encodes src into EncodedLen(len(src)) bytes of dst. As a convenience,
// it returns the number of bytes written to dst, but this value is always
// EncodedLen(len(src)). Encode implements hexadecimal encoding.
func Encode(src []byte) []byte {
	dst := make([]byte, len(src)*2)
	hex.Encode(dst, src)
	return dst
}

// EncodeToString returns the hexadecimal encoding of src.
func EncodeToString(src []byte) string {
	return unsafeConvert.StringPointer(Encode(src))
}

func Decode(b []byte) ([]byte, error) {
	n, err := hex.Decode(b, b)
	return b[:n], err
}

// DecodeString returns the bytes represented by the hexadecimal string s.
//
// DecodeString expects that src contains only hexadecimal
// characters and that src has even length.
// If the input is malformed, DecodeString returns
// the bytes decoded before the error.
func DecodeString(s string) ([]byte, error) {
	src := unsafeConvert.BytePointer(s)
	// We can use the source slice itself as the destination
	// because the decode loop increments by one and then the 'seen' byte is not used anymore.
	n, err := hex.Decode(src, src)
	return src[:n], err
}

func DecodeStringCopy(s string) ([]byte, error) {
	src := unsafeConvert.BytePointer(s)
	csrc := make([]byte, len(src))
	copy(csrc, src)
	n, err := hex.Decode(csrc, csrc)
	return csrc[:n], err
}
