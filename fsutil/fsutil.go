package fsutil

import "errors"

var (
	ErrNotExist error = errors.New("ulib.fsutil: no file/folder found")
	ErrMethods  error = errors.New("ulib.fsutil: don't use weird methods")
)
