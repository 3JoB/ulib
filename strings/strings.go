package strings

import (
	"io"
	"strings"

	"github.com/3JoB/unsafeConvert"
)

func NewBuilder(b []byte) strings.Builder {
	builder := strings.Builder{}
	builder.Write(b)
	return builder
}

func NewBuilders() strings.Builder {
	return strings.Builder{}
}

func NewBuilderString(s string) strings.Builder {
	builder := strings.Builder{}
	builder.WriteString(s)
	return builder
}

func ReadFrom(b strings.Builder, r io.Reader) (n int64, err error) {
	buf := make([]byte, b.Cap())
	for {
		rn, err := r.Read(buf)
		if err != nil || rn == 0 {
			return n, err
		}
		b.Write(buf[:rn])
		n += int64(rn)
	}
}

func Bytes(b strings.Builder) []byte {
	return unsafeConvert.BytesReflect(b.String())
}
