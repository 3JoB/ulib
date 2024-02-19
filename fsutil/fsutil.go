package fsutil

import (
	"errors"
)

var (
	ErrNotExist error = errors.New("no file/folder found")
	ErrMethods  error = errors.New("don't use weird methods")
)
