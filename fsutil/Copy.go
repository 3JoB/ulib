package fsutil

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func copyTo(src, dst string) error {
	if src == dst {
		return ErrMethods
	}
	s, err := OpenFile(src, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	d, err := OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer s.Close()
	defer d.Close()

	sb := bufio.NewReader(s)
	db := bufio.NewWriter(d)

	if _, err := io.Copy(db, sb); err != nil {
		return err
	}
	if err := db.Flush(); err != nil {
		return err
	}
	return err
}

func Copy(src, dst string) error {
	return copyTo(src, dst)
}

func Move(src, dst string) error {
	if err := CopyAll(src, dst); err != nil {
		return err
	}
	return Remove(src)
}

func CopyAll(src, dst string) error {
	src, dst = CleanPaths(src), CleanPaths(dst)
	if IsDir(src) {
		if !IsDir(dst) {
			if IsFile(dst) {
				return fmt.Errorf("cannot copy directory to file src=%v dst=%v", src, dst)
			}
		}
		s, err := os.Stat(src)
		if err != nil {
			return err
		}
		if err := Mkdir(dst, s.Mode()); err != nil {
			return err
		}
		if entries, err := ReadDirRaw(src); err != nil {
			return err
		} else {
			for _, entry := range entries {
				srcPath := JoinPaths(src, entry.Name())
				dstPath := JoinPaths(dst, entry.Name())

				if entry.IsDir() {
					if err := copyTo(src, dst); err != nil {
						return err
					}
				} else {
					// Skip symlinks.
					if entry.Type()&os.ModeSymlink != 0 {
						continue
					}
					if err := copyTo(srcPath, dstPath); err != nil {
						return err
					}
				}
			}
		}
	} else {
		if IsFile(dst) {
			return copyTo(src, dst)
		}
		return copyTo(src, JoinPaths(dst, BasePaths(src)))
	}
	return nil
}
