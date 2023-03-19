package fsutil

import (
	"io/fs"
	"os"
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

// This function will check if the target is a directory.
//
// **When the target does not exist or other errors occur, it will return `false`**
func IsDir(path string) bool {
	info, err := Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// This function checks if the target exists.
func IsExist(path string) bool {
	if _, err := Stat(path); err == nil {
		return true
	}
	// return os.IsNotExist(err)
	return false
}

func ReadDirRaw(path string) ([]fs.DirEntry, error) {
	return os.ReadDir(path)
}

func ReadDirInfo(path string) (r []fs.FileInfo) {
	if fr, err := ReadDirRaw(path); err != nil {
		return nil
	} else {
		for _, fs := range fr {
			info, _ := fs.Info()
			r = append(r, info)
		}
		return r
	}
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
