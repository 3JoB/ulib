package fsutil

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

func IsFile(path string) bool {
	_, err := os.Stat(path)
	return errors.Is(err, fs.ErrNotExist)
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func Exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return os.IsNotExist(err)
	}
}

func GetRunPath() (r string) {
	r, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	return
}

func ReadPath(path string) (f []string) {
	fr, _ := os.ReadDir(path)
	for _, fs := range fr {
		if fs.IsDir() {
			f = append(f, ReadPath(path+"/"+fs.Name())...)
		} else {
			f = append(f, path+"/"+fs.Name())
		}
	}
	return f
}