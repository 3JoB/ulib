package fsutil

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sync"
)

var bytePool = sync.Pool{
	New: func() any {
		return make([]byte, 4096)
	},
}

func Copy(w io.Writer, r io.Reader) (int64, error) {
	if wt, ok := r.(io.WriterTo); ok {
		return wt.WriteTo(w)
	}
	if rt, ok := w.(io.ReaderFrom); ok {
		return rt.ReadFrom(r)
	}
	v := bytePool.Get()
	buf := v.([]byte)
	n, err := io.CopyBuffer(w, r, buf)
	bytePool.Put(v)
	return n, err
}

func copyTo(src, dst string) error {
	if src == dst {
		return ErrMethods
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

	sb, db := ReaderWriter(s, d)

	if _, err := io.Copy(db, sb); err != nil {
		return err
	}
	if err := db.Flush(); err != nil {
		return err
	}
	return err
}

func Move(src, dst string) error {
	if err := CopyAll(src, dst); err != nil {
		return err
	}
	return Remove(src)
}

func CopyAll(src, dst string) error {
	src, dst = filepath.Clean(src), filepath.Clean(dst)
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
				srcPath := path.Join(src, entry.Name())
				dstPath := path.Join(dst, entry.Name())

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
		return copyTo(src, path.Join(dst, path.Base(src)))
	}
	return nil
}
