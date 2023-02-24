package fsutil

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"os"

	"github.com/3JoB/unsafeConvert"
)

type FS struct {
	Path  string
	Data  string
	TRUNC bool
}

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

func File(path string) *FS {
	fs := &FS{
		Path: path,
	}
	return fs
}

func (f *FS) CopyTo(paths string) error {
	if f.Path == paths {
		return errors.New("ulib.fsutil: don't use weird methods")
	}
	s, err := os.OpenFile(f.Path, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	d, err := os.OpenFile(paths, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer s.Close()
	defer d.Close()

	sb := bufio.NewReader(s)
	db := bufio.NewWriter(d)

	if _ ,err = io.Copy(db,sb); err != nil {
		return err
	}
	if err:= db.Flush(); err != nil {
		return err
	}
	return err
}

func (f *FS) SetTrunc() *FS {
	if f.TRUNC {
		f.TRUNC = false
	} else {
		f.TRUNC = true
	}
	return f
}

func (f *FS) Write(d string) error {
	var (
		file *os.File
		err  error
	)
	if f.TRUNC {
		file, err = os.OpenFile(f.Path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	} else {
		file, err = os.OpenFile(f.Path, os.O_WRONLY|os.O_CREATE, 0666)
	}
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write(unsafeConvert.Bytes(d))
	writer.Flush()
	return nil
}
