package null

import "io"

type NUL struct{}

func New() io.Writer {
	l := &NUL{}
	return l
}

func (l *NUL) Write(b []byte) (n int, err error) {
	return 0, nil
}
