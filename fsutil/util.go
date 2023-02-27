package fsutil

import (
	"os"
	"path/filepath"
)

func GetRunPath() (r string) {
	r, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	return
}
