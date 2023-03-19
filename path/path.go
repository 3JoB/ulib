package path

import "path/filepath"

func Abs(v string) (string, error) {
	return filepath.Abs(v)
}

func IsAbs(path string) bool {
	return filepath.IsAbs(path)
}

func IsLocal(path string) bool {
	return filepath.IsLocal(path)
}

func Join(v ...string) string {
	return filepath.Join(v...)
}

func Base(src string) string {
	return filepath.Base(src)
}

func Clean(src string) string {
	return filepath.Clean(src)
}

func Split(src string) (string, string) {
	return filepath.Split(src)
}

func Ext(src string) string {
	return filepath.Ext(src)
}

func Dir(path string) string {
	return filepath.Dir(path)
}

func DirPath(src string) string {
	dir, _ := Split(src)
	return Clean(dir)
}
