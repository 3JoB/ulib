package fsutil

import (
	"io/fs"
	"os"

	ph "github.com/3JoB/ulib/fsutil/path"
)

func IsFile(path string) bool {
	if !IsExist(path) {
		return false
	}
	if i, _ := Stat(path); i.IsDir() {
		return false
	}
	return true
}

func IsDir(path string) bool {
	info, err := Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func IsExist(path string) bool {
	if _, err := Stat(path); err == nil {
		return true
	}
	// return os.IsNotExist(err)
	return false
}

func GetRunPath() (r string) {
	r, _ = ph.Abs(ph.Dir(os.Args[0]))
	return
}

func ReadDirRaw(r string) ([]fs.DirEntry, error) {
	return os.ReadDir(r)
}

func ReadDir(path string) (f []string) {
	if fr, err := ReadDirRaw(path); err != nil {
		return nil
	} else {
		for _, fs := range fr {
			f = append(f, path+"/"+fs.Name())
		}
		return f
	}
}

func ReadDirAll(path string) (f []string) {
	if fr, err := ReadDirRaw(path); err != nil {
		return nil
	} else {
		for _, fs := range fr {
			if fs.IsDir() {
				f = append(f, ReadDirAll(path+"/"+fs.Name())...)
			} else {
				f = append(f, path+"/"+fs.Name())
			}
		}
		return f
	}
}
