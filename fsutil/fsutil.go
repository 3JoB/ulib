package fsutil

import (
	"github.com/3JoB/ulib/err"
)

var (
	ErrNotExist error = &err.Err{Op: "ulib.fsutil", Err: "no file/folder found"}
	ErrMethods  error = &err.Err{Op: "ulib.fsutil", Err: "don't use weird methods"}
)
