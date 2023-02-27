package fsutil

import "path/filepath"

func JoinPaths(v ...string) string {
	return filepath.Join(v...)
}

func BasePaths(src string) string {
	return filepath.Base(src)
}

func CleanPaths(src string) string {
	return filepath.Clean(src)
}

func SplitPath(src string) (string, string) {
	return filepath.Split(src)
}

func Ext(src string) string {
	return filepath.Ext(src)
}

func DirPath(src string) string {
	dir, _ := SplitPath(src)
	return CleanPaths(dir)
}
