package fsutil

import (
	"bufio"
	"errors"
	"io/fs"
	"os"

	"github.com/3JoB/unsafeConvert"
	"github.com/spf13/cast"
)

var (
	ErrNotExist error = errors.New("no file/folder found")
)

type FS struct {
	Path  string
	Data  string
	TRUNC bool
}

func File(src string) *FS {
	fs := &FS{
		Path: src,
	}
	return fs
}

func (f *FS) CopyTo(dst string) error {
	return copyTo(f.Path, dst)
}

func (f *FS) SetTrunc() *FS {
	if f.TRUNC {
		f.TRUNC = false
	} else {
		f.TRUNC = true
	}
	return f
}

func (f *FS) Write(d any) error {
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
		file.Close()
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	switch d := d.(type) {
	case string:
		writer.Write(unsafeConvert.Bytes(d))
	case []byte:
		writer.Write(d)
	default:
		writer.Write(unsafeConvert.Bytes(cast.ToString(d)))
	}
	writer.Flush()
	return nil
}

func Mkdir(path string, mode ...fs.FileMode) error {
	if len(mode) != 0 {
		return os.MkdirAll(path, mode[0])
	}
	return os.MkdirAll(path, os.ModePerm)
}