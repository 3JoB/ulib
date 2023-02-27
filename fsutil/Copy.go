package fsutil

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyTo(src, dst string) error {
	if src == dst {
		return errors.New("ulib.fsutil: don't use weird methods")
	}
	s, err := os.OpenFile(src, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	d, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
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

func CopyAll(src, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)
	if !Exists(src) {
		return ErrNotExist
	}
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
		Mkdir(dst, s.Mode())
		if entries, err := os.ReadDir(src); err != nil {
			return err
		} else {
			for _, entry := range entries {
				srcPath := JoinPaths(src, entry.Name())
				dstPath := JoinPaths(dst, entry.Name())

				if entry.IsDir() {
					copyTo(src, dst)
				} else {
					// Skip symlinks.
					if entry.Type()&os.ModeSymlink != 0 {
						continue
					}
					copyTo(srcPath, dstPath)
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