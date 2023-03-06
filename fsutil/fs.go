package fsutil

import (
	"bufio"
	"io"
	"io/fs"
	"os"

	"github.com/3JoB/unsafeConvert"
	"github.com/spf13/cast"
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

func Open(v string) (*os.File, error) {
	return os.Open(v)
}

func OpenRead(v string) ([]byte, error) {
	o, err := Open(v)
	if err != nil {
		return nil, err
	}
	defer o.Close()
	if data, err := ReadAll(o); err != nil {
		return nil, err
	} else {
		return data, err
	}
}

func ReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

func OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
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
		file, err = OpenFile(f.Path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	} else {
		file, err = OpenFile(f.Path, os.O_WRONLY|os.O_CREATE, 0666)
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

func Remove(v string) error {
	return os.RemoveAll(v)
}
