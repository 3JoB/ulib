package strings

import (
	"io"
	"strings"
)

func ReadFrom(b *strings.Builder, r io.Reader) (n int64, err error) {
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
