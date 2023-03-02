package compress

import "errors"

var (
	ErrTargetType error = errors.New("ulib.fsutil.compress: the target directory type is file")
)
